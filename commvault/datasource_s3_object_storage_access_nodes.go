package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceS3ObjectStorageAccessNodes() *schema.Resource {
	return &schema.Resource{
		Read: datasourceReadS3ObjectStorageAccessNodes,
		Schema: map[string]*schema.Schema{
			"region_id":             {Type: schema.TypeInt, Optional: true, ValidateFunc: validateS3IntRange(1, 2147483647)},
			"exclude_client_groups": {Type: schema.TypeBool, Optional: true, Default: false},
			"exclude_clients":       {Type: schema.TypeBool, Optional: true, Default: false},
			"access_nodes": {
				Type: schema.TypeList, Computed: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"client_id":   {Type: schema.TypeInt, Computed: true},
					"client_name": {Type: schema.TypeString, Computed: true},
					"entity_type": {Type: schema.TypeInt, Computed: true},
				}},
			},
		},
	}
}

func datasourceReadS3ObjectStorageAccessNodes(d *schema.ResourceData, m interface{}) error {
	regionID := d.Get("region_id").(int)
	excludeGroups, excludeClients := d.Get("exclude_client_groups").(bool), d.Get("exclude_clients").(bool)
	resp, err := handler.CvGetS3AccessNodes(excludeGroups, excludeClients, regionID)
	if err != nil {
		return fmt.Errorf("list S3 access nodes: %w", err)
	}
	nodes := make([]map[string]interface{}, 0, len(resp.AccessNodes))
	for _, node := range resp.AccessNodes {
		if node.Entity == nil || node.Entity.ClientId == nil || node.Entity.ClientName == nil {
			return fmt.Errorf("S3 access-node response is missing entity, clientId, or clientName")
		}
		entry := map[string]interface{}{"client_id": *node.Entity.ClientId, "client_name": *node.Entity.ClientName}
		s3PutInt(entry, "entity_type", node.Entity.Type)
		nodes = append(nodes, entry)
	}
	if err := d.Set("access_nodes", nodes); err != nil {
		return err
	}
	d.SetId("AMAZON_S3:" + strconv.Itoa(regionID) + ":" + strconv.FormatBool(excludeGroups) + ":" + strconv.FormatBool(excludeClients))
	return nil
}
