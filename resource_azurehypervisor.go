package main

import (
	"fmt"

	"github.com/TerraformProvider/handler"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"application_password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"access_nodes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
	apiResp := handler.AzureHypCreateHandler(displayname, subscriptionID, tenantID, applicationID, applicationPWD, accessnodes)
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
