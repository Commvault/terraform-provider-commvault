package commvault

import (
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceS3ObjectStorage() *schema.Resource {
	return &schema.Resource{
		Create:   resourceCreateS3ObjectStorage,
		Read:     resourceReadS3ObjectStorage,
		Update:   resourceUpdateS3ObjectStorage,
		Delete:   resourceDeleteS3ObjectStorage,
		Importer: &schema.ResourceImporter{State: schema.ImportStatePassthrough},
		Schema:   s3ObjectStorageSchema(),
	}
}

func s3ObjectStorageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"client_name":   {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validateS3NonEmptyString},
		"instance_name": {Type: schema.TypeString, Required: true, ValidateFunc: validateS3NonEmptyString},
		"client_id":     {Type: schema.TypeInt, Computed: true},
		"instance_id":   {Type: schema.TypeInt, Computed: true},
		"comm_cell_id":  {Type: schema.TypeInt, Optional: true, Default: 2, ForceNew: true, ValidateFunc: validateS3IntRange(0, 2147483647)},
		"plan_name":     {Type: schema.TypeString, Required: true, ValidateFunc: validateS3NonEmptyString},
		"host_url":      {Type: schema.TypeString, Optional: true, Default: "s3.amazonaws.com", ValidateFunc: validateS3NonEmptyString},
		"credential_type": {
			Type: schema.TypeString, Optional: true, Default: "AMAZON_S3",
			ValidateFunc: validateS3CredentialType,
		},
		"credential_name":                   {Type: schema.TypeString, Optional: true},
		"use_iam_role":                      {Type: schema.TypeBool, Computed: true, Description: "Derived from credential_type: true for AMAZON_IAM_ROLE."},
		"number_of_backup_streams":          {Type: schema.TypeInt, Optional: true, Default: 4, ValidateFunc: validateS3IntRange(1, 2147483647)},
		"member_servers":                    {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
		"region_name":                       {Type: schema.TypeString, Optional: true, Description: "Region for hosted infrastructure; allows an empty member_servers list."},
		"validate_region_info":              {Type: schema.TypeBool, Optional: true, Default: false},
		"is_customer_managed_access_node":   {Type: schema.TypeBool, Optional: true, Computed: true},
		"object_storage_indexing_schema_v2": {Type: schema.TypeBool, Optional: true, Computed: true},
		"object_storage_multinode":          {Type: schema.TypeInt, Optional: true, Computed: true, ValidateFunc: validateS3IntRange(0, 2147483647)},
		"is_cloud_encryption_key_set":       {Type: schema.TypeBool, Optional: true, Default: false},
		"use_resource_pool_info":            {Type: schema.TypeBool, Optional: true, Default: false, ForceNew: true},
		"company_name":                      {Type: schema.TypeString, Optional: true, Default: "Commcell", Description: "Company name used when deconfiguring the instance."},
		"multi_commcell_id":                 {Type: schema.TypeInt, Optional: true, Default: 0, ValidateFunc: validateS3IntRange(0, 2147483647)},
	}
}

