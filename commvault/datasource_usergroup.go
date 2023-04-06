package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadUserGroup,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadUserGroup(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvUserGroupIdByName(d.Get("name").(string))

	if resp.UserGroups != nil && len(resp.UserGroups) > 0 && resp.UserGroups[0].UserGroupEntity.UserGroupId > 0 {
		d.SetId(strconv.Itoa(resp.UserGroups[0].UserGroupEntity.UserGroupId))
	} else {
		return fmt.Errorf("unknown user group %s", d.Get("name").(string))
	}

	return nil
}
