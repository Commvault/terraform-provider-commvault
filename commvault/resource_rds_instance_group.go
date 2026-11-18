package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRDSInstanceGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceCreateRDSInstanceGroup,
		Read:     resourceReadRDSInstanceGroup,
		Update:   resourceUpdateRDSInstanceGroup,
		Delete:   resourceDeleteRDSInstanceGroup,
		Importer: &schema.ResourceImporter{State: schema.ImportStatePassthrough},
		Schema: map[string]*schema.Schema{
			"client_name": {Type: schema.TypeString, Required: true, ForceNew: true,
				ValidateFunc: validateNonEmptyString},
			"instance_name": {Type: schema.TypeString, Required: true,
				ValidateFunc: validateNonEmptyString},
			"display_name": {Type: schema.TypeString, Required: true, ForceNew: true,
				ValidateFunc: validateNonEmptyString, Description: "Cloud account display name."},
			"subclient_name": {Type: schema.TypeString, Required: true,
				ValidateFunc: validateNonEmptyString},
			"plan_name": {Type: schema.TypeString, Required: true,
				ValidateFunc: validateNonEmptyString},
			"number_of_backup_streams":           {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateIntRange(1, 2147483647)},
			"job_start_time":                     {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateIntRange(0, 86399)},
			"auto_update_list":                   {Type: schema.TypeBool, Optional: true, Computed: true},
			"enable_external_snapshot_discovery": {Type: schema.TypeBool, Optional: true, Computed: true},
			"enable_snapshot_backup_copy":        {Type: schema.TypeBool, Optional: true, Computed: true},
			"wait_for_snapshot":                  {Type: schema.TypeBool, Optional: true, Computed: true},
			"content": {
				Type: schema.TypeList, Required: true, MinItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"display_name": {Type: schema.TypeString, Required: true, ValidateFunc: validateNonEmptyString},
					"name":         {Type: schema.TypeString, Required: true, ValidateFunc: validateNonEmptyString},
					"path":         {Type: schema.TypeString, Required: true},
					"type": {Type: schema.TypeString, Required: true,
						ValidateFunc: validateRDSStringEnum("Cluster", "Instance", "REGION")},
					"update_type": {Type: schema.TypeInt, Required: true, ValidateFunc: validateIntRange(1, 2147483647),
						Description: "Numeric content type: REGION=4, Cluster=48, Instance=50."},
					"value": {Type: schema.TypeString, Required: true, ValidateFunc: validateNonEmptyString,
						Description: "Update selection value, such as region or region/database."},
					"all_or_any_children": {Type: schema.TypeBool, Optional: true, Default: true},
					"negation":            {Type: schema.TypeBool, Optional: true, Default: false},
				}},
			},
			"subclient_id": {Type: schema.TypeInt, Computed: true},
			"instance_id":  {Type: schema.TypeInt, Computed: true},
			"client_id":    {Type: schema.TypeInt, Computed: true},
		},
	}
}

func buildRDSGroupContent(d *schema.ResourceData) ([]handler.MsgRDSCreateContent, []handler.MsgRDSModifyContent) {
	create := make([]handler.MsgRDSCreateContent, 0)
	modify := make([]handler.MsgRDSModifyContent, 0)
	for _, raw := range d.Get("content").([]interface{}) {
		item := raw.(map[string]interface{})
		display, name, path := item["display_name"].(string), item["name"].(string), item["path"].(string)
		contentType, updateType, value := item["type"].(string), item["update_type"].(int), item["value"].(string)
		all, negation := item["all_or_any_children"].(bool), item["negation"].(bool)
		create = append(create, handler.MsgRDSCreateContent{
			DisplayName: &display, Name: &name, Path: &path, Type: &contentType, AllOrAnyChildren: &all, Negation: &negation,
		})
		modify = append(modify, handler.MsgRDSModifyContent{
			DisplayName: &display, Name: &name, Path: &path, Type: &updateType, Value: &value, AllOrAnyChildren: &all, Negation: &negation,
		})
	}
	return create, modify
}

