package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourcePlan() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadPlan,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadPlan(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvPlanIdByName(d.Get("name").(string))

	if resp.Plans != nil && len(resp.Plans) > 0 && resp.Plans[0].Plan.PlanId > 0 {
		d.SetId(strconv.Itoa(resp.Plans[0].Plan.PlanId))
	} else {
		return fmt.Errorf("unknown plan %s", d.Get("name").(string))
	}

	return nil
}
