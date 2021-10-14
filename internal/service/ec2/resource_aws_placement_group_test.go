package aws

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/service/ec2/finder"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)

func init() {
	resource.AddTestSweepers("aws_placement_group", &resource.Sweeper{
		Name: "aws_placement_group",
		F:    testSweepEc2PlacementGroups,
		Dependencies: []string{
			"aws_autoscaling_group",
			"aws_instance",
			"aws_launch_template",
			"aws_spot_fleet_request",
		},
	})
}

func testSweepEc2PlacementGroups(region string) error {
	client, err := sweep.SharedRegionalSweepClient(region)
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}
	conn := client.(*conns.AWSClient).EC2Conn
	input := &ec2.DescribePlacementGroupsInput{}
	sweepResources := make([]*sweep.SweepResource, 0)

	output, err := conn.DescribePlacementGroups(input)

	if sweep.SkipSweepError(err) {
		log.Printf("[WARN] Skipping EC2 Placement Group sweep for %s: %s", region, err)
		return nil
	}

	if err != nil {
		return fmt.Errorf("error listing EC2 Placement Groups (%s): %w", region, err)
	}

	for _, placementGroup := range output.PlacementGroups {
		r := ResourcePlacementGroup()
		d := r.Data(nil)
		d.SetId(aws.StringValue(placementGroup.GroupName))

		sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
	}

	err = sweep.SweepOrchestrator(sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping EC2 Placement Groups (%s): %w", region, err)
	}

	return nil
}

func TestAccAWSPlacementGroup_basic(t *testing.T) {
	var pg ec2.PlacementGroup
	resourceName := "aws_placement_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, ec2.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAWSPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSPlacementGroupConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "strategy", "cluster"),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "ec2", fmt.Sprintf("placement-group/%s", rName)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAWSPlacementGroup_disappears(t *testing.T) {
	var pg ec2.PlacementGroup
	resourceName := "aws_placement_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, ec2.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAWSPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSPlacementGroupConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					acctest.CheckResourceDisappears(acctest.Provider, ResourcePlacementGroup(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAWSPlacementGroup_Tags(t *testing.T) {
	var pg ec2.PlacementGroup
	resourceName := "aws_placement_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, ec2.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAWSPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSPlacementGroupConfigTags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAWSPlacementGroupConfigTags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccAWSPlacementGroupConfigTags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2")),
			},
		},
	})
}

func TestAccAWSPlacementGroup_PartitionCount(t *testing.T) {
	var pg ec2.PlacementGroup
	resourceName := "aws_placement_group.test"
	rName := sdkacctest.RandomWithPrefix("tf-acc-partition")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		ErrorCheck:   acctest.ErrorCheck(t, ec2.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckAWSPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSPlacementGroupConfigPartitionCount(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSPlacementGroupExists(resourceName, &pg),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "strategy", "partition"),
					resource.TestCheckResourceAttr(resourceName, "partition_count", "7"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAWSPlacementGroupDestroy(s *terraform.State) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_placement_group" {
			continue
		}

		_, err := tfec2.FindPlacementGroupByName(conn, rs.Primary.ID)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("EC2 Placement Group %s still exists", rs.Primary.ID)
	}

	return nil
}

func testAccCheckAWSPlacementGroupExists(n string, v *ec2.PlacementGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No EC2 Placement Group ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn

		output, err := tfec2.FindPlacementGroupByName(conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccAWSPlacementGroupConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_placement_group" "test" {
  name     = %[1]q
  strategy = "cluster"
}
`, rName)
}

func testAccAWSPlacementGroupConfigTags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_placement_group" "test" {
  name     = %[1]q
  strategy = "cluster"

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccAWSPlacementGroupConfigTags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_placement_group" "test" {
  name     = %[1]q
  strategy = "cluster"

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccAWSPlacementGroupConfigPartitionCount(rName string) string {
	return fmt.Sprintf(`
resource "aws_placement_group" "test" {
  name            = %[1]q
  strategy        = "partition"
  partition_count = 7
}
`, rName)
}
