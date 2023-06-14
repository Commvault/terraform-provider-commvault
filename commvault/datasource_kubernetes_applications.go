package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceKubernetesApplications() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadKubernetesApplications,

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

func datasourceReadKubernetesApplications(d *schema.ResourceData, m interface{}) error {
	resp, _ := handler.CvKubernetesApplicationsByName(d.Get("clusterid").(int), d.Get("namespace").(string), d.Get("name").(string))

	if len(resp.KubernetesGuid) > 0 {
		d.SetId(resp.KubernetesGuid)
	} else {
		return fmt.Errorf("unknown kubernetes application %s", d.Get("name").(string))
	}

	return nil
}
