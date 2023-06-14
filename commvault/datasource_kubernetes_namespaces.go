package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceKubernetesNamespaces() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadKubernetesNamespaces,

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
		},
	}
}

func datasourceReadKubernetesNamespaces(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvKubernetesNamespacesByName(d.Get("clusterid").(int), d.Get("name").(string))

	if len(resp.KubernetesGuid) > 0 {
		d.SetId(resp.KubernetesGuid)
	} else {
		return fmt.Errorf("unknown kubernetes namespace %s", d.Get("name").(string))
	}

	return nil
}
