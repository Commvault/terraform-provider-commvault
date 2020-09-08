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
			"displayname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username":&schema.Schema{
				Type:schema.TypeString,
				Required:true,
			},
			"password":&schema.Schema{
				Type:schema.TypeString,
				Required:true,
			},
			"accessnodes":&schema.Schema{
				Type:schema.TypeString,
				Required:true,
			},
		},
	}
}

func resourceVMWareHypervisorCreate(d *schema.ResourceData, m interface{}) error {
	displayname := d.Get("displayname").(string)
	hostname := d.Get("hostname").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	accessnodes := d.Get("accessnodes").(string)
	apiResp := handler.VMWareHypCreateHandler(displayname,hostname,username,password,accessnodes)
	d.SetId(string(apiResp.Response.Entity.ClientId))
	return resourceVMWareHypervisorRead(d,m)
}

func resourceVMWareHypervisorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVMWareHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceVMWareHypervisorRead(d, m)
}

func resourceVMWareHypervisorDelete(d *schema.ResourceData, m interface{}) error {
	clientId := d.Id()
	resp := handler.VMWareHypDeleteHandler(clientId)
	if resp.Response.ErrorCode != "0" {
		return fmt.Errorf("Error retiring the VMWare Hypervisor")
	}
	d.SetId("")
	return nil
}



