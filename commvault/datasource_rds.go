package commvault

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceRDSInstance() *schema.Resource {
	fields := resourceRDSInstance().Schema
	// These creation-only values are not provided by the instance GET contract.
	delete(fields, "default_subclient_id")
	delete(fields, "use_resource_pool_info")
	makeComputed(fields)
	fields["instance_id"] = &schema.Schema{Type: schema.TypeInt, Required: true, ValidateFunc: validateIntRange(1, 2147483647)}
	return &schema.Resource{
		Schema: fields,
		Read: func(d *schema.ResourceData, m interface{}) error {
			d.SetId(strconv.Itoa(d.Get("instance_id").(int)))
			return readRDSInstance(d, true)
		},
	}
}

func datasourceRDSInstanceGroup() *schema.Resource {
	fields := resourceRDSInstanceGroup().Schema
	makeComputed(fields)
	fields["subclient_id"] = &schema.Schema{Type: schema.TypeInt, Required: true, ValidateFunc: validateIntRange(1, 2147483647)}
	return &schema.Resource{
		Schema: fields,
		Read: func(d *schema.ResourceData, m interface{}) error {
			d.SetId(strconv.Itoa(d.Get("subclient_id").(int)))
			return readRDSInstanceGroup(d, true)
		},
	}
}