func buildS3ObjectStorageProperties(d *schema.ResourceData) (*handler.MsgS3InstanceProperties, error) {
	credentialType := d.Get("credential_type").(string)
	credentialName := d.Get("credential_name").(string)
	if credentialType != "AMAZON_IAM_ROLE" && credentialName == "" {
		return nil, fmt.Errorf("credential_name is required for credential_type %s", credentialType)
	}
	if credentialType == "AMAZON_IAM_ROLE" && credentialName != "" {
		return nil, fmt.Errorf("credential_name must be omitted for AMAZON_IAM_ROLE")
	}
	servers := make([]handler.MsgS3MemberServer, 0)
	for _, value := range d.Get("member_servers").([]interface{}) {
		name, entityType := value.(string), 3
		if strings.TrimSpace(name) == "" {
			return nil, fmt.Errorf("member_servers must not contain empty access-node names")
		}
		servers = append(servers, handler.MsgS3MemberServer{Client: &handler.MsgS3Entity{ClientName: &name, Type: &entityType}})
	}
	if len(servers) == 0 && d.Get("region_name").(string) == "" {
		return nil, fmt.Errorf("member_servers must contain an access node unless region_name is set for hosted infrastructure")
	}
	if len(servers) == 0 && credentialType == "AMAZON_IAM_ROLE" {
		return nil, fmt.Errorf("AMAZON_IAM_ROLE requires member_servers")
	}
	applicationID, instanceType, displayName := 134, "AMAZON_S3", "Amazon S3"
	useIAMRole := credentialType == "AMAZON_IAM_ROLE"
	general := &handler.MsgS3GeneralCloudProperties{
		ObjectStorageIndexingSchemaV2: cloudOptionalBool(d, "object_storage_indexing_schema_v2"),
		ObjectStorageMultinode:        intPtrFromResource(d, "object_storage_multinode"),
		NumberOfBackupStreams:         intPtrFromResource(d, "number_of_backup_streams"),
		MemberServers:                 servers,
	}
	if credentialName != "" {
		general.Credentials = &handler.MsgCredentialInfo{CredentialName: &credentialName}
	}
	objectStorage := &handler.MsgS3ObjectStorageInstance{
		IsCustomerManagedAccessNode: cloudOptionalBool(d, "is_customer_managed_access_node"),
		IsCloudEncryptionKeySet:     boolPtrFromResource(d, "is_cloud_encryption_key_set"),
	}
	if region := stringPtrFromResource(d, "region_name"); region != nil {
		objectStorage.RegionInfo = &handler.MsgS3RegionInfo{RegionName: region}
	}
	return &handler.MsgS3InstanceProperties{
		Instance: &handler.MsgS3Entity{
			InstanceName:  stringPtrFromResource(d, "instance_name"),
			ClientName:    stringPtrFromResource(d, "client_name"),
			ApplicationId: &applicationID,
			CommCellId:    intPtrFromResource(d, "comm_cell_id"),
		},
		CloudAppsInstance: &handler.MsgS3CloudAppsInstance{
			InstanceType: &instanceType, InstanceTypeDisplayName: &displayName,
			GeneralCloudProperties: general, ObjectStorageInstance: objectStorage,
			S3Instance:         &handler.MsgS3Connection{HostURL: stringPtrFromResource(d, "host_url"), UseIamRole: &useIAMRole},
			CredentialType:     &credentialType,
			ValidateRegionInfo: boolPtrFromResource(d, "validate_region_info"),
		},
		PlanEntity:          &handler.MsgS3Plan{PlanName: stringPtrFromResource(d, "plan_name")},
		UseResourcePoolInfo: boolPtrFromResource(d, "use_resource_pool_info"),
	}, nil
}

func resourceCreateS3ObjectStorage(d *schema.ResourceData, m interface{}) error {
	properties, err := buildS3ObjectStorageProperties(d)
	if err != nil {
		return err
	}
	clientType, instanceType := 15, "AMAZON_S3"
	clientName, plan := properties.Instance.ClientName, properties.PlanEntity
	properties.Instance.ClientName, properties.PlanEntity = nil, nil
	results, err := handler.CvCreateS3ObjectStorage(handler.MsgCreateS3ObjectStorageRequest{
		Entity: &handler.MsgS3Entity{ClientName: clientName},
		ClientInfo: &handler.MsgS3ClientInfo{
			ClientType: &clientType, Plan: plan,
			CloudClonnectorProperties: &handler.MsgS3CloudConnector{InstanceType: &instanceType, Instance: properties},
		},
	})
	if err != nil {
		return fmt.Errorf("create S3 object storage: %w", err)
	}
	for _, result := range results {
		if result.Entity != nil && result.Entity.InstanceId != nil && *result.Entity.InstanceId > 0 {
			d.SetId(strconv.Itoa(*result.Entity.InstanceId))
			break
		}
	}
	if d.Id() == "" {
		entity, err := handler.CvFindS3ObjectStorage(d.Get("client_name").(string), d.Get("instance_name").(string))
		if err != nil {
			return fmt.Errorf("S3 client was created but instance ID lookup failed; import the instance before retrying: %w", err)
		}
		d.SetId(strconv.Itoa(*entity.InstanceId))
	}
	return resourceReadS3ObjectStorage(d, m)
}

func resourceReadS3ObjectStorage(d *schema.ResourceData, m interface{}) error {
	return readS3ObjectStorage(d, false)
}