func resourceCreateRDSInstanceGroup(d *schema.ResourceData, m interface{}) error {
	appName, instanceType := "Cloud Apps", "AMAZON_RDS"
	subclientID, contentOperation := -1, 2
	content, _ := buildRDSGroupContent(d)
	response, err := handler.CvCreateRDSInstanceGroup(handler.MsgCreateRDSInstanceGroupRequest{
		SubclientProperties: &handler.MsgRDSCreateSubclientProperties{
			SubclientEntity: &handler.MsgRDSSubclientEntity{
				ClientName: stringPtrFromResource(d, "client_name"), InstanceName: stringPtrFromResource(d, "instance_name"),
				DisplayName: stringPtrFromResource(d, "display_name"), SubclientName: stringPtrFromResource(d, "subclient_name"),
				AppName: &appName, SubclientId: &subclientID,
			},
			PlanEntity:             &handler.MsgRDSPlan{PlanName: stringPtrFromResource(d, "plan_name")},
			ContentOperationType:   &contentOperation,
			CloudAppsSubclientProp: &handler.MsgRDSCloudAppsSubclientProperties{InstanceType: &instanceType, RDSSubclient: buildRDSSubclient(d)},
			CommonProperties: &handler.MsgS3CommonProperties{
				NumberOfBackupStreams: intPtrFromResource(d, "number_of_backup_streams"), JobStartTime: intPtrFromResource(d, "job_start_time"),
			},
			CloudDbContent: &handler.MsgRDSCreateCloudDbContent{Children: content},
		},
	})
	if err != nil {
		return fmt.Errorf("create RDS instance group: %w", err)
	}
	for _, result := range response.Response {
		if result.Entity != nil && result.Entity.SubclientId != nil && *result.Entity.SubclientId > 0 {
			d.SetId(strconv.Itoa(*result.Entity.SubclientId))
			if err := d.Set("subclient_id", *result.Entity.SubclientId); err != nil {
				return fmt.Errorf("set RDS subclient_id: %w", err)
			}
			break
		}
	}
	// The create endpoint can ignore snapshot settings (notably explicit false).
	for _, key := range []string{"number_of_backup_streams", "job_start_time", "auto_update_list",
		"enable_external_snapshot_discovery", "enable_snapshot_backup_copy", "wait_for_snapshot"} {
		if _, configured := d.GetOkExists(key); !configured {
			continue
		}
		_, err := handler.CvModifyRDSInstanceGroup(d.Id(), handler.MsgModifyRDSInstanceGroupRequest{
			SubclientProperties: &handler.MsgRDSModifySubclientProperties{
				CloudAppsSubclientProp: &handler.MsgRDSCloudAppsSubclientProperties{InstanceType: &instanceType, RDSSubclient: buildRDSSubclient(d)},
				CommonProperties: &handler.MsgS3CommonProperties{
					NumberOfBackupStreams: intPtrFromResource(d, "number_of_backup_streams"), JobStartTime: intPtrFromResource(d, "job_start_time"),
				},
			},
			Association: &handler.MsgRDSAssociation{Entity: []handler.MsgRDSAssociationEntity{{
				ClientName: stringPtrFromResource(d, "client_name"), InstanceName: stringPtrFromResource(d, "instance_name"),
				SubclientName: stringPtrFromResource(d, "subclient_name"), AppName: &appName,
			}}},
		})
		if err != nil {
			return fmt.Errorf("set initial RDS group settings (created ID retained): %w", err)
		}
		break
	}
	return resourceReadRDSInstanceGroup(d, m)
}

