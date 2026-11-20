package commvault

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRDSInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceCreateRDSInstance,
		Read:     resourceReadRDSInstance,
		Update:   resourceUpdateRDSInstance,
		Delete:   resourceDeleteRDSInstance,
		Importer: &schema.ResourceImporter{State: schema.ImportStatePassthrough},
		Schema: map[string]*schema.Schema{
			"client_name": {Type: schema.TypeString, Required: true, ForceNew: true,
				ValidateFunc: validateNonEmptyString, Description: "Existing Commvault cloud account name."},
			"instance_name": {Type: schema.TypeString, Required: true,
				ValidateFunc: validateNonEmptyString},
			"plan_name": {Type: schema.TypeString, Required: true,
				ValidateFunc: validateNonEmptyString},
			"access_node": {
				Type: schema.TypeList, Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"client_id": {Type: schema.TypeInt, Required: true, ValidateFunc: validateIntRange(1, 2147483647)},
					"client_name": {Type: schema.TypeString, Required: true,
						ValidateFunc: validateNonEmptyString},
					"type": {Type: schema.TypeString, Optional: true, Default: "CLIENT_ENTITY",
						ValidateFunc: validateRDSStringEnum("CLIENT_ENTITY", "CLIENTGROUP_ENTITY")},
				}},
			},
			"use_resource_pool_info":                  {Type: schema.TypeBool, Optional: true, Default: false, ForceNew: true},
			"number_of_backup_streams":                {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateIntRange(1, 2147483647)},
			"region_end_points":                       {Type: schema.TypeString, Optional: true, Computed: true},
			"point_in_time_restore_enabled":           {Type: schema.TypeBool, Optional: true, Computed: true},
			"point_in_time_restore_retention_period":  {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateIntRange(0, 35)},
			"point_in_time_restore_backup_period":     {Type: schema.TypeString, Optional: true, Computed: true},
			"point_in_time_restore_apply_immediately": {Type: schema.TypeBool, Optional: true, Computed: true},
			"credential_name": {Type: schema.TypeString, Optional: true,
				Description: "Optional existing credential name applied after creation. Omit to inherit the cloud account credential; clearing a managed credential is unsupported."},
			"use_iam_role": {Type: schema.TypeBool, Optional: true, Default: false,
				Description: "Used only with credential_name, via the documented credential update."},
			"instance_id":          {Type: schema.TypeInt, Computed: true},
			"client_id":            {Type: schema.TypeInt, Computed: true},
			"default_subclient_id": {Type: schema.TypeInt, Computed: true},
		},
		CustomizeDiff: func(d *schema.ResourceDiff, m interface{}) error {
			if d.Get("credential_name").(string) == "" &&
				(d.Get("use_iam_role").(bool) || (d.Id() != "" && d.HasChange("credential_name"))) {
				return fmt.Errorf("credential_name is required for IAM-role configuration or credential changes; clearing credentials is unsupported")
			}
			return nil
		},
	}
}

func resourceCreateRDSInstance(d *schema.ResourceData, m interface{}) error {
	applicationID, instanceType := 134, "AMAZON_RDS"
	properties := &handler.MsgRDSInstanceProperties{
		Instance: &handler.MsgRDSInstance{
			ClientName: stringPtrFromResource(d, "client_name"), InstanceName: stringPtrFromResource(d, "instance_name"),
			ApplicationId: &applicationID,
		},
		PlanEntity:          &handler.MsgRDSPlan{PlanName: stringPtrFromResource(d, "plan_name")},
		UseResourcePoolInfo: boolPtrFromResource(d, "use_resource_pool_info"),
		CloudAppsInstance: &handler.MsgRDSCloudAppsInstance{
			InstanceType: &instanceType, GeneralCloudProperties: buildRDSGeneralProperties(d),
		},
	}
	response, err := handler.CvCreateRDSInstance(handler.MsgCreateRDSInstanceRequest{InstanceProperties: properties})
	if err != nil {
		return fmt.Errorf("create RDS instance: %w", err)
	}
	for _, result := range response.Response {
		entity := result.Entity
		if entity == nil || entity.InstanceId == nil || *entity.InstanceId <= 0 {
			continue
		}
		d.SetId(strconv.Itoa(*entity.InstanceId))
		values := map[string]*int{
			"instance_id": entity.InstanceId, "client_id": entity.ClientId, "default_subclient_id": entity.SubclientId,
		}
		for key, value := range values {
			if value != nil {
				if err := d.Set(key, *value); err != nil {
					return fmt.Errorf("set RDS %s: %w", key, err)
				}
			}
		}
		break
	}
	// Preserve the created ID if the subsequent credential update fails.
	if d.Get("credential_name").(string) != "" {
		if err := modifyRDSInstance(d, true); err != nil {
			return err
		}
	}
	return resourceReadRDSInstance(d, m)
}

