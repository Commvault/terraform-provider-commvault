package commvault

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func planDestinationErrorMessage(message *string) string {
	if message == nil {
		return "(no error message returned)"
	}
	return *message
}

func validateServerPlanDestinationMappings(destination handler.MsgCreatePlanBackupDestinationSet) error {
	for _, mapping := range destination.Mappings {
		if mapping.Vendor == nil || (*mapping.Vendor != "AMAZON" && *mapping.Vendor != "AZURE") {
			continue
		}
		if mapping.Source == nil || mapping.Source.Name == nil || strings.TrimSpace(*mapping.Source.Name) == "" ||
			mapping.Target == nil || mapping.Target.Name == nil || strings.TrimSpace(*mapping.Target.Name) == "" {
			return fmt.Errorf("%s snapshot mappings require non-empty source.name and target.name cloud region names; numeric region IDs alone are not supported", *mapping.Vendor)
		}
	}
	return nil
}

func validateServerPlanDestinationUpdateFields(fields map[string]interface{}) error {
	for _, key := range []string{"isconfiguredforreplication", "backupstarttime", "optimizeforinstantclone", "storagepolicy"} {
		value := fields[key]
		if value != nil && !reflect.ValueOf(value).IsZero() {
			if list, ok := value.([]interface{}); ok && len(list) == 0 {
				continue
			}
			return fmt.Errorf("backup destination %q: %s is not supported by the server-plan destination update contract", fields["backupdestinationname"], key)
		}
	}
	return nil
}

func serverPlanDestinationsByName(items []interface{}) (map[string]map[string]interface{}, error) {
	result := make(map[string]map[string]interface{}, len(items))
	for _, item := range items {
		fields := item.(map[string]interface{})
		name, _ := fields["backupdestinationname"].(string)
		if strings.TrimSpace(name) == "" {
			return nil, fmt.Errorf("backupdestinations requires a non-empty backupdestinationname for each destination")
		}
		if _, exists := result[name]; exists {
			return nil, fmt.Errorf("duplicate backupdestinationname %q", name)
		}
		result[name] = fields
	}
	return result, nil
}

func hashServerPlanDestination(value interface{}) int {
	fields := resourcePlan_Server().Schema["backupdestinations"].Elem.(*schema.Resource).Schema
	stable := make(map[string]*schema.Schema, len(fields))
	for key, field := range fields {
		if !field.Computed {
			stable[key] = field
		}
	}
	// Optional/computed defaults must not change a set element's identity during refresh.
	return schema.HashResource(&schema.Resource{Schema: stable})(value)
}

func serverPlanDestinationEqual(a, b interface{}) bool {
	resource := resourcePlan_Server().Schema["backupdestinations"].Elem.(*schema.Resource)
	hash := schema.HashResource(resource)
	return hash(a) == hash(b)
}

func customizeServerPlanDestinations(d *schema.ResourceDiff, m interface{}) error {
	if d.Id() == "" || !d.NewValueKnown("backupdestinations") {
		return nil
	}
	destinations := d.Get("backupdestinations").(*schema.Set)
	oldValue, _ := d.GetChange("backupdestinations")
	oldNames, err := serverPlanDestinationsByName(oldValue.(*schema.Set).List())
	if err != nil {
		return err
	}
	for _, item := range destinations.List() {
		fields := item.(map[string]interface{})
		name := fields["backupdestinationname"].(string)
		old, exists := oldNames[name]
		if !exists {
			continue
		}
		for _, key := range []string{"isdefault", "copytype", "copyprecedence"} {
			path := fmt.Sprintf("backupdestinations.%d.%s", destinations.F(item), key)
			next, specified := d.GetOkExists(path)
			if specified && d.NewValueKnown(path) && !reflect.DeepEqual(old[key], next) {
				return fmt.Errorf("changing %s on existing backup destination %q is not supported by the current provider SDK; configure this setting when creating or adding the destination", key, name)
			}
		}
	}
	return nil
}

