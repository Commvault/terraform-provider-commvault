package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceClientGroup() *schema.Resource { 
	return &schema.Resource{
		Read: datasourceReadClientGroup,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadClientGroup(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvClientGroupIdByName(d.Get("name").(string))

	if resp.Groups != nil && len(resp.Groups) > 0 && resp.Groups[0].ClientGroup.ClientGroupId > 0 {
		d.SetId(strconv.Itoa(resp.Groups[0].ClientGroup.ClientGroupId))
	} else {
		return fmt.Errorf("unknown client group %s", d.Get("name").(string))
	}

	return nil
}
