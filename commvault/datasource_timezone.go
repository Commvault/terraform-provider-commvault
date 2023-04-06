package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceTimezone() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadTimezone,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadTimezone(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvTimezoneIdByName(d.Get("name").(string))

	if resp.TimezoneId > 0 {
		d.SetId(strconv.Itoa(resp.TimezoneId))
	} else {
		return fmt.Errorf("unknown timezone %s", d.Get("name").(string))
	}

	return nil
}
