package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
import "github.com/TerraformProvider/handler"

func resourceAssociateVMToPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssociateVMToPlanCreate,
		Read:   resourceAssociateVMToPlanRead,
		Update: resourceAssociateVMToPlanUpdate,
		Delete: resourceAssociateVMToPlanDelete,

		Schema: map[string]*schema.Schema{
			"plan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cname":&schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nplan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAssociateVMToPlanUpdate(data *schema.ResourceData, i interface{}) error {
	planName := data.Get("plan").(string)
	newPlanName := data.Get("nplan").(string)
	vmname := data.Get("cname").(string)
	handler.AssociatePlanToVM(newPlanName,vmname)
	sid := planName + vmname
	data.SetId(sid)
	return resourceAssociateVMToPlanRead(data,i)
}

func resourceAssociateVMToPlanRead(data *schema.ResourceData, i interface{}) error {
	return nil
}

func resourceAssociateVMToPlanCreate(data *schema.ResourceData, i interface{}) error {
	planName := data.Get("plan").(string)
	vmname := data.Get("cname").(string)
	handler.AssociatePlanToVM(planName,vmname)
	sid := planName + vmname
	data.SetId(sid)
	return resourceAssociateVMToPlanRead(data,i)
}

func resourceAssociateVMToPlanDelete(data *schema.ResourceData, i interface{}) error {
	return nil
}