func resourceUpdateRDSInstanceGroup(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("instance_name") {
		properties, err := handler.CvGetRDSInstanceGroup(d.Id())
		if err != nil {
			return err
		}
		if len(properties) != 1 || properties[0].SubclientEntity == nil {
			return fmt.Errorf("cannot verify parent identity for RDS group %s", d.Id())
		}
		name, err := rdsGroupInstanceName(properties[0].SubclientEntity)
		if err != nil {
			return err
		}
		if name == nil || *name != d.Get("instance_name").(string) {
			return fmt.Errorf("instance_name can follow a parent rename, but moving an RDS group to another instance is unsupported")
		}
	}
	if !d.HasChanges("content", "plan_name", "subclient_name", "instance_name", "number_of_backup_streams",
		"job_start_time", "auto_update_list", "enable_external_snapshot_discovery", "enable_snapshot_backup_copy", "wait_for_snapshot") {
		return resourceReadRDSInstanceGroup(d, m)
	}
	_, content := buildRDSGroupContent(d)
	instanceType, appName, entityType, backupset := "AMAZON_RDS", "Cloud Apps", "SUBCLIENT_ENTITY", "defaultBackupSet"
	properties := &handler.MsgRDSModifySubclientProperties{}
	if d.HasChanges("number_of_backup_streams", "job_start_time") {
		properties.CommonProperties = &handler.MsgS3CommonProperties{
			NumberOfBackupStreams: intPtrFromResource(d, "number_of_backup_streams"), JobStartTime: intPtrFromResource(d, "job_start_time"),
		}
	}
	if d.HasChanges("auto_update_list", "enable_external_snapshot_discovery", "enable_snapshot_backup_copy", "wait_for_snapshot") {
		properties.CloudAppsSubclientProp = &handler.MsgRDSCloudAppsSubclientProperties{InstanceType: &instanceType, RDSSubclient: buildRDSSubclient(d)}
	}
	if d.HasChange("content") {
		if properties.CloudAppsSubclientProp == nil {
			properties.CloudAppsSubclientProp = &handler.MsgRDSCloudAppsSubclientProperties{InstanceType: &instanceType}
		}
		properties.CloudDbContent = &handler.MsgRDSModifyCloudDbContent{Children: content}
	}
	if d.HasChange("plan_name") {
		properties.PlanEntity = &handler.MsgRDSPlan{PlanName: stringPtrFromResource(d, "plan_name")}
	}
	request := handler.MsgModifyRDSInstanceGroupRequest{
		SubclientProperties: properties,
		Association: &handler.MsgRDSAssociation{Entity: []handler.MsgRDSAssociationEntity{{
			ClientName: stringPtrFromResource(d, "client_name"), InstanceName: stringPtrFromResource(d, "instance_name"),
			DisplayName: stringPtrFromResource(d, "display_name"), SubclientName: stringPtrFromResource(d, "subclient_name"),
			AppName: &appName, Type: &entityType, BackupsetName: &backupset,
		}}},
	}
	if id := d.Get("instance_id").(int); id > 0 {
		subclientID, err := strconv.Atoi(d.Id())
		if err != nil {
			return err
		}
		applicationID := 134
		request.Association.Entity = []handler.MsgRDSAssociationEntity{{
			InstanceId: &id, ClientId: intPtrFromResource(d, "client_id"), SubclientId: &subclientID,
			ApplicationId: &applicationID,
		}}
	}
	if d.HasChange("subclient_name") {
		old, _ := d.GetChange("subclient_name")
		oldName := old.(string)
		if request.Association.Entity[0].SubclientId == nil {
			request.Association.Entity[0].SubclientName = &oldName
		}
		request.NewName = stringPtrFromResource(d, "subclient_name")
	}
	_, err := handler.CvModifyRDSInstanceGroup(d.Id(), request)
	if err != nil {
		return fmt.Errorf("modify RDS instance group: %w", err)
	}
	return resourceReadRDSInstanceGroup(d, m)
}

func resourceReadRDSInstanceGroup(d *schema.ResourceData, m interface{}) error {
	return readRDSInstanceGroup(d, false)
}

func readRDSInstanceGroup(d *schema.ResourceData, datasource bool) error {
	properties, err := handler.CvGetRDSInstanceGroup(d.Id())
	if rdsResourceAbsent(d, datasource, err, len(properties)) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("read RDS instance group %s: %w", d.Id(), err)
	}
	if len(properties) != 1 {
		return fmt.Errorf("expected one RDS group for %s, got %d", d.Id(), len(properties))
	}
	prop := properties[0]
	entity := prop.SubclientEntity
	if entity == nil || entity.SubclientId == nil || strconv.Itoa(*entity.SubclientId) != d.Id() ||
		entity.ApplicationId == nil || *entity.ApplicationId != 134 ||
		prop.CloudAppsSubclientProp == nil || prop.CloudAppsSubclientProp.InstanceType == nil ||
		*prop.CloudAppsSubclientProp.InstanceType != "AMAZON_RDS" {
		return fmt.Errorf("subclient %s does not identify the requested Cloud Apps AMAZON_RDS group", d.Id())
	}
	values := map[string]interface{}{"subclient_id": *entity.SubclientId}
	s3PutInt(values, "instance_id", entity.InstanceId)
	s3PutInt(values, "client_id", entity.ClientId)
	if common := prop.CommonProperties; common != nil {
		s3PutInt(values, "number_of_backup_streams", common.NumberOfBackupStreams)
		s3PutInt(values, "job_start_time", common.JobStartTime)
	}
	if settings := prop.CloudAppsSubclientProp.RDSSubclient; settings != nil {
		for key, value := range map[string]*bool{
			"auto_update_list": settings.AutoUpdateList, "enable_external_snapshot_discovery": settings.EnableExternalSnapshotDiscovery,
			"enable_snapshot_backup_copy": settings.EnableSnapshotBackupCopy, "wait_for_snapshot": settings.WaitForSnapshot,
		} {
			if value != nil {
				values[key] = *value
			}
		}
	}
	rdsPutString(values, "client_name", entity.ClientName)
	instanceName, err := rdsGroupInstanceName(entity)
	if err != nil {
		return err
	}
	rdsPutString(values, "instance_name", instanceName)
	rdsPutString(values, "display_name", entity.DisplayName)
	rdsPutString(values, "subclient_name", entity.SubclientName)
	if prop.PlanEntity != nil {
		rdsPutString(values, "plan_name", prop.PlanEntity.Name)
		rdsPutString(values, "plan_name", prop.PlanEntity.PlanName)
	}
	content := make([]map[string]interface{}, 0)
	if prop.CloudDbContent == nil {
		return fmt.Errorf("RDS group %s GET omits cloudDbContent; cannot refresh managed selections", d.Id())
	}
	for _, selection := range prop.CloudDbContent.Children {
		item, err := flattenRDSContent(selection, d.Get("content").([]interface{}))
		if err != nil {
			return fmt.Errorf("read RDS group %s content: %w", d.Id(), err)
		}
		content = append(content, item)
	}
	values["content"] = content
	return setRDSState(d, values)
}

