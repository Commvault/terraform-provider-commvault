package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceCompany() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadCompany,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadCompany(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvCompanyIdByName(d.Get("name").(string))

	if resp.Providers != nil && len(resp.Providers) > 0 && resp.Providers[0].ShortName.Id > 0 {
		d.SetId(strconv.Itoa(resp.Providers[0].ShortName.Id))
	} else {
		return fmt.Errorf("unknown company %s", d.Get("name").(string))
	}

	return nil
}
