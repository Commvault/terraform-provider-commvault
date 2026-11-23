package commvault

import (
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceS3ObjectStorageSubclient() *schema.Resource {
	return &schema.Resource{
		Create:   resourceCreateS3ObjectStorageSubclient,
		Read:     resourceReadS3ObjectStorageSubclient,
		Update:   resourceUpdateS3ObjectStorageSubclient,
		Delete:   resourceDeleteS3ObjectStorageSubclient,
		Importer: &schema.ResourceImporter{State: schema.ImportStatePassthrough},
		Schema:   s3ObjectStorageSubclientSchema(),
	}
}

func s3ObjectStorageSubclientSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"client_id":                {Type: schema.TypeInt, Required: true, ForceNew: true, ValidateFunc: validateS3IntRange(1, 2147483647)},
		"instance_id":              {Type: schema.TypeInt, Required: true, ForceNew: true, ValidateFunc: validateS3IntRange(1, 2147483647)},
		"subclient_id":             {Type: schema.TypeInt, Computed: true},
		"subclient_name":           {Type: schema.TypeString, Required: true, ValidateFunc: validateS3NonEmptyString},
		"plan_id":                  {Type: schema.TypeInt, Required: true, ValidateFunc: validateS3IntRange(1, 2147483647)},
		"plan_name":                {Type: schema.TypeString, Computed: true},
		"number_of_backup_streams": {Type: schema.TypeInt, Optional: true, Default: 4, ValidateFunc: validateS3IntRange(1, 2147483647)},
		"job_start_time":           {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateS3IntRange(0, 86399)},
		"collect_acl":              {Type: schema.TypeBool, Optional: true, Default: false},
		"event_scan_enabled":       {Type: schema.TypeBool, Optional: true, Computed: true},
		"event_scan_staging_path":  {Type: schema.TypeString, Optional: true, Computed: true},
		"backup_content_type":      {Type: schema.TypeInt, Computed: true, Description: "0 for path-based content, 1 for tag-based content."},
		"content": {
			Type: schema.TypeList, Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"path":         {Type: schema.TypeString, Optional: true},
				"exclude_path": {Type: schema.TypeString, Optional: true},
				"include_path": {Type: schema.TypeString, Optional: true},
			}},
		},
		"tag_filter": {
			Type: schema.TypeList, Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"rule_item_type": {Type: schema.TypeInt, Optional: true, Default: 2, ValidateFunc: validateS3IntRange(2, 2)},
				"rules": {
					Type: schema.TypeList, Required: true, MinItems: 1,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name":  {Type: schema.TypeString, Required: true, ValidateFunc: validateS3NonEmptyString},
						"value": {Type: schema.TypeString, Required: true},
					}},
				},
			}},
		},
	}
}

func buildS3SubclientProperties(d *schema.ResourceData) (*handler.MsgS3SubclientProperties, error) {
	applicationID, instanceType := 134, "AMAZON_S3"
	content := make([]handler.MsgS3Content, 0)
	hasPath, hasExclude, hasInclude := false, false, false
	for _, raw := range d.Get("content").([]interface{}) {
		entry := raw.(map[string]interface{})
		item := handler.MsgS3Content{}
		count := 0
		if value := entry["path"].(string); value != "" {
			item.Path, hasPath = &value, true
			count++
		}
		if value := entry["exclude_path"].(string); value != "" {
			item.ExcludePath, hasExclude = &value, true
			count++
		}
		if value := entry["include_path"].(string); value != "" {
			item.IncludePath, hasInclude = &value, true
			count++
		}
		if count != 1 {
			return nil, fmt.Errorf("each content block must set exactly one non-empty path, exclude_path, or include_path")
		}
		content = append(content, item)
	}
	filters := make([]handler.MsgS3TagFilter, 0)
	for _, raw := range d.Get("tag_filter").([]interface{}) {
		entry := raw.(map[string]interface{})
		ruleType := entry["rule_item_type"].(int)
		filter := handler.MsgS3TagFilter{RuleItemType: &ruleType}
		for _, rawRule := range entry["rules"].([]interface{}) {
			rule := rawRule.(map[string]interface{})
			name, value := rule["name"].(string), rule["value"].(string)
			filter.Rules = append(filter.Rules, handler.MsgS3TagRule{Name: &name, Value: &value})
		}
		filters = append(filters, filter)
	}
	contentType := 0
	if len(filters) > 0 {
		if len(content) > 0 {
			return nil, fmt.Errorf("tag_filter and path-based content cannot be combined")
		}
		contentType = 1
	} else if !hasPath {
		return nil, fmt.Errorf("at least one content.path or tag_filter is required")
	}
	contentOp, excludeOp, includeOp := "CLEAR", "CLEAR", "CLEAR"
	if hasPath {
		contentOp = "OVERWRITE"
	}
	if hasExclude {
		excludeOp = "OVERWRITE"
	}
	if hasInclude {
		includeOp = "OVERWRITE"
	}
	objectStorage := &handler.MsgS3ObjectStorageSubclient{BackupContentType: &contentType}
	objectStorage.IsEventScanEnabled = cloudOptionalBool(d, "event_scan_enabled")
	if path := cloudOptionalString(d, "event_scan_staging_path"); path != nil {
		objectStorage.EventScanProperties = &handler.MsgS3EventScanProperties{CloudStagingPath: path}
	}
	if objectStorage.IsEventScanEnabled != nil && *objectStorage.IsEventScanEnabled &&
		(objectStorage.EventScanProperties == nil || *objectStorage.EventScanProperties.CloudStagingPath == "") {
		return nil, fmt.Errorf("event_scan_staging_path is required when event_scan_enabled is true")
	}
	if len(filters) > 0 {
		objectStorage.ObjectStorageTagFilter = &handler.MsgS3TagFilters{TagFilters: filters}
	}
	return &handler.MsgS3SubclientProperties{
		SubClientEntity: &handler.MsgS3Entity{
			SubclientName: stringPtrFromResource(d, "subclient_name"),
			ClientId:      intPtrFromResource(d, "client_id"), InstanceId: intPtrFromResource(d, "instance_id"),
			ApplicationId: &applicationID,
		},
		PlanEntity: &handler.MsgS3Plan{PlanId: intPtrFromResource(d, "plan_id")},
		CloudAppsSubClientProp: &handler.MsgS3CloudAppsSubclient{
			InstanceType: &instanceType, CollectACL: boolPtrFromResource(d, "collect_acl"),
			ObjectStorageSubclient: objectStorage,
		},
		CommonProperties:       &handler.MsgS3CommonProperties{NumberOfBackupStreams: intPtrFromResource(d, "number_of_backup_streams")},
		Content:                &content,
		FsContentOperationType: &contentOp, FsExcludeFilterOperationType: &excludeOp, FsIncludeFilterOperationType: &includeOp,
	}, nil
}

