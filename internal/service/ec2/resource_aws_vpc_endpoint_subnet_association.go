package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfec2 "github.com/hashicorp/terraform-provider-aws/aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/service/ec2/finder"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/service/ec2/waiter"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
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
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)

func ResourceVPCEndpointSubnetAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCEndpointSubnetAssociationCreate,
		Read:   resourceVPCEndpointSubnetAssociationRead,
		Delete: resourceVPCEndpointSubnetAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceVPCEndpointSubnetAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).EC2Conn

	endpointID := d.Get("vpc_endpoint_id").(string)
	subnetID := d.Get("subnet_id").(string)
	// Human friendly ID for error messages since d.Id() is non-descriptive
	id := fmt.Sprintf("%s/%s", endpointID, subnetID)

	input := &ec2.ModifyVpcEndpointInput{
		VpcEndpointId: aws.String(endpointID),
		AddSubnetIds:  aws.StringSlice([]string{subnetID}),
	}

	log.Printf("[DEBUG] Creating VPC Endpoint Subnet Association: %s", input)

	// See https://github.com/hashicorp/terraform-provider-aws/issues/3382.
	// Prevent concurrent subnet association requests and delay between requests.
	mk := "vpc_endpoint_subnet_association_" + endpointID
	conns.GlobalMutexKV.Lock(mk)
	defer conns.GlobalMutexKV.Unlock(mk)

	c := &resource.StateChangeConf{
		Delay:   1 * time.Minute,
		Timeout: 3 * time.Minute,
		Target:  []string{"ok"},
		Refresh: func() (interface{}, string, error) {
			output, err := conn.ModifyVpcEndpoint(input)

			return output, "ok", err
		},
	}
	_, err := c.WaitForState()

	if err != nil {
		return fmt.Errorf("error creating VPC Endpoint Subnet Association (%s): %w", id, err)
	}

	d.SetId(tfec2.VPCEndpointSubnetAssociationCreateID(endpointID, subnetID))

	_, err = tfec2.WaitVPCEndpointAvailable(conn, endpointID, d.Timeout(schema.TimeoutCreate))

	if err != nil {
		return fmt.Errorf("error waiting for VPC Endpoint (%s) to become available: %w", endpointID, err)
	}

	return resourceVPCEndpointSubnetAssociationRead(d, meta)
}

func resourceVPCEndpointSubnetAssociationRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).EC2Conn

	endpointID := d.Get("vpc_endpoint_id").(string)
	subnetID := d.Get("subnet_id").(string)
	// Human friendly ID for error messages since d.Id() is non-descriptive
	id := fmt.Sprintf("%s/%s", endpointID, subnetID)

	err := tfec2.FindVPCEndpointSubnetAssociationExists(conn, endpointID, subnetID)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] VPC Endpoint Subnet Association (%s) not found, removing from state", id)
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error reading VPC Endpoint Subnet Association (%s): %w", id, err)
	}

	return nil
}

func resourceVPCEndpointSubnetAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).EC2Conn

	endpointID := d.Get("vpc_endpoint_id").(string)
	subnetID := d.Get("subnet_id").(string)
	// Human friendly ID for error messages since d.Id() is non-descriptive
	id := fmt.Sprintf("%s/%s", endpointID, subnetID)

	input := &ec2.ModifyVpcEndpointInput{
		VpcEndpointId:   aws.String(endpointID),
		RemoveSubnetIds: aws.StringSlice([]string{subnetID}),
	}

	log.Printf("[DEBUG] Deleting VPC Endpoint Subnet Association: %s", id)
	_, err := conn.ModifyVpcEndpoint(input)

	if tfawserr.ErrCodeEquals(err, tfec2.ErrCodeInvalidVPCEndpointIdNotFound) || tfawserr.ErrCodeEquals(err, tfec2.ErrCodeInvalidSubnetIdNotFound) || tfawserr.ErrCodeEquals(err, tfec2.ErrCodeInvalidParameter) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error deleting VPC Endpoint Subnet Association (%s): %w", id, err)
	}

	_, err = tfec2.WaitVPCEndpointAvailable(conn, endpointID, d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return fmt.Errorf("error waiting for VPC Endpoint (%s) to become available: %w", endpointID, err)
	}

	return nil
}
