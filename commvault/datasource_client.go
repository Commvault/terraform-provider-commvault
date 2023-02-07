package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceClient() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadClient,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadClient(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvClientIdByName(d.Get("name").(string))

	d.Set("type", 3)

	if resp.ClientProperties != nil && len(resp.ClientProperties) > 0 && resp.ClientProperties[0].Client.ClientEntity.ClientId > 0 {
		d.SetId(strconv.Itoa(resp.ClientProperties[0].Client.ClientEntity.ClientId))
	} else {
		return fmt.Errorf("unknown user %s", d.Get("name").(string))
	}

	return nil
}