func buildRDSGeneralProperties(d *schema.ResourceData) *handler.MsgRDSGeneralCloudProperties {
	general := &handler.MsgRDSGeneralCloudProperties{
		NumberOfBackupStreams:              intPtrFromResource(d, "number_of_backup_streams"),
		RegionEndPoints:                    cloudOptionalString(d, "region_end_points"),
		PointInTimeRestoreEnabled:          cloudOptionalBool(d, "point_in_time_restore_enabled"),
		PointInTimeRestoreRetentionPeriod:  intPtrFromResource(d, "point_in_time_restore_retention_period"),
		PointInTimeRestoreBackupPeriod:     cloudOptionalString(d, "point_in_time_restore_backup_period"),
		PointInTimeRestoreApplyImmediately: cloudOptionalBool(d, "point_in_time_restore_apply_immediately"),
	}
	if nodes := d.Get("access_node").([]interface{}); len(nodes) > 0 {
		servers := make([]handler.MsgRDSMemberServer, 0, len(nodes))
		for _, node := range nodes {
			values := node.(map[string]interface{})
			id, name, entityType := values["client_id"].(int), values["client_name"].(string), values["type"].(string)
			servers = append(servers, handler.MsgRDSMemberServer{
				Client: &handler.MsgRDSAccessNodeClient{ClientId: &id, ClientName: &name, Type: &entityType},
			})
		}
		general.AccessNodes = &handler.MsgRDSAccessNodes{MemberServers: servers}
	}
	return general
}

func modifyRDSInstance(d *schema.ResourceData, credentialsOnly bool) error {
	instanceType, appName := "AMAZON_RDS", "Cloud Apps"
	clientName, instanceName := d.Get("client_name").(string), d.Get("instance_name").(string)
	properties := &handler.MsgRDSInstanceProperties{}
	request := handler.MsgModifyRDSInstanceRequest{InstanceProperties: properties}
	if !credentialsOnly && d.HasChange("instance_name") {
		old, _ := d.GetChange("instance_name")
		instanceName = old.(string)
		request.NewName = stringPtrFromResource(d, "instance_name")
		id, err := strconv.Atoi(d.Id())
		if err != nil || id <= 0 {
			return fmt.Errorf("invalid RDS instance ID %q", d.Id())
		}
		properties.Instance = &handler.MsgRDSInstance{InstanceId: &id, ClientId: intPtrFromResource(d, "client_id")}
		properties.CloudAppsInstance = &handler.MsgRDSCloudAppsInstance{InstanceType: &instanceType, GeneralCloudProperties: buildRDSGeneralProperties(d)}
	}
	if !credentialsOnly && d.HasChanges("access_node", "number_of_backup_streams", "region_end_points",
		"point_in_time_restore_enabled", "point_in_time_restore_retention_period",
		"point_in_time_restore_backup_period", "point_in_time_restore_apply_immediately") {
		general := buildRDSGeneralProperties(d)
		if d.HasChange("access_node") && general.AccessNodes == nil {
			return fmt.Errorf("clearing all RDS access nodes is unsupported; supply at least one access_node")
		}
		properties.CloudAppsInstance = &handler.MsgRDSCloudAppsInstance{InstanceType: &instanceType, GeneralCloudProperties: general}
	}
	if credentialsOnly || d.HasChanges("credential_name", "use_iam_role") {
		name := d.Get("credential_name").(string)
		if name == "" {
			return fmt.Errorf("credential_name is required for the RDS credential update")
		}
		if properties.CloudAppsInstance == nil {
			properties.CloudAppsInstance = &handler.MsgRDSCloudAppsInstance{
				InstanceType: &instanceType, GeneralCloudProperties: &handler.MsgRDSGeneralCloudProperties{},
			}
		}
		if properties.CloudAppsInstance.GeneralCloudProperties == nil {
			properties.CloudAppsInstance.GeneralCloudProperties = &handler.MsgRDSGeneralCloudProperties{}
		}
		general := properties.CloudAppsInstance.GeneralCloudProperties
		general.Credentials = &handler.MsgCredentialInfo{CredentialName: &name}
		general.AmazonInstanceInfo = &handler.MsgRDSAmazonInstanceInfo{UseIamRole: boolPtrFromResource(d, "use_iam_role")}
	}
	if !credentialsOnly && d.HasChange("plan_name") {
		properties.PlanEntity = &handler.MsgRDSPlan{PlanName: stringPtrFromResource(d, "plan_name")}
	}
	request.Association = &handler.MsgRDSAssociation{Entity: []handler.MsgRDSAssociationEntity{{
		ClientName: &clientName, InstanceName: &instanceName, DisplayName: &clientName, AppName: &appName,
	}}}
	_, err := handler.CvModifyRDSInstance(d.Id(), request)
	if err != nil {
		return fmt.Errorf("modify RDS instance %s: %w", d.Id(), err)
	}
	return nil
}