func rdsGroupInstanceName(entity *handler.MsgRDSSubclientEntity) (*string, error) {
	if entity.InstanceId == nil {
		return entity.InstanceName, nil
	}
	// Subclient GET can retain the old embedded name after its parent is renamed.
	properties, err := handler.CvGetRDSInstance(strconv.Itoa(*entity.InstanceId))
	if err != nil {
		return nil, fmt.Errorf("read RDS group parent: %w", err)
	}
	if len(properties) != 1 || properties[0].Instance == nil || properties[0].Instance.InstanceId == nil ||
		*properties[0].Instance.InstanceId != *entity.InstanceId || properties[0].CloudAppsInstance == nil ||
		properties[0].CloudAppsInstance.InstanceType == nil || *properties[0].CloudAppsInstance.InstanceType != "AMAZON_RDS" {
		return nil, fmt.Errorf("RDS group parent %d has an invalid identity", *entity.InstanceId)
	}
	return properties[0].Instance.InstanceName, nil
}

func flattenRDSContent(selection handler.MsgRDSReadContent, configured []interface{}) (map[string]interface{}, error) {
	if len(selection.Children) != 0 {
		return nil, fmt.Errorf("nested RDS selections are not supported; refusing to flatten away selection semantics")
	}
	var named string
	if err := json.Unmarshal(selection.Type, &named); err != nil {
		var number int
		if err := json.Unmarshal(selection.Type, &number); err != nil {
			return nil, fmt.Errorf("invalid RDS content type %s", selection.Type)
		}
		switch number {
		case 4:
			named = "REGION"
		case 50:
			named = "Instance"
		case 48:
			named = "Cluster"
		default:
			for _, raw := range configured {
				item := raw.(map[string]interface{})
				if item["update_type"].(int) == number {
					named = item["type"].(string)
					break
				}
			}
			if named == "" {
				return nil, fmt.Errorf("unknown RDS content type %d; a verified named/numeric mapping is required", number)
			}
		}
	}
	updateType := 0
	switch named {
	case "REGION":
		updateType = 4
	case "Instance":
		updateType = 50
	case "Cluster":
		updateType = 48
	default:
		return nil, fmt.Errorf("unsupported RDS content type %q", named)
	}
	if updateType <= 0 || selection.Name == nil || selection.DisplayName == nil ||
		selection.AllOrAnyChildren == nil || selection.Negation == nil {
		return nil, fmt.Errorf("incomplete RDS %s selection or unverified update type", named)
	}
	path := ""
	if selection.Path != nil {
		path = *selection.Path
	}
	value := *selection.Name
	if named != "REGION" && path != "" {
		value = strings.TrimSuffix(path, "/") + "/" + value
	}
	if selection.Value != nil {
		value = *selection.Value
	}
	return map[string]interface{}{
		"display_name": *selection.DisplayName, "name": *selection.Name, "path": path,
		"type": named, "update_type": updateType, "value": value,
		"all_or_any_children": *selection.AllOrAnyChildren, "negation": *selection.Negation,
	}, nil
}

func buildRDSSubclient(d *schema.ResourceData) *handler.MsgRDSSubclient {
	return &handler.MsgRDSSubclient{
		AutoUpdateList:                  cloudOptionalBool(d, "auto_update_list"),
		EnableExternalSnapshotDiscovery: cloudOptionalBool(d, "enable_external_snapshot_discovery"),
		EnableSnapshotBackupCopy:        cloudOptionalBool(d, "enable_snapshot_backup_copy"),
		WaitForSnapshot:                 cloudOptionalBool(d, "wait_for_snapshot"),
	}
}

func resourceDeleteRDSInstanceGroup(d *schema.ResourceData, m interface{}) error {
	if err := resourceReadRDSInstanceGroup(d, m); err != nil || d.Id() == "" {
		return err
	}
	if d.Get("subclient_name").(string) == "default" {
		return fmt.Errorf("refusing to delete the RDS default subclient; manage only explicitly created instance groups")
	}
	if _, err := handler.CvDeleteRDSInstanceGroup(d.Id()); err != nil && !strings.HasPrefix(err.Error(), "status: 404,") {
		return fmt.Errorf("delete RDS group %s: %w", d.Id(), err)
	}
	d.SetId("")
	return nil
}
