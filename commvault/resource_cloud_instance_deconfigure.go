package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudInstanceDeconfigure() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateCloudInstanceDeconfigure,
		// This resource is a one-shot operation receipt, not ownership of the instance.
		Read:   func(d *schema.ResourceData, m interface{}) error { return nil },
		Delete: func(d *schema.ResourceData, m interface{}) error { d.SetId(""); return nil },
		Schema: map[string]*schema.Schema{
			"instance_id":   {Type: schema.TypeInt, Required: true, ForceNew: true, ValidateFunc: validateIntRange(1, 2147483647)},
			"client_name":   {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateNonEmptyString},
			"instance_name": {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateNonEmptyString},
			"workload": {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateRDSStringEnum("AMAZON_RDS", "AMAZON_S3"),
				Description: "Aurora snapshot registrations use AMAZON_RDS."},
			"confirm_deconfigure": {Type: schema.TypeBool, Required: true, ForceNew: true},
			"completed":           {Type: schema.TypeBool, Computed: true},
		},
		CustomizeDiff: func(d *schema.ResourceDiff, m interface{}) error {
			if !d.Get("confirm_deconfigure").(bool) {
				return fmt.Errorf("confirm_deconfigure must be true; deconfiguration disables protection")
			}
			return nil
		},
	}
}

func resourceCreateCloudInstanceDeconfigure(d *schema.ResourceData, m interface{}) error {
	if !d.Get("confirm_deconfigure").(bool) {
		return fmt.Errorf("confirm_deconfigure must be true")
	}
	id := strconv.Itoa(d.Get("instance_id").(int))
	state, workload, err := handler.CvGetCloudInstanceState(id)
	if err != nil {
		return fmt.Errorf("verify instance before deconfiguration: %w", err)
	}
	entity := state.Instance
	if entity == nil || entity.InstanceId == nil || strconv.Itoa(*entity.InstanceId) != id ||
		entity.ApplicationId == nil || *entity.ApplicationId != 134 || entity.ClientName == nil ||
		*entity.ClientName != d.Get("client_name").(string) || entity.InstanceName == nil ||
		*entity.InstanceName != d.Get("instance_name").(string) || workload != d.Get("workload").(string) {
		return fmt.Errorf("instance %s does not match the confirmed client, name and workload; refusing deconfiguration", id)
	}
	deconfigured := func(state *handler.MsgCloudInstanceState) bool {
		return (state.IsDeleted != nil && *state.IsDeleted) || (state.Status != nil && *state.Status&2 != 0)
	}
	if !deconfigured(state) {
		if err := handler.CvDeconfigureCloudInstance(id, handler.MsgDeleteS3ObjectStorageRequest{Instance: entity}); err != nil {
			return fmt.Errorf("deconfigure cloud instance %s: %w", id, err)
		}
	}
	d.SetId(id)
	state, _, err = handler.CvGetCloudInstanceState(id)
	if err != nil {
		return fmt.Errorf("deconfiguration succeeded but verification failed; do not retry blindly: %w", err)
	}
	if state.Instance == nil || state.Instance.InstanceId == nil || strconv.Itoa(*state.Instance.InstanceId) != id || !deconfigured(state) {
		return fmt.Errorf("deconfiguration accepted for %s but disabled status was not confirmed; inspect before retrying", id)
	}
	if err := d.Set("completed", true); err != nil {
		return err
	}
	return nil
}