func buildServerPlanDestinationUpdate(d *schema.ResourceData) ([]handler.MsgPlanBackupDestinationDetailsSet, error) {
	oldValue, newValue := d.GetChange("backupdestinations")
	oldSet, newSet := oldValue.(*schema.Set), newValue.(*schema.Set)
	oldNames, err := serverPlanDestinationsByName(oldSet.List())
	if err != nil {
		return nil, err
	}
	newNames, err := serverPlanDestinationsByName(newSet.List())
	if err != nil {
		return nil, err
	}
	operation := d.Get("backupdestinationsoperationtype").(string)
	var selected []interface{}
	switch operation {
	case "ADD":
		for name, old := range oldNames {
			if next, exists := newNames[name]; !exists || !serverPlanDestinationEqual(old, next) {
				return nil, fmt.Errorf("ADD cannot modify or remove existing backup destination %q; retain its configuration", name)
			}
		}
		for _, item := range newSet.List() {
			if _, exists := oldNames[item.(map[string]interface{})["backupdestinationname"].(string)]; !exists {
				selected = append(selected, item)
			}
		}
	case "MODIFY":
		if len(oldNames) != len(newNames) {
			return nil, fmt.Errorf("MODIFY cannot add or remove backup destinations")
		}
		for name := range oldNames {
			if _, exists := newNames[name]; !exists {
				return nil, fmt.Errorf("MODIFY cannot rename or remove backup destination %q", name)
			}
		}
		for _, item := range newSet.List() {
			name := item.(map[string]interface{})["backupdestinationname"].(string)
			if !serverPlanDestinationEqual(oldNames[name], item) {
				selected = append(selected, item)
			}
		}
	case "DELETE":
		for name, next := range newNames {
			if old, exists := oldNames[name]; !exists || !serverPlanDestinationEqual(old, next) {
				return nil, fmt.Errorf("DELETE cannot add or modify backup destination %q", name)
			}
		}
		selected = oldSet.Difference(newSet).List()
	case "OVERWRITE":
		selected = newSet.List()
	default:
		return nil, fmt.Errorf("set backupdestinationsoperationtype to ADD, MODIFY, DELETE, or OVERWRITE when changing an existing plan's destinations")
	}
	return buildServerPlanSelectedDestinationUpdate(d, operation, selected)
}

func buildServerPlanSelectedDestinationUpdate(d *schema.ResourceData, operation string, selected []interface{}) ([]handler.MsgPlanBackupDestinationDetailsSet, error) {
	current, err := handler.CvGetPlanById(d.Id())
	if err != nil {
		return nil, fmt.Errorf("resolve server plan backup destinations before %s: %w", operation, err)
	}
	if current.Plan == nil {
		return nil, fmt.Errorf("resolve server plan %s: response omitted plan identity", d.Id())
	}
	if current.BackupDestinations == nil && d.Get("backupdestinations").(*schema.Set).Len() > 0 {
		return nil, fmt.Errorf("resolve server plan %s: response omitted backupDestinations", d.Id())
	}
	identities := make(map[string]*handler.MsgIdName)
	for _, destination := range current.BackupDestinations {
		if destination.PlanBackupDestination == nil || destination.PlanBackupDestination.Name == nil {
			return nil, fmt.Errorf("server plan %s returned a backup destination without a name", d.Id())
		}
		if _, exists := identities[*destination.PlanBackupDestination.Name]; exists {
			return nil, fmt.Errorf("server plan %s returned duplicate backup destination %q", d.Id(), *destination.PlanBackupDestination.Name)
		}
		identities[*destination.PlanBackupDestination.Name] = destination.PlanBackupDestination
	}
	result := make([]handler.MsgPlanBackupDestinationDetailsSet, 0, len(selected))
	for _, item := range selected {
		fields := item.(map[string]interface{})
		name := fields["backupdestinationname"].(string)
		identity, exists := identities[name]
		if operation == "ADD" && exists {
			return nil, fmt.Errorf("backup destination %q already exists on server plan %s; refresh before retrying ADD", name, d.Id())
		}
		if operation == "MODIFY" || operation == "DELETE" {
			if !exists || identity.Id == nil || *identity.Id <= 0 {
				return nil, fmt.Errorf("cannot resolve backup destination %q on server plan %s", name, d.Id())
			}
		}
		if operation == "DELETE" {
			result = append(result, handler.MsgPlanBackupDestinationDetailsSet{PlanBackupDestination: identity})
			continue
		}
		// These creation fields have no equivalent in the current modify-plan contract.
		if err := validateServerPlanDestinationUpdateFields(fields); err != nil {
			return nil, err
		}
		created := build_plan_server_msgcreateplanbackupdestinationset_array(d, []interface{}{item})[0]
		if err := validateServerPlanDestinationMappings(created); err != nil {
			return nil, err
		}
		if source := created.SourceCopy; source != nil && source.Name != nil && strings.TrimSpace(*source.Name) != "" && (source.Id == nil || *source.Id == 0) {
			resolved, exists := identities[*source.Name]
			if !exists || resolved.Id == nil || *resolved.Id <= 0 {
				return nil, fmt.Errorf("cannot resolve source copy %q for backup destination %q on server plan %s", *source.Name, name, d.Id())
			}
			created.SourceCopy = &handler.MsgIdName{Id: resolved.Id, Name: source.Name}
		}
		destination := serverPlanDestinationUpdateMessage(created)
		if exists {
			destination.PlanBackupDestination.Id = identity.Id
		}
		result = append(result, destination)
	}
	return result, nil
}

