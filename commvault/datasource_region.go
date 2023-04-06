package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceRegion() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadRegion,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadRegion(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvRegionIdByName(d.Get("name").(string))

	if resp.RegionId > 0 {
		d.SetId(strconv.Itoa(resp.RegionId))
	} else {
		return fmt.Errorf("unknown region %s", d.Get("name").(string))
	}

	return nil
}
