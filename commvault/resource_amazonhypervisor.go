package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAmazonHypervisor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAmazonHypervisorCreate,
		Read:   resourceAmazonHypervisorRead,
		Update: resourceAmazonHypervisorUpdate,
		Delete: resourceAmazonHypervisorDelete,

		Schema: map[string]*schema.Schema{
			"client_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The name of the Amazon hypervisor.",
			},
			"regions": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Specifies the regions used for the Hypervisor",
			},
			"use_iam_role": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Specifies whether you want to use IAM role.",
			},
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies The access key ID for your Amazon account.",
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies The secret key ID for your Amazon account.",
			},
			"access_nodes": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The clients that have the VSA package installed and that act as proxy clients for Amazon hypervisors.",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the company id to which the Amazon Hypervisor should be associated with.",
			},
		},
	}
}

func resourceAmazonHypervisorCreate(d *schema.ResourceData, m interface{}) error {
	clientname := d.Get("client_name").(string)
	regions := d.Get("regions").(string)
	useiamrole := d.Get("use_iam_role").(bool)
	accesskey := d.Get("access_key").(string)
	secretkey := d.Get("secret_key").(string)
	accessnodes := d.Get("access_nodes").(string)
	companyid := d.Get("company_id").(int)
	if !useiamrole {
		if accesskey == "" || secretkey == "" {
			return fmt.Errorf("Accesskey and Secretkey Cannot be Blank!")
		}
	}
	apiResp := handler.AmazonHypCreateHandler(clientname, regions, useiamrole, accesskey, secretkey, accessnodes, companyid)
	d.SetId(string(apiResp.Response.Entity.ClientId))
	return resourceAmazonHypervisorRead(d, m)
}

func resourceAmazonHypervisorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAmazonHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAmazonHypervisorRead(d, m)
}

func resourceAmazonHypervisorDelete(d *schema.ResourceData, m interface{}) error {
	clientID := d.Id()
	resp := handler.AmazonHypDeleteHandler(clientID)
	if resp.Response.ErrorCode != "0" {
		return fmt.Errorf("Error retiring the VMWare Hypervisor")
	}
	d.SetId("")
	return nil
}
