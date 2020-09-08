package main

import(
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/TerraformProvider/handler"
)

func resourcePlan() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlanCreate,
		Read:   resourcePlanRead,
		Update: resourcePlanUpdate,
		Delete: resourcePlanDelete,

		Schema: map[string]*schema.Schema{
			"planname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rpoinminutes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},			
			"slainminutes":&schema.Schema{
				Type:schema.TypeString,
				Optional:true,
			},
			"backupdestname":&schema.Schema{
				Type:schema.TypeString,
				Optional:true,
			},
			"backupdeststorage":&schema.Schema{
				Type:schema.TypeString,
				Optional:true,
			},			
		},
	}
}


func resourcePlanCreate(d *schema.ResourceData, m interface{}) error {
	planName := d.Get("planname").(string)
	backupDestName := d.Get("backupdestname").(string)
	backupDestStorage := d.Get("backupdeststorage").(string)
	rpoinmin := d.Get("rpoinminutes").(string)
	apiResp := handler.PlanCreate(planName,backupDestName,backupDestStorage,rpoinmin)
	d.SetId(string(apiResp.Plan.Summary.Plan.PlanId))
	return resourcePlanRead(d,m)
}

func resourcePlanRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePlanUpdate(d *schema.ResourceData, m interface{}) error {
	rpoinmin := d.Get("rpoinminutes").(string)
	slainminutes := d.Get("slainminutes").(string)
	id := d.Id()
	handler.PlanUpdate(rpoinmin,slainminutes,id)
	return resourcePlanRead(d, m)
}

func resourcePlanDelete(d *schema.ResourceData, m interface{}) error {
	planId := d.Id()
	handler.PlanDelete(planId)
	d.SetId("")
	return nil
}