func serverPlanDestinationUpdateMessage(created handler.MsgCreatePlanBackupDestinationSet) handler.MsgPlanBackupDestinationDetailsSet {
	var storage *handler.MsgStoragePool
	if created.StoragePool != nil {
		storage = &handler.MsgStoragePool{Id: created.StoragePool.Id, Name: created.StoragePool.Name}
	}
	var region *handler.MsgIdNameDisplayName
	if created.Region != nil {
		region = &handler.MsgIdNameDisplayName{Id: created.Region.Id, Name: created.Region.Name}
	}
	return handler.MsgPlanBackupDestinationDetailsSet{
		PlanBackupDestination: &handler.MsgIdName{Name: created.BackupDestinationName},
		General: &handler.MsgPlanBackupDestinationGeneralInfo{
			IsMirrorCopy: created.IsMirrorCopy, IsSnapCopy: created.IsSnapCopy,
			Storage: storage, Source: created.SourceCopy, NetAppCloudTarget: created.NetAppCloudTarget,
			IsDefault: created.IsDefault, CopyType: created.CopyType, CopyPrecedence: created.CopyPrecedence,
			StorageType: created.StorageType, StorageTemplateTags: created.StorageTemplateTags,
		},
		RetentionRules: &handler.MsgPlanBackupDestinationRetentionRuleInfo{
			RetentionPeriodDays: created.RetentionPeriodDays, ExtendedRetentionRules: created.ExtendedRetentionRules,
			RetentionRuleType: created.RetentionRuleType, SnapRecoveryPoints: created.SnapRecoveryPoints,
			UseExtendedRetentionRules: created.UseExtendedRetentionRules,
			OverrideRetentionSettings: created.OverrideRetentionSettings,
		},
		BackupType: &handler.MsgPlanBackupDestinationBackupTypeInfo{
			BackupTypesToCopy: created.BackupsToCopy, FullBackupTypesToCopy: created.FullBackupTypesToCopy,
		},
		Mappings: created.Mappings, Region: region,
	}
}

