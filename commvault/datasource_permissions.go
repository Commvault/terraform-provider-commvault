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
			"categoryid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadPermissions(d *schema.ResourceData, m interface{}) error {
	pid, cid := handler.CvPermissionIdByName(d.Get("name").(string))
	print("pid %d cid %d", pid, cid)
	if pid > 0 {
		d.SetId(strconv.Itoa(pid))
		d.Set("categoryid", cid)
	} else {
		return fmt.Errorf("unknown permission %s", d.Get("name").(string))
	}

	return nil
}
