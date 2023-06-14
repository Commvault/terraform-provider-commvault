package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceKubernetesLabels() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadKubernetesLabels,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"clusterid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},
	}
}

func datasourceReadKubernetesLabels(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvKubernetesLabelsByName(d.Get("clusterid").(int), d.Get("namespace").(string), d.Get("name").(string))

	if len(resp.KubernetesGuid) > 0 {
		d.SetId(resp.KubernetesGuid)
	} else {
		return fmt.Errorf("unknown kubernetes label %s", d.Get("name").(string))
	}

	return nil
}
