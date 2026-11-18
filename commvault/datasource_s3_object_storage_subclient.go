package commvault

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceS3ObjectStorageSubclient() *schema.Resource {
	fields := s3ObjectStorageSubclientSchema()
	makeComputed(fields)
	fields["subclient_id"] = &schema.Schema{Type: schema.TypeInt, Required: true, ValidateFunc: validateS3IntRange(1, 2147483647)}
	return &schema.Resource{
		Schema: fields,
		Read: func(d *schema.ResourceData, m interface{}) error {
			d.SetId(strconv.Itoa(d.Get("subclient_id").(int)))
			return readS3ObjectStorageSubclient(d, true)
		},
	}
}
