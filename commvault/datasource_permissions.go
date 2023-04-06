package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourcePermissions() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadPermissions,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadPermissions(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvPermissionIdByName(d.Get("name").(string))

	if resp.PermissionId > 0 {
		d.SetId(strconv.Itoa(resp.PermissionId))
	} else {
		return fmt.Errorf("unknown permission %s", d.Get("name").(string))
	}

	return nil
}
