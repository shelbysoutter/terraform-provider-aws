package lightsail_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lightsail"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func init() {
	resource.AddTestSweepers("aws_lightsail_static_ip", &resource.Sweeper{
		Name: "aws_lightsail_static_ip",
		F:    testSweepLightsailStaticIps,
	})
}

func testSweepLightsailStaticIps(region string) error {
	client, err := acctest.SharedRegionalSweeperClient(region)
	if err != nil {
		return fmt.Errorf("Error getting client: %s", err)
	}
	conn := client.(*conns.AWSClient).LightsailConn

	input := &lightsail.GetStaticIpsInput{}

	for {
		output, err := conn.GetStaticIps(input)
		if err != nil {
			if acctest.SkipSweepError(err) {
				log.Printf("[WARN] Skipping Lightsail Static IP sweep for %s: %s", region, err)
				return nil
			}
			return fmt.Errorf("Error retrieving Lightsail Static IPs: %s", err)
		}

		if len(output.StaticIps) == 0 {
			log.Print("[DEBUG] No Lightsail Static IPs to sweep")
			return nil
		}

		for _, staticIp := range output.StaticIps {
			name := aws.StringValue(staticIp.Name)

			log.Printf("[INFO] Deleting Lightsail Static IP %s", name)
			_, err := conn.ReleaseStaticIp(&lightsail.ReleaseStaticIpInput{
				StaticIpName: aws.String(name),
			})
			if err != nil {
				return fmt.Errorf("Error deleting Lightsail Static IP %s: %s", name, err)
			}
		}

		if output.NextPageToken == nil {
			break
		}
		input.PageToken = output.NextPageToken
	}

	return nil
}

func TestAccLightsailStaticIP_basic(t *testing.T) {
	var staticIp lightsail.StaticIp
	staticIpName := fmt.Sprintf("tf-test-lightsail-%s", sdkacctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, lightsail.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckStaticIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStaticIPConfig_basic(staticIpName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckStaticIPExists("aws_lightsail_static_ip.test", &staticIp),
				),
			},
		},
	})
}

func TestAccLightsailStaticIP_disappears(t *testing.T) {
	var staticIp lightsail.StaticIp
	staticIpName := fmt.Sprintf("tf-test-lightsail-%s", sdkacctest.RandString(5))

	staticIpDestroy := func(*terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).LightsailConn
		_, err := conn.ReleaseStaticIp(&lightsail.ReleaseStaticIpInput{
			StaticIpName: aws.String(staticIpName),
		})

		if err != nil {
			return fmt.Errorf("Error deleting Lightsail Static IP in disapear test")
		}

		return nil
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); testAccPreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, lightsail.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckStaticIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStaticIPConfig_basic(staticIpName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckStaticIPExists("aws_lightsail_static_ip.test", &staticIp),
					staticIpDestroy,
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckStaticIPExists(n string, staticIp *lightsail.StaticIp) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Lightsail Static IP ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).LightsailConn

		resp, err := conn.GetStaticIp(&lightsail.GetStaticIpInput{
			StaticIpName: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if resp == nil || resp.StaticIp == nil {
			return fmt.Errorf("Static IP (%s) not found", rs.Primary.ID)
		}
		*staticIp = *resp.StaticIp
		return nil
	}
}

func testAccCheckStaticIPDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_lightsail_static_ip" {
			continue
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).LightsailConn

		resp, err := conn.GetStaticIp(&lightsail.GetStaticIpInput{
			StaticIpName: aws.String(rs.Primary.ID),
		})

		if err == nil {
			if resp.StaticIp != nil {
				return fmt.Errorf("Lightsail Static IP %q still exists", rs.Primary.ID)
			}
		}

		// Verify the error
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return nil
			}
		}
		return err
	}

	return nil
}

func testAccStaticIPConfig_basic(staticIpName string) string {
	return fmt.Sprintf(`
resource "aws_lightsail_static_ip" "test" {
  name = "%s"
}
`, staticIpName)
}
