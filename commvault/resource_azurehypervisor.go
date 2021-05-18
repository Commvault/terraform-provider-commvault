package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAzureHypervisor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureHypervisorCreate,
		Read:   resourceAzureHypervisorRead,
		Update: resourceAzureHypervisorUpdate,
		Delete: resourceAzureHypervisorDelete,

		Schema: map[string]*schema.Schema{
			"hypervisor_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The name of the hypervisor.",
			},
			"subscription_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The subscription ID for your Azure account.",
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The tenant ID for your Azure account.",
			},
			"application_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The application ID of the tenant.",
			},
			"application_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The password for the application ID of the tenant.",
			},
			"access_nodes": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The clients that have the VSA package installed and that act as proxy clients for Azure hypervisors.",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the company id to which the Azure Hypervisor should be associated with.",
			},
		},
	}
}

func resourceAzureHypervisorCreate(d *schema.ResourceData, m interface{}) error {
	displayname := d.Get("hypervisor_name").(string)
	subscriptionID := d.Get("subscription_id").(string)
	tenantID := d.Get("tenant_id").(string)
	applicationID := d.Get("application_id").(string)
	applicationPWD := d.Get("application_password").(string)
	accessnodes := d.Get("access_nodes").(string)
	companyid := d.Get("company_id").(int)
	apiResp := handler.AzureHypCreateHandler(displayname, subscriptionID, tenantID, applicationID, applicationPWD, accessnodes, companyid)
	d.SetId(string(apiResp.Response.Entity.ClientId))
	return resourceAzureHypervisorRead(d, m)
}

func resourceAzureHypervisorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzureHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAzureHypervisorRead(d, m)
}

func resourceAzureHypervisorDelete(d *schema.ResourceData, m interface{}) error {
	clientID := d.Id()
	resp := handler.AzureHypDeleteHandler(clientID)
	if resp.Response.ErrorCode != "0" {
		return fmt.Errorf("Error retiring the Azure Hypervisor")
	}
	d.SetId("")
	return nil
}
