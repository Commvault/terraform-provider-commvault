package commvault

import (
	"fmt"
	"reflect"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func rdsSnapshotReplicaSchema() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeList, Optional: true, MaxItems: 1,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"copy_id": {Type: schema.TypeInt, Required: true, ValidateFunc: validateIntRange(1, 2147483647)},
			"region_mapping": {
				Type: schema.TypeSet, Required: true, MinItems: 1,
				Set: hashRDSRegionMapping,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"source_region":      {Type: schema.TypeString, Required: true, ValidateFunc: validateNonEmptyString},
					"destination_region": {Type: schema.TypeString, Required: true, ValidateFunc: validateNonEmptyString},
				}},
			},
		}},
	}
}

func hashRDSRegionMapping(value interface{}) int {
	mapping := value.(map[string]interface{})
	return schema.HashString(fmt.Sprintf("%q:%q", mapping["source_region"], mapping["destination_region"]))
}

func validateRDSGroupReplication(share, copy, account, replica interface{}) error {
	if share.(bool) && copy.(bool) {
		return fmt.Errorf("RDS share-only and full-copy cross-account modes are mutually exclusive")
	}
	if (share.(bool) || copy.(bool)) && len(account.([]interface{})) != 1 {
		return fmt.Errorf("RDS cross-account protection requires destination_account with a Commvault client_id")
	}
	for _, raw := range replica.([]interface{}) {
		item := raw.(map[string]interface{})
		if item["copy_id"].(int) <= 0 {
			return fmt.Errorf("snapshot_replica.copy_id must be positive")
		}
		mappings := item["region_mapping"].(*schema.Set)
		if mappings.Len() == 0 {
			return fmt.Errorf("snapshot_replica requires region_mapping; remove the block to clear overrides")
		}
		sources := make(map[string]bool)
		for _, rawMapping := range mappings.List() {
			mapping := rawMapping.(map[string]interface{})
			source := mapping["source_region"].(string)
			if sources[source] {
				return fmt.Errorf("RDS source region %q has multiple destination mappings", source)
			}
			sources[source] = true
		}
	}
	return nil
}

func customizeRDSGroupReplication(d *schema.ResourceDiff, m interface{}) error {
	if err := validateRDSReplicaPlanChange(d); err != nil {
		return err
	}
	share := d.NewValueKnown("is_cross_account_enabled") && d.Get("is_cross_account_enabled").(bool)
	copy := d.NewValueKnown("is_cross_account_copy_enabled") && d.Get("is_cross_account_copy_enabled").(bool)
	if share && copy {
		return fmt.Errorf("RDS share-only and full-copy cross-account modes are mutually exclusive")
	}
	if d.NewValueKnown("destination_account") && (share || copy) && len(d.Get("destination_account").([]interface{})) == 0 {
		return fmt.Errorf("RDS cross-account protection requires destination_account with a Commvault client_id")
	}
	if d.NewValueKnown("snapshot_replica") && d.NewValueKnown("snapshot_replica.0.copy_id") {
		return validateRDSGroupReplication(false, false, []interface{}{}, d.Get("snapshot_replica"))
	}
	return nil
}

func validateRDSReplicaPlanChange(d interface {
	HasChange(string) bool
	GetChange(string) (interface{}, interface{})
	Get(string) interface{}
}) error {
	oldPlan, _ := d.GetChange("plan_name")
	if d.HasChange("plan_name") && oldPlan != nil && oldPlan.(string) != "" {
		oldReplica, _ := d.GetChange("snapshot_replica")
		if len(oldReplica.([]interface{})) != 0 || len(d.Get("snapshot_replica").([]interface{})) != 0 {
			return fmt.Errorf("remove snapshot_replica and apply before changing plan_name; configure the new plan's replica in a subsequent apply")
		}
	}
	return nil
}

func buildRDSDestinationAccount(d *schema.ResourceData) *handler.MsgRDSCrossAccountDetails {
	accounts := d.Get("destination_account").([]interface{})
	if len(accounts) == 0 {
		return nil
	}
	id := accounts[0].(map[string]interface{})["client_id"].(int)
	return &handler.MsgRDSCrossAccountDetails{AmazonClient: &handler.MsgRDSAccountClient{ClientId: &id}}
}