func resourceCreateS3ObjectStorageSubclient(d *schema.ResourceData, m interface{}) error {
	properties, err := buildS3SubclientProperties(d)
	if err != nil {
		return err
	}
	results, err := handler.CvCreateS3Subclient(handler.MsgS3SubclientRequest{SubClientProperties: properties})
	if err != nil {
		return fmt.Errorf("create S3 subclient: %w", err)
	}
	for _, result := range results {
		if result.Entity != nil && result.Entity.SubclientId != nil && *result.Entity.SubclientId > 0 {
			d.SetId(strconv.Itoa(*result.Entity.SubclientId))
			break
		}
	}
	if d.Id() == "" {
		return fmt.Errorf("S3 subclient creation returned no subclientId; verify creation and import before retrying")
	}
	if start := intPtrFromResource(d, "job_start_time"); start != nil {
		_, err := handler.CvModifyS3Subclient(d.Id(), handler.MsgS3SubclientRequest{
			SubClientProperties: &handler.MsgS3SubclientProperties{CommonProperties: &handler.MsgS3CommonProperties{JobStartTime: start}},
		})
		if err != nil {
			return fmt.Errorf("set S3 subclient job start time: %w", err)
		}
	}
	return resourceReadS3ObjectStorageSubclient(d, m)
}

func resourceReadS3ObjectStorageSubclient(d *schema.ResourceData, m interface{}) error {
	return readS3ObjectStorageSubclient(d, false)
}