func resourceUpdateRDSInstance(d *schema.ResourceData, m interface{}) error {
	expectedName := d.Get("instance_name").(string)
	renamed := d.HasChange("instance_name")
	if d.HasChanges("credential_name", "use_iam_role", "plan_name", "instance_name", "access_node",
		"number_of_backup_streams", "region_end_points", "point_in_time_restore_enabled",
		"point_in_time_restore_retention_period", "point_in_time_restore_backup_period", "point_in_time_restore_apply_immediately") {
		if err := modifyRDSInstance(d, false); err != nil {
			return err
		}
	}
	if renamed {
		return refreshCloudInstanceRename(d, m, expectedName, resourceReadRDSInstance)
	}
	return resourceReadRDSInstance(d, m)
}

func refreshCloudInstanceRename(d *schema.ResourceData, m interface{}, expected string, read func(*schema.ResourceData, interface{}) error) error {
	for attempt := 0; attempt < 30; attempt++ {
		if err := read(d, m); err != nil {
			return err
		}
		if d.Id() == "" {
			return fmt.Errorf("renamed cloud instance disappeared during refresh")
		}
		if d.Get("instance_name").(string) == expected {
			return nil
		}
		// Reads can lag a successful rename; never repeat the write.
		time.Sleep(time.Second)
	}
	return fmt.Errorf("cloud instance %s rename was accepted but %q was not observed within 30 seconds; inspect before retrying", d.Id(), expected)
}

func resourceReadRDSInstance(d *schema.ResourceData, m interface{}) error {
	return readRDSInstance(d, false)
}