func verifyRDSCrossAccount(d *schema.ResourceData, expected *handler.MsgRDSSubclient) error {
	for key, value := range map[string]*bool{
		"is_cross_account_enabled":      expected.IsCrossAccountEnabled,
		"is_cross_account_copy_enabled": expected.IsCrossAccountCopyEnabled,
	} {
		if value != nil && d.Get(key).(bool) != *value {
			return fmt.Errorf("RDS group %s did not persist %s", d.Id(), key)
		}
	}
	if expected.DestinationAccount != nil {
		actual := buildRDSDestinationAccount(d)
		if actual == nil || *actual.AmazonClient.ClientId != *expected.DestinationAccount.AmazonClient.ClientId {
			return fmt.Errorf("RDS group %s did not persist destination_account.client_id", d.Id())
		}
	}
	return nil
}

func validateRDSReplicaCopy(prop handler.MsgRDSReadSubclientProperties, copyID int) error {
	if prop.PlanEntity == nil {
		return fmt.Errorf("cannot verify RDS snapshot replica without an assigned plan")
	}
	planID := 0
	if prop.PlanEntity.PlanId != nil {
		planID = *prop.PlanEntity.PlanId
	}
	if planID <= 0 {
		name := prop.PlanEntity.PlanName
		if name == nil {
			name = prop.PlanEntity.Name
		}
		if name == nil || *name == "" {
			return fmt.Errorf("RDS group assigned plan omits both ID and name")
		}
		plans, err := handler.CvPlanIdByName(*name)
		if err != nil {
			return fmt.Errorf("resolve RDS replica plan: %w", err)
		}
		if plans == nil || len(plans.Plans) != 1 || plans.Plans[0].Plan.PlanId <= 0 {
			return fmt.Errorf("cannot uniquely resolve RDS replica plan %q", *name)
		}
		planID = plans.Plans[0].Plan.PlanId
	}
	detail, err := handler.CvGetServerPlanBackupDestinationDetails(strconv.Itoa(planID), strconv.Itoa(copyID))
	if err != nil {
		return fmt.Errorf("verify RDS replica copy: %w", err)
	}
	if detail == nil || detail.Plan == nil || detail.Plan.Id == nil || *detail.Plan.Id != planID ||
		detail.PlanBackupDestination == nil || detail.PlanBackupDestination.Id == nil || *detail.PlanBackupDestination.Id != copyID {
		return fmt.Errorf("copy %d does not identify a destination of the RDS group's assigned plan %d", copyID, planID)
	}
	general := detail.General
	if general == nil || general.IsSnapCopy == nil || !*general.IsSnapCopy ||
		(general.IsMirrorCopy != nil && *general.IsMirrorCopy) ||
		(general.SnapCopyType != nil && *general.SnapCopyType == "SNAPSHOT_PRIMARY") ||
		general.Source == nil || general.Source.Id == nil || *general.Source.Id <= 0 {
		return fmt.Errorf("copy %d must be a secondary snapshot replica, not a disk, primary snapshot or mirror copy", copyID)
	}
	return nil
}

func rdsRegionMapValues(mappings []handler.MsgRDSReadRegionMapping) (map[string]string, error) {
	values := make(map[string]string)
	for _, mapping := range mappings {
		if mapping.SourceRegionName == nil || *mapping.SourceRegionName == "" ||
			mapping.DestinationRegionName == nil || *mapping.DestinationRegionName == "" {
			return nil, fmt.Errorf("RDS region mapping response omits source or destination region names")
		}
		source := *mapping.SourceRegionName
		if _, exists := values[source]; exists {
			return nil, fmt.Errorf("RDS region mapping response repeats source region %q", source)
		}
		values[source] = *mapping.DestinationRegionName
	}
	return values, nil
}

func rdsConfiguredRegionMaps(block map[string]interface{}) map[string]string {
	values := make(map[string]string)
	for _, raw := range block["region_mapping"].(*schema.Set).List() {
		mapping := raw.(map[string]interface{})
		values[mapping["source_region"].(string)] = mapping["destination_region"].(string)
	}
	return values
}

func readRDSReplicaMappings(d *schema.ResourceData, prop handler.MsgRDSReadSubclientProperties, values map[string]interface{}) error {
	replicas := d.Get("snapshot_replica").([]interface{})
	if len(replicas) == 0 {
		return nil
	}
	copyID := replicas[0].(map[string]interface{})["copy_id"].(int)
	if err := validateRDSReplicaCopy(prop, copyID); err != nil {
		return err
	}
	mappings, err := handler.CvGetRDSRegionMappings(strconv.Itoa(copyID), d.Id())
	if err != nil {
		return fmt.Errorf("read RDS group replica mappings: %w", err)
	}
	actual, err := rdsRegionMapValues(mappings)
	if err != nil {
		return err
	}
	items := make([]interface{}, 0, len(actual))
	for source, target := range actual {
		items = append(items, map[string]interface{}{"source_region": source, "destination_region": target})
	}
	values["snapshot_replica"] = []interface{}{map[string]interface{}{"copy_id": copyID, "region_mapping": items}}
	return nil
}