func readS3ObjectStorage(d *schema.ResourceData, datasource bool) error {
	if err := validateS3ID(d.Id()); err != nil {
		return err
	}
	properties, err := handler.CvGetS3ObjectStorage(d.Id())
	if err != nil {
		if !datasource && strings.HasPrefix(err.Error(), "status: 404,") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("read S3 object storage %s: %w", d.Id(), err)
	}
	if len(properties) == 0 && !datasource {
		d.SetId("")
		return nil
	}
	if len(properties) != 1 {
		return fmt.Errorf("expected one S3 instance for ID %s, got %d", d.Id(), len(properties))
	}
	prop := properties[0]
	if prop.Instance == nil || prop.CloudAppsInstance == nil || prop.CloudAppsInstance.S3Instance == nil {
		return fmt.Errorf("instance %s does not contain S3 object storage properties", d.Id())
	}
	if prop.Instance.ApplicationId == nil || *prop.Instance.ApplicationId != 134 {
		return fmt.Errorf("instance %s is not a Cloud Apps instance", d.Id())
	}
	cloud := prop.CloudAppsInstance
	if cloud.InstanceType != nil && *cloud.InstanceType != "AMAZON_S3" {
		return fmt.Errorf("instance %s is not AMAZON_S3", d.Id())
	}
	values := map[string]interface{}{}
	id, _ := strconv.Atoi(d.Id())
	values["instance_id"] = id
	s3PutString(values, "client_name", prop.Instance.ClientName)
	s3PutString(values, "instance_name", prop.Instance.InstanceName)
	s3PutInt(values, "client_id", prop.Instance.ClientId)
	s3PutInt(values, "comm_cell_id", prop.Instance.CommCellId)
	if prop.UseResourcePoolInfo != nil {
		values["use_resource_pool_info"] = *prop.UseResourcePoolInfo
	}
	s3PutString(values, "host_url", cloud.S3Instance.HostURL)
	s3PutString(values, "credential_type", cloud.CredentialType)
	if cloud.S3Instance.UseIamRole != nil {
		values["use_iam_role"] = *cloud.S3Instance.UseIamRole
	}
	if cloud.ValidateRegionInfo != nil {
		values["validate_region_info"] = *cloud.ValidateRegionInfo
	}
	if general := cloud.GeneralCloudProperties; general != nil {
		s3PutInt(values, "object_storage_multinode", general.ObjectStorageMultinode)
		if general.ObjectStorageIndexingSchemaV2 != nil {
			values["object_storage_indexing_schema_v2"] = *general.ObjectStorageIndexingSchemaV2
		}
		s3PutInt(values, "number_of_backup_streams", general.NumberOfBackupStreams)
		servers := make([]string, 0, len(general.MemberServers))
		for _, server := range general.MemberServers {
			if server.Client == nil || server.Client.ClientName == nil {
				return fmt.Errorf("S3 instance %s returned a member server without clientName", d.Id())
			}
			servers = append(servers, *server.Client.ClientName)
		}
		values["member_servers"] = servers
		values["credential_name"] = ""
		if general.Credentials != nil {
			s3PutString(values, "credential_name", general.Credentials.CredentialName)
		}
	}
	values["region_name"] = ""
	if object := cloud.ObjectStorageInstance; object != nil {
		if object.IsCustomerManagedAccessNode != nil {
			values["is_customer_managed_access_node"] = *object.IsCustomerManagedAccessNode
		}
		if object.IsCloudEncryptionKeySet != nil {
			values["is_cloud_encryption_key_set"] = *object.IsCloudEncryptionKeySet
		}
		if object.RegionInfo != nil {
			s3PutString(values, "region_name", object.RegionInfo.RegionName)
		}
	}
	if prop.PlanEntity != nil {
		s3PutString(values, "plan_name", prop.PlanEntity.Name)
		s3PutString(values, "plan_name", prop.PlanEntity.PlanName)
	}
	return setS3State(d, values)
}

