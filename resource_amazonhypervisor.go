package main

import (
	"fmt"

	"github.com/TerraformProvider/handler"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"regions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_iam_role": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_nodes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
	if !useiamrole {
		if accesskey == "" || secretkey == "" {
			return fmt.Errorf("Accesskey and Secretkey Cannot be Blank!")
		}
	}
	apiResp := handler.AmazonHypCreateHandler(clientname, regions, useiamrole, accesskey, secretkey, accessnodes)
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