func readRDSInstance(d *schema.ResourceData, datasource bool) error {
	properties, err := handler.CvGetRDSInstance(d.Id())
	if rdsResourceAbsent(d, datasource, err, len(properties)) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("read RDS instance %s: %w", d.Id(), err)
	}
	if len(properties) != 1 {
		return fmt.Errorf("expected one RDS instance for %s, got %d", d.Id(), len(properties))
	}
	prop := properties[0]
	if prop.Instance == nil || prop.Instance.InstanceId == nil || strconv.Itoa(*prop.Instance.InstanceId) != d.Id() ||
		prop.Instance.ApplicationId == nil || *prop.Instance.ApplicationId != 134 ||
		prop.CloudAppsInstance == nil || prop.CloudAppsInstance.InstanceType == nil || *prop.CloudAppsInstance.InstanceType != "AMAZON_RDS" {
		return fmt.Errorf("instance %s does not identify the requested Cloud Apps AMAZON_RDS instance", d.Id())
	}
	values := map[string]interface{}{"instance_id": *prop.Instance.InstanceId}
	rdsPutString(values, "client_name", prop.Instance.ClientName)
	rdsPutString(values, "instance_name", prop.Instance.InstanceName)
	if prop.Instance.ClientId != nil {
		values["client_id"] = *prop.Instance.ClientId
	}
	if prop.PlanEntity != nil {
		rdsPutString(values, "plan_name", prop.PlanEntity.Name)
		rdsPutString(values, "plan_name", prop.PlanEntity.PlanName)
	}
	if !datasource {
		if prop.UseResourcePoolInfo != nil {
			values["use_resource_pool_info"] = *prop.UseResourcePoolInfo
		} else {
			log.Printf("[WARN] RDS instance %s: GET omits create-only use_resource_pool_info; retaining its configured value", d.Id())
		}
	}
	if general := prop.CloudAppsInstance.GeneralCloudProperties; general != nil {
		s3PutInt(values, "number_of_backup_streams", general.NumberOfBackupStreams)
		rdsPutString(values, "region_end_points", general.RegionEndPoints)
		s3PutInt(values, "point_in_time_restore_retention_period", general.PointInTimeRestoreRetentionPeriod)
		rdsPutString(values, "point_in_time_restore_backup_period", general.PointInTimeRestoreBackupPeriod)
		if general.PointInTimeRestoreEnabled != nil {
			values["point_in_time_restore_enabled"] = *general.PointInTimeRestoreEnabled
		}
		if general.PointInTimeRestoreApplyImmediately != nil {
			values["point_in_time_restore_apply_immediately"] = *general.PointInTimeRestoreApplyImmediately
		}
		if general.AccessNodes != nil {
			nodes := make([]map[string]interface{}, 0, len(general.AccessNodes.MemberServers))
			for _, node := range general.AccessNodes.MemberServers {
				if node.Client == nil || node.Client.ClientId == nil || *node.Client.ClientId <= 0 || node.Client.Type == nil {
					return fmt.Errorf("RDS instance %s returned an incomplete access node", d.Id())
				}
				name := node.Client.ClientName
				if name == nil || *name == "" {
					if *node.Client.Type != "CLIENT_ENTITY" {
						return fmt.Errorf("RDS instance %s returned an unnamed non-client access node", d.Id())
					}
					var err error
					name, err = handler.CvGetCloudAccessNodeName(*node.Client.ClientId)
					if err != nil {
						return fmt.Errorf("resolve RDS access-node name: %w", err)
					}
				}
				nodes = append(nodes, map[string]interface{}{
					"client_id": *node.Client.ClientId, "client_name": *name, "type": *node.Client.Type,
				})
			}
			values["access_node"] = nodes
		}
		// Inherited account credentials are intentionally not managed by this resource.
		if datasource || d.Get("credential_name").(string) != "" {
			if general.Credentials != nil {
				rdsPutString(values, "credential_name", general.Credentials.CredentialName)
			}
			if general.AmazonInstanceInfo != nil && general.AmazonInstanceInfo.UseIamRole != nil {
				values["use_iam_role"] = *general.AmazonInstanceInfo.UseIamRole
			}
		}
	}
	return setRDSState(d, values)
}

func resourceDeleteRDSInstance(d *schema.ResourceData, m interface{}) error {
	// Verify workload identity before sending a destructive generic instance request.
	if err := resourceReadRDSInstance(d, m); err != nil || d.Id() == "" {
		return err
	}
	if _, err := handler.CvDeleteRDSInstance(d.Id()); err != nil && !strings.HasPrefix(err.Error(), "status: 404,") {
		return fmt.Errorf("delete RDS instance %s: %w", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func rdsResourceAbsent(d *schema.ResourceData, datasource bool, err error, count int) bool {
	if !datasource && ((err != nil && strings.HasPrefix(err.Error(), "status: 404,")) || (err == nil && count == 0)) {
		d.SetId("")
		return true
	}
	return false
}

func rdsPutString(values map[string]interface{}, key string, value *string) {
	if value != nil {
		values[key] = *value
	}
}

func cloudOptionalBool(d *schema.ResourceData, key string) *bool {
	if value, ok := d.GetOkExists(key); ok {
		result := value.(bool)
		return &result
	}
	return nil
}

func cloudOptionalString(d *schema.ResourceData, key string) *string {
	if value, ok := d.GetOkExists(key); ok {
		result := value.(string)
		return &result
	}
	return nil
}

func setRDSState(d *schema.ResourceData, values map[string]interface{}) error {
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("set RDS %s: %w", key, err)
		}
	}
	return nil
}

func validateRDSStringEnum(allowed ...string) schema.SchemaValidateFunc {
	return func(value interface{}, key string) ([]string, []error) {
		for _, candidate := range allowed {
			if value.(string) == candidate {
				return nil, nil
			}
		}
		return nil, []error{fmt.Errorf("%s must be one of %v", key, allowed)}
	}
}