func readS3ObjectStorageSubclient(d *schema.ResourceData, datasource bool) error {
	if err := validateS3ID(d.Id()); err != nil {
		return err
	}
	properties, err := handler.CvGetS3Subclient(d.Id())
	if err != nil {
		if !datasource && strings.HasPrefix(err.Error(), "status: 404,") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("read S3 subclient %s: %w", d.Id(), err)
	}
	if len(properties) == 0 && !datasource {
		d.SetId("")
		return nil
	}
	if len(properties) != 1 {
		return fmt.Errorf("expected one S3 subclient for ID %s, got %d", d.Id(), len(properties))
	}
	prop := properties[0]
	if prop.SubClientEntity == nil || prop.SubClientEntity.ApplicationId == nil || *prop.SubClientEntity.ApplicationId != 134 ||
		prop.CloudAppsSubClientProp == nil || prop.CloudAppsSubClientProp.InstanceType == nil || *prop.CloudAppsSubClientProp.InstanceType != "AMAZON_S3" {
		return fmt.Errorf("subclient %s is not an AMAZON_S3 Cloud Apps subclient", d.Id())
	}
	values := map[string]interface{}{}
	id, _ := strconv.Atoi(d.Id())
	values["subclient_id"] = id
	s3PutString(values, "subclient_name", prop.SubClientEntity.SubclientName)
	s3PutInt(values, "client_id", prop.SubClientEntity.ClientId)
	s3PutInt(values, "instance_id", prop.SubClientEntity.InstanceId)
	if prop.PlanEntity != nil {
		s3PutInt(values, "plan_id", prop.PlanEntity.Id)
		s3PutInt(values, "plan_id", prop.PlanEntity.PlanId)
		s3PutString(values, "plan_name", prop.PlanEntity.Name)
		s3PutString(values, "plan_name", prop.PlanEntity.PlanName)
	}
	if prop.CommonProperties != nil {
		s3PutInt(values, "number_of_backup_streams", prop.CommonProperties.NumberOfBackupStreams)
		s3PutInt(values, "job_start_time", prop.CommonProperties.JobStartTime)
	}
	if prop.CloudAppsSubClientProp.CollectACL != nil {
		values["collect_acl"] = *prop.CloudAppsSubClientProp.CollectACL
	}
	content := make([]map[string]interface{}, 0)
	if prop.Content != nil {
		for _, item := range *prop.Content {
			entry := map[string]interface{}{}
			s3PutString(entry, "path", item.Path)
			s3PutString(entry, "exclude_path", item.ExcludePath)
			s3PutString(entry, "include_path", item.IncludePath)
			content = append(content, entry)
		}
	}
	values["content"] = content
	filters := make([]map[string]interface{}, 0)
	if object := prop.CloudAppsSubClientProp.ObjectStorageSubclient; object != nil {
		if object.IsEventScanEnabled != nil {
			values["event_scan_enabled"] = *object.IsEventScanEnabled
		}
		if object.EventScanProperties != nil {
			s3PutString(values, "event_scan_staging_path", object.EventScanProperties.CloudStagingPath)
		}
		s3PutInt(values, "backup_content_type", object.BackupContentType)
		if object.ObjectStorageTagFilter != nil {
			for _, filter := range object.ObjectStorageTagFilter.TagFilters {
				entry := map[string]interface{}{}
				s3PutInt(entry, "rule_item_type", filter.RuleItemType)
				rules := make([]map[string]interface{}, 0)
				for _, rule := range filter.Rules {
					value := map[string]interface{}{}
					s3PutString(value, "name", rule.Name)
					s3PutString(value, "value", rule.Value)
					rules = append(rules, value)
				}
				entry["rules"] = rules
				filters = append(filters, entry)
			}
		}
	}
	values["tag_filter"] = filters
	return setS3State(d, values)
}

func resourceUpdateS3ObjectStorageSubclient(d *schema.ResourceData, m interface{}) error {
	if err := validateS3ID(d.Id()); err != nil {
		return err
	}
	properties, err := buildS3SubclientProperties(d)
	if err != nil {
		return err
	}
	if d.HasChanges("subclient_name", "plan_id", "number_of_backup_streams", "job_start_time", "collect_acl", "event_scan_enabled", "event_scan_staging_path") {
		props := &handler.MsgS3SubclientProperties{
			PlanEntity: properties.PlanEntity, CommonProperties: properties.CommonProperties,
			CloudAppsSubClientProp: &handler.MsgS3CloudAppsSubclient{
				InstanceType: properties.CloudAppsSubClientProp.InstanceType,
				CollectACL:   properties.CloudAppsSubClientProp.CollectACL,
			},
		}
		props.CommonProperties.JobStartTime = intPtrFromResource(d, "job_start_time")
		if d.HasChanges("event_scan_enabled", "event_scan_staging_path") {
			props.CloudAppsSubClientProp.ObjectStorageSubclient = &handler.MsgS3ObjectStorageSubclient{
				IsEventScanEnabled:  properties.CloudAppsSubClientProp.ObjectStorageSubclient.IsEventScanEnabled,
				EventScanProperties: properties.CloudAppsSubClientProp.ObjectStorageSubclient.EventScanProperties,
			}
		}
		req := handler.MsgS3SubclientRequest{SubClientProperties: props}
		if d.HasChange("subclient_name") {
			req.NewName = stringPtrFromResource(d, "subclient_name")
		}
		if _, err := handler.CvModifyS3Subclient(d.Id(), req); err != nil {
			return fmt.Errorf("modify S3 subclient properties: %w", err)
		}
	}
	if d.HasChanges("content", "tag_filter") {
		props := &handler.MsgS3SubclientProperties{
			CloudAppsSubClientProp: properties.CloudAppsSubClientProp, Content: properties.Content,
			FsContentOperationType:       properties.FsContentOperationType,
			FsExcludeFilterOperationType: properties.FsExcludeFilterOperationType,
			FsIncludeFilterOperationType: properties.FsIncludeFilterOperationType,
		}
		if _, err := handler.CvModifyS3SubclientContent(d.Id(), handler.MsgS3SubclientRequest{SubClientProperties: props}); err != nil {
			return fmt.Errorf("modify S3 subclient content: %w", err)
		}
	}
	return resourceReadS3ObjectStorageSubclient(d, m)
}

func resourceDeleteS3ObjectStorageSubclient(d *schema.ResourceData, m interface{}) error {
	if err := validateS3ID(d.Id()); err != nil {
		return err
	}
	_, err := handler.CvDeleteS3Subclient(d.Id())
	if err != nil && !strings.HasPrefix(err.Error(), "status: 404,") {
		return fmt.Errorf("delete S3 subclient: %w", err)
	}
	d.SetId("")
	return nil
}
