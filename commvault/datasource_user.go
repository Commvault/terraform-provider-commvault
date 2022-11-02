package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceUser() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadUser,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadUser(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvGetUserByName(d.Get("name").(string))

	if resp.Users != nil && len(resp.Users) > 0 && resp.Users[0].UserEntity.UserId > 0 {
		d.SetId(strconv.Itoa(resp.Users[0].UserEntity.UserId))
	} else {
		return fmt.Errorf("unknown user %s", d.Get("name").(string))
	}

	return nil
}