func flattenServerPlanDestinations(d *schema.ResourceData, destinations []handler.MsgPlanBackupDestinationSet) ([]interface{}, error) {
	configured, err := serverPlanDestinationsByName(d.Get("backupdestinations").(*schema.Set).List())
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, 0, len(destinations))
	seen := make(map[string]bool)
	for _, destination := range destinations {
		identity := destination.PlanBackupDestination
		if identity == nil || identity.Name == nil || strings.TrimSpace(*identity.Name) == "" {
			return nil, fmt.Errorf("server plan %s returned a backup destination without a name", d.Id())
		}
		name := *identity.Name
		if seen[name] {
			return nil, fmt.Errorf("server plan %s returned duplicate backup destination %q", d.Id(), name)
		}
		seen[name] = true
		// Optimized V4 plan reads omit mappings; secondary snapshot copies need their detail response.
		if destination.Mappings == nil && destination.IsSnapCopy != nil && *destination.IsSnapCopy && destination.SourceCopy != nil {
			if identity.Id == nil || *identity.Id <= 0 {
				return nil, fmt.Errorf("server plan %s returned snapshot destination %q without an ID", d.Id(), name)
			}
			details, err := handler.CvGetServerPlanBackupDestinationDetails(d.Id(), strconv.Itoa(*identity.Id))
			if err != nil {
				return nil, fmt.Errorf("read snapshot mappings for backup destination %q: %w", name, err)
			}
			if details.Plan == nil || details.Plan.Id == nil || strconv.Itoa(*details.Plan.Id) != d.Id() ||
				details.PlanBackupDestination == nil || details.PlanBackupDestination.Id == nil ||
				*details.PlanBackupDestination.Id != *identity.Id {
				return nil, fmt.Errorf("snapshot mapping details for backup destination %q omitted or mismatched the plan/copy identity", name)
			}
			destination.Mappings = details.Mappings
		}
		fields := map[string]interface{}{
			"backupdestinationname": name,
		}
		putBool := func(key string, value *bool) {
			fields[key] = nil
			if value != nil {
				fields[key] = fmt.Sprintf("%t", *value)
			}
		}
		putInt := func(key string, value *int) {
			fields[key] = nil
			if value != nil {
				fields[key] = *value
			}
		}
		putString := func(key string, value *string) {
			fields[key] = nil
			if value != nil {
				fields[key] = *value
			}
		}
		putBool("ismirrorcopy", destination.IsMirrorCopy)
		putBool("issnapcopy", destination.IsSnapCopy)
		if destination.IsMirrorCopy == nil {
			fields["ismirrorcopy"] = "false"
		}
		if destination.IsSnapCopy == nil {
			fields["issnapcopy"] = "false"
		}
		putBool("isconfiguredforreplication", destination.IsConfiguredForReplication)
		putBool("useextendedretentionrules", destination.UseExtendedRetentionRules)
		putBool("overrideretentionsettings", destination.OverrideRetentionSettings)
		putBool("netappcloudtarget", destination.NetAppCloudTarget)
		putInt("retentionperioddays", destination.RetentionPeriodDays)
		putInt("snaprecoverypoints", destination.SnapRecoveryPoints)
		putInt("backupstarttime", destination.BackupStartTime)
		putInt("copyprecedence", destination.CopyPrecedence)
		putString("backupstocopy", destination.BackupsToCopy)
		putString("retentionruletype", destination.RetentionRuleType)
		putString("fullbackuptypestocopy", destination.FullBackupTypesToCopy)
		putString("copytype", destination.CopyType)
		putString("storagetype", destination.StorageType)
		fields["isdefault"] = destination.IsDefault
		if destination.IsDefault != nil {
			fields["isdefault"] = *destination.IsDefault
		}
		fields["storagepool"] = []map[string]interface{}{}
		if destination.StoragePool != nil {
			if value, ok := serialize_plan_server_msgidname(d, &handler.MsgIdName{Id: destination.StoragePool.Id, Name: destination.StoragePool.Name}); ok {
				fields["storagepool"] = value
			}
		}
		fields["sourcecopy"], _ = serialize_plan_server_msgidname(d, destination.SourceCopy)
		fields["region"], _ = serialize_plan_server_msgidname(d, destination.Region)
		fields["extendedretentionrules"], _ = serialize_plan_backupdestination_msgextendedretentionrules(d, destination.ExtendedRetentionRules)
		fields["mappings"], _ = serialize_plan_backupdestination_msgsnapshotcopymappingset_array(d, destination.Mappings)
		fields["storagetemplatetags"] = flattenServerPlanStorageTags(destination.StorageTemplateTags)
		if previous, exists := configured[name]; exists {
			fields = projectServerPlanDestination(fields, previous)
		}
		if destination.IsDefault != nil {
			fields["isdefault"] = *destination.IsDefault
		}
		if destination.CopyType != nil {
			fields["copytype"] = *destination.CopyType
		}
		if destination.CopyPrecedence != nil {
			fields["copyprecedence"] = *destination.CopyPrecedence
		}
		result = append(result, fields)
	}
	return result, nil
}

// Keep optional server defaults out of configuration while refreshing configured values.
func projectServerPlanDestination(server, configured map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(configured))
	for key, previous := range configured {
		if value, exists := server[key]; exists {
			if previous == nil || reflect.ValueOf(previous).IsZero() {
				continue
			}
			result[key] = projectServerPlanDestinationValue(value, previous)
		} else {
			result[key] = previous
		}
	}
	result["backupdestinationname"] = server["backupdestinationname"]
	return result
}

func projectServerPlanDestinationValue(server, configured interface{}) interface{} {
	if current, ok := server.([]map[string]interface{}); ok {
		values := make([]interface{}, len(current))
		for i := range current {
			values[i] = current[i]
		}
		server = values
	}
	if current, ok := server.([]interface{}); ok {
		var previous []interface{}
		switch value := configured.(type) {
		case []interface{}:
			previous = value
		case *schema.Set:
			previous = value.List()
		}
		values := make([]interface{}, len(current))
		for i, item := range current {
			values[i] = item
			if i < len(previous) {
				matched := previous[i]
				for _, candidate := range previous {
					if serverPlanDestinationIdentityMatches(item, candidate) {
						matched = candidate
						break
					}
				}
				values[i] = projectServerPlanDestinationValue(item, matched)
			}
		}
		return values
	}
	if current, ok := server.(map[string]interface{}); ok {
		if previous, ok := configured.(map[string]interface{}); ok {
			values := make(map[string]interface{})
			for key, old := range previous {
				if old != nil && !reflect.ValueOf(old).IsZero() {
					values[key] = projectServerPlanDestinationValue(current[key], old)
				}
			}
			return values
		}
	}
	return server
}