func syncRDSReplicaMappings(d *schema.ResourceData, creating bool) error {
	if !creating && !d.HasChange("snapshot_replica") {
		return nil
	}
	oldRaw, newRaw := d.GetChange("snapshot_replica")
	oldBlocks, newBlocks := oldRaw.([]interface{}), newRaw.([]interface{})
	if creating {
		oldBlocks = nil
		newBlocks = d.Get("snapshot_replica").([]interface{})
	}
	if len(oldBlocks) == 0 && len(newBlocks) == 0 {
		return nil
	}
	properties, err := handler.CvGetRDSInstanceGroup(d.Id())
	if err != nil {
		return fmt.Errorf("verify RDS group before mapping update: %w", err)
	}
	if len(properties) != 1 || properties[0].SubclientEntity == nil ||
		properties[0].SubclientEntity.SubclientId == nil || strconv.Itoa(*properties[0].SubclientEntity.SubclientId) != d.Id() ||
		properties[0].SubclientEntity.ApplicationId == nil || *properties[0].SubclientEntity.ApplicationId != 134 ||
		properties[0].CloudAppsSubclientProp == nil || properties[0].CloudAppsSubclientProp.InstanceType == nil ||
		*properties[0].CloudAppsSubclientProp.InstanceType != "AMAZON_RDS" {
		return fmt.Errorf("cannot verify AMAZON_RDS subclient %s before mapping update", d.Id())
	}
	desired := make(map[int]map[string]string)
	copyIDs := make([]int, 0, 2)
	if len(oldBlocks) != 0 {
		copyID := oldBlocks[0].(map[string]interface{})["copy_id"].(int)
		desired[copyID] = map[string]string{}
		copyIDs = append(copyIDs, copyID)
	}
	if len(newBlocks) != 0 {
		block := newBlocks[0].(map[string]interface{})
		copyID := block["copy_id"].(int)
		if _, exists := desired[copyID]; !exists {
			copyIDs = append(copyIDs, copyID)
		}
		desired[copyID] = rdsConfiguredRegionMaps(block)
	}
	type pendingUpdate struct {
		copyID   string
		request  handler.MsgRDSRegionMapRequest
		expected map[string]string
	}
	updates := make([]pendingUpdate, 0, len(desired))
	for _, copyID := range copyIDs {
		target := desired[copyID]
		if err := validateRDSReplicaCopy(properties[0], copyID); err != nil {
			return err
		}
		id := strconv.Itoa(copyID)
		current, err := handler.CvGetRDSRegionMappings(id, d.Id())
		if err != nil {
			return err
		}
		actual, err := rdsRegionMapValues(current)
		if err != nil {
			return err
		}
		request := handler.MsgRDSRegionMapRequest{}
		subclientID := *properties[0].SubclientEntity.SubclientId
		appendMapping := func(source, destination string) {
			request.RegionMap = append(request.RegionMap, handler.MsgRDSRegionMapping{
				SourceRegion: &source, DestinationRegion: &destination,
				SubClient: &handler.MsgRDSSubclientEntity{SubclientId: &subclientID},
			})
		}
		for source := range actual {
			if _, retained := target[source]; !retained {
				appendMapping(source, "")
			}
		}
		// Always write desired rows, even if GET returned identical inherited plan mappings.
		for source, destination := range target {
			appendMapping(source, destination)
		}
		expected := target
		if len(target) == 0 {
			defaults, err := handler.CvGetRDSRegionMappings(id, "")
			if err != nil {
				return err
			}
			expected, err = rdsRegionMapValues(defaults)
			if err != nil {
				return err
			}
		}
		updates = append(updates, pendingUpdate{id, request, expected})
	}
	for _, update := range updates {
		if len(update.request.RegionMap) != 0 {
			if err := handler.CvSetRDSRegionMappings(update.copyID, update.request); err != nil {
				return fmt.Errorf("update RDS replica %s mappings: %w", update.copyID, err)
			}
		}
		actualRows, err := handler.CvGetRDSRegionMappings(update.copyID, d.Id())
		if err != nil {
			return err
		}
		actual, err := rdsRegionMapValues(actualRows)
		if err != nil {
			return err
		}
		if !reflect.DeepEqual(actual, update.expected) {
			return fmt.Errorf("RDS replica %s did not persist the requested effective region mappings", update.copyID)
		}
	}
	return nil
}
