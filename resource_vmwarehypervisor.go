package main

import (
	"fmt"

	"github.com/TerraformProvider/handler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVMWareHypervisor() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMWareHypervisorCreate,
		Read:   resourceVMWareHypervisorRead,
		Update: resourceVMWareHypervisorUpdate,
		Delete: resourceVMWareHypervisorDelete,

		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
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

func resourceVMWareHypervisorCreate(d *schema.ResourceData, m interface{}) error {
	displayname := d.Get("display_name").(string)
	hostname := d.Get("host_name").(string)
	username := d.Get("user_name").(string)
	password := d.Get("password").(string)
	accessnodes := d.Get("access_nodes").(string)
	apiResp := handler.VMWareHypCreateHandler(displayname, hostname, username, password, accessnodes)
	d.SetId(string(apiResp.Response.Entity.ClientId))
	return resourceVMWareHypervisorRead(d, m)
}

func resourceVMWareHypervisorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVMWareHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceVMWareHypervisorRead(d, m)
}

func resourceVMWareHypervisorDelete(d *schema.ResourceData, m interface{}) error {
	clientID := d.Id()
	resp := handler.VMWareHypDeleteHandler(clientID)
	if resp.Response.ErrorCode != "0" {
		return fmt.Errorf("Error retiring the VMWare Hypervisor")
	}
	d.SetId("")
	return nil
}