func serverPlanDestinationIdentityMatches(server, configured interface{}) bool {
	if value := reflect.ValueOf(server); value.IsValid() && value.Kind() == reflect.Ptr && !value.IsNil() {
		server = value.Elem().Interface()
	}
	switch previous := configured.(type) {
	case map[string]interface{}:
		current, ok := server.(map[string]interface{})
		if !ok {
			return false
		}
		for _, key := range []string{"id", "name", "vendor", "source", "target", "sourcevendor", "targetvendor"} {
			if value := previous[key]; value != nil && !reflect.ValueOf(value).IsZero() {
				if !serverPlanDestinationIdentityMatches(current[key], value) {
					return false
				}
			}
		}
		return true
	case []interface{}:
		if len(previous) == 0 {
			return true
		}
		switch current := server.(type) {
		case []interface{}:
			return len(current) > 0 && serverPlanDestinationIdentityMatches(current[0], previous[0])
		case []map[string]interface{}:
			return len(current) > 0 && serverPlanDestinationIdentityMatches(current[0], previous[0])
		}
		return false
	default:
		return reflect.DeepEqual(server, configured)
	}
}

func flattenServerPlanStorageTags(tags []handler.MsgIdNameValueSet) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		fields := make(map[string]interface{})
		if tag.Id != nil {
			fields["id"] = *tag.Id
		}
		if tag.Name != nil {
			fields["name"] = *tag.Name
		}
		if tag.Value != nil {
			fields["value"] = *tag.Value
		}
		result = append(result, fields)
	}
	return result
}

type serverPlanDestinationExpectation struct {
	destinations *schema.Set
	unknown      map[string][]string
	configured   bool
}

func expectedServerPlanDestinations(d *schema.ResourceData) serverPlanDestinationExpectation {
	destinations := d.Get("backupdestinations").(*schema.Set)
	expected := serverPlanDestinationExpectation{destinations: destinations, unknown: make(map[string][]string)}
	expected.configured = destinations.Len() > 0 ||
		(d.HasChange("backupdestinations") && (d.Get("backupdestinationsoperationtype") == "DELETE" || d.Get("backupdestinationsoperationtype") == "OVERWRITE"))
	for _, item := range destinations.List() {
		fields := item.(map[string]interface{})
		name := fields["backupdestinationname"].(string)
		for _, key := range []string{"isdefault", "copytype", "copyprecedence"} {
			path := fmt.Sprintf("backupdestinations.%d.%s", destinations.F(item), key)
			if _, exists := d.GetOkExists(path); !exists {
				expected.unknown[name] = append(expected.unknown[name], key)
			}
		}
	}
	return expected
}

func verifyServerPlanDestinations(d *schema.ResourceData, expected serverPlanDestinationExpectation) error {
	actual := d.Get("backupdestinations").(*schema.Set)
	if expected.configured && actual.Len() != expected.destinations.Len() {
		return fmt.Errorf("server plan %s did not confirm the configured destination list after the write; refresh and inspect the plan before retrying", d.Id())
	}
	names, err := serverPlanDestinationsByName(actual.List())
	if err != nil {
		return err
	}
	for _, item := range expected.destinations.List() {
		fields := item.(map[string]interface{})
		name := fields["backupdestinationname"].(string)
		server, exists := names[name]
		if !exists {
			return fmt.Errorf("server plan %s did not confirm backup destination %q after the write; refresh and inspect the plan before retrying", d.Id(), name)
		}
		comparable := make(map[string]interface{}, len(fields))
		for key, value := range fields {
			comparable[key] = value
		}
		for _, key := range expected.unknown[name] {
			comparable[key] = server[key]
		}
		if !serverPlanDestinationEqual(server, comparable) {
			return fmt.Errorf("server plan %s did not confirm the configured settings for backup destination %q after the write; refresh and inspect the plan before retrying", d.Id(), name)
		}
	}
	return nil
}