func resourceUpdateS3ObjectStorage(d *schema.ResourceData, m interface{}) error {
	if !d.HasChanges("plan_name", "host_url", "credential_type", "credential_name", "number_of_backup_streams",
		"member_servers", "region_name", "validate_region_info", "is_cloud_encryption_key_set", "instance_name",
		"is_customer_managed_access_node", "object_storage_indexing_schema_v2", "object_storage_multinode") {
		return nil
	}
	properties, err := buildS3ObjectStorageProperties(d)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(d.Id())
	if err != nil || id <= 0 {
		return fmt.Errorf("invalid S3 instance ID %q", d.Id())
	}
	properties.Instance.InstanceId = &id
	properties.Instance.ClientId = intPtrFromResource(d, "client_id")
	clientName := properties.Instance.ClientName
	properties.Instance.ClientName, properties.Instance.InstanceName, properties.Instance.CommCellId = nil, nil, nil
	// This flag is create-only in the supplied API spec.
	properties.UseResourcePoolInfo = nil
	request := handler.MsgModifyS3ObjectStorageRequest{
		InstanceProperties: properties,
		Association:        &handler.MsgS3Association{Entity: []handler.MsgS3Entity{{ClientName: clientName}}},
	}
	if d.HasChange("instance_name") {
		request.NewName = stringPtrFromResource(d, "instance_name")
	}
	_, err = handler.CvModifyS3ObjectStorage(d.Id(), request)
	if err != nil {
		return fmt.Errorf("modify S3 object storage: %w", err)
	}
	if request.NewName != nil {
		return refreshCloudInstanceRename(d, m, *request.NewName, resourceReadS3ObjectStorage)
	}
	return resourceReadS3ObjectStorage(d, m)
}

func resourceDeleteS3ObjectStorage(d *schema.ResourceData, m interface{}) error {
	if err := validateS3ID(d.Id()); err != nil {
		return err
	}
	applicationID, appName := 134, "Cloud Apps"
	_, err := handler.CvDeleteS3ObjectStorage(d.Id(), handler.MsgDeleteS3ObjectStorageRequest{
		Instance: &handler.MsgS3Entity{
			ClientName: stringPtrFromResource(d, "client_name"), InstanceName: stringPtrFromResource(d, "instance_name"),
			ApplicationId: &applicationID, AppName: &appName,
			EntityInfo: &handler.MsgS3EntityInfo{
				CompanyName: stringPtrFromResource(d, "company_name"), MultiCommcellId: intPtrFromResource(d, "multi_commcell_id"),
			},
		},
	})
	if err != nil && !strings.HasPrefix(err.Error(), "status: 404,") {
		return fmt.Errorf("deconfigure S3 object storage: %w", err)
	}
	d.SetId("")
	return nil
}

func validateS3ID(value string) error {
	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 {
		return fmt.Errorf("S3 entity ID must be a positive integer, got %q", value)
	}
	return nil
}

func s3PutString(values map[string]interface{}, key string, value *string) {
	if value != nil {
		values[key] = *value
	}
}

func s3PutInt(values map[string]interface{}, key string, value *int) {
	if value != nil {
		values[key] = *value
	}
}

func setS3State(d *schema.ResourceData, values map[string]interface{}) error {
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("set S3 state field %s: %w", key, err)
		}
	}
	return nil
}

func validateS3NonEmptyString(value interface{}, key string) ([]string, []error) {
	return validateNonEmptyString(value, key)
}

func validateNonEmptyString(value interface{}, key string) ([]string, []error) {
	if strings.TrimSpace(value.(string)) == "" {
		return nil, []error{fmt.Errorf("%s must not be empty", key)}
	}
	return nil, nil
}

func validateS3IntRange(min, max int) schema.SchemaValidateFunc {
	return validateIntRange(min, max)
}

func validateIntRange(min, max int) schema.SchemaValidateFunc {
	return func(value interface{}, key string) ([]string, []error) {
		number := value.(int)
		if number < min || number > max {
			return nil, []error{fmt.Errorf("%s must be between %d and %d", key, min, max)}
		}
		return nil, nil
	}
}

func validateS3CredentialType(value interface{}, key string) ([]string, []error) {
	switch value.(string) {
	case "AMAZON_S3", "AMAZON_IAM_ROLE", "AMAZON_STS_IAM_ROLE":
		return nil, nil
	default:
		return nil, []error{fmt.Errorf("%s must be AMAZON_S3, AMAZON_IAM_ROLE, or AMAZON_STS_IAM_ROLE", key)}
	}
}
