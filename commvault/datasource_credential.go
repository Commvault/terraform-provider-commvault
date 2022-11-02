package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceCredential() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadCredential,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadCredential(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvCredentialByName(d.Get("name").(string))

	if resp.Id > 0 {
		d.SetId(strconv.Itoa(resp.Id))
	} else {
		return fmt.Errorf("unknown credential %s", d.Get("name").(string))
	}

	return nil
}
