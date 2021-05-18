package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVMGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMGroupCreate,
		Read:   resourceVMGroupRead,
		Update: resourceVMGroupUpdate,
		Delete: resourceVMGroupDelete,

		Schema: map[string]*schema.Schema{
			"vm_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"plan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"vms": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"company_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceVMGroupCreate(d *schema.ResourceData, m interface{}) error {
	vmGroupName := d.Get("vm_group_name").(string)
	if vmGroupName == "" {
		return fmt.Errorf("vmgroup Name cannot be empty")
	}
	clientid := d.Get("client_id").(int)
	planid := d.Get("plan_id").(int)
	vms := d.Get("vms").(*schema.Set).List()
	tags := d.Get("tags").(*schema.Set).List()
	companyid := d.Get("company_id").(int)
	if len(vms) == 0 && len(tags) == 0 {
		return fmt.Errorf("VM's list or tags  list are missing from input")
	}
	vmnames := make([]string, len(vms))
	tagnames := make([]string, len(tags))
	for i, n := range vms {
		vmnames[i] = n.(string)
	}
	for i, n := range tags {
		tagnames[i] = n.(string)
	}
	apiResp := handler.VMGroupCreate(vmGroupName, planid, clientid, vmnames, tagnames, companyid)
	if apiResp.Response.ErrorCode != "0" {
		return fmt.Errorf("Error creating vmgroup")
	}
	d.SetId(apiResp.Response.Entity.SubclientId)
	return resourceVMGroupRead(d, m)
}

func resourceVMGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVMGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceVMGroupRead(d, m)
}

func resourceVMGroupDelete(d *schema.ResourceData, m interface{}) error {
	vmgroupid := d.Id()
	resp := handler.VMGroupDelete(vmgroupid)
	if resp.Response.ErrorCode != "0" {
		return fmt.Errorf("Couldnt delete vmgroupid " + vmgroupid)
	}
	d.SetId("")
	return nil
}
