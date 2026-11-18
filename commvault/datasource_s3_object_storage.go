package commvault

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceS3ObjectStorage() *schema.Resource {
	fields := s3ObjectStorageSchema()
	// Deconfiguration metadata is not returned by the instance GET endpoint.
	delete(fields, "company_name")
	delete(fields, "multi_commcell_id")
	makeComputed(fields)
	fields["instance_id"] = &schema.Schema{Type: schema.TypeInt, Required: true, ValidateFunc: validateS3IntRange(1, 2147483647)}
	return &schema.Resource{
		Schema: fields,
		Read: func(d *schema.ResourceData, m interface{}) error {
			d.SetId(strconv.Itoa(d.Get("instance_id").(int)))
			return readS3ObjectStorage(d, true)
		},
	}
}

func makeComputed(fields map[string]*schema.Schema) {
	for _, field := range fields {
		field.Required, field.Optional, field.ForceNew = false, false, false
		field.Computed = true
		field.Default, field.DefaultFunc, field.ValidateFunc = nil, nil, nil
		field.MinItems, field.MaxItems = 0, 0
		if nested, ok := field.Elem.(*schema.Resource); ok {
			makeComputed(nested.Schema)
		}
	}
}
