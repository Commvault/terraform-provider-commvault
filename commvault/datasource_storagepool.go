package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceStoragePool() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadStoragePool,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadStoragePool(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvStoragePoolIdByName(d.Get("name").(string))

	if resp.StoragePolicyEntity.StoragePolicyId > 0 {
		d.SetId(strconv.Itoa(resp.StoragePolicyEntity.StoragePolicyId))
	} else {
		return fmt.Errorf("unknown storage pool %s", d.Get("name").(string))
	}

	return nil
}
