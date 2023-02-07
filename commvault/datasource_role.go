package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceRole() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadRole,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadRole(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvRoleIdByName(d.Get("name").(string))

	if resp.RoleId > 0 {
		d.SetId(strconv.Itoa(resp.RoleId))
	} else {
		return fmt.Errorf("unknown plan %s", d.Get("name").(string))
	}

	return nil
}
