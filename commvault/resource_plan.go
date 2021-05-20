package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePlan() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlanCreate,
		Read:   resourcePlanRead,
		Update: resourcePlanUpdate,
		Delete: resourcePlanDelete,

		Schema: map[string]*schema.Schema{
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the Plan name used for creation of the plan.",
			},
			"retention_period_days": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the number of days that the software retains the data.",
			},
			"backup_destination_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the destination name for the backup.",
			},
			"backup_destination_storage": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the backup destination storage used for the plan.",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the companyid to which the created plan needs to be associated with.",
			},
		},
	}
}

func resourcePlanCreate(d *schema.ResourceData, m interface{}) error {

	var createPlanRequest handler.ApiCreatePlanReq
	createPlanRequest.PlanName = d.Get("plan_name").(string)
	var backupDestination handler.BackupDestination
	backupDestination.BackupDestinationName = d.Get("backup_destination_name").(string)
	backupDestination.RetentionPeriodDays = d.Get("retention_period_days").(int)
	backupDestination.StoragePool.Name = d.Get("backup_destination_storage").(string)
	createPlanRequest.BackupDestinations = append(createPlanRequest.BackupDestinations, backupDestination)
	apiResp := handler.PlanCreate(createPlanRequest, d.Get("company_id").(int))
	if apiResp.Plan.ID > 0 {
		d.SetId(strconv.Itoa(apiResp.Plan.ID))
		return resourcePlanRead(d, m)
	}
	return fmt.Errorf("error in creation of plan")

}

func resourcePlanRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePlanUpdate(d *schema.ResourceData, m interface{}) error {
	rpoinmin := d.Get("rpoinminutes").(string)
	slainminutes := d.Get("slainminutes").(string)
	id := d.Id()
	handler.PlanUpdate(rpoinmin, slainminutes, id)
	return resourcePlanRead(d, m)
}

func resourcePlanDelete(d *schema.ResourceData, m interface{}) error {
	planID := d.Id()
	genericResp := handler.PlanDelete(planID)
	if genericResp.ErrorCode != 0 {
		return fmt.Errorf("Error in deletion of plan")
	}
	d.SetId("")
	return nil
}
