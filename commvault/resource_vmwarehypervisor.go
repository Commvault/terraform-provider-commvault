package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

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
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The display name of the hypervisor.",
			},
			"host_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The host name of the hypervisor.",
			},
			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The user name for the account.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The password for the account.",
			},
			"access_nodes": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The clients that have the VSA package installed and that act as proxy clients for hypervisors.",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the company id to which the Vmware Hypervisor should be associated with.",
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
	companyid := d.Get("company_id").(int)
	apiResp := handler.VMWareHypCreateHandler(displayname, hostname, username, password, accessnodes, companyid)
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
