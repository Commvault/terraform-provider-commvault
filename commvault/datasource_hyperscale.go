package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceHyperscale() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadHyperscale,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadHyperscale(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvHyperscaleIdByName(d.Get("name").(string))

	if resp.HyperscaleId > 0 {
		d.SetId(strconv.Itoa(resp.HyperscaleId))
	} else {
		return fmt.Errorf("unknown hyperscale %s", d.Get("name").(string))
	}

	return nil
}
