package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func RetrieveBackupLocationAccessPathId(req MsgAddMediaAgentRequest, resp *MsgAddMediaAgentResponse, d *schema.ResourceData, m interface{}) string {
	id, _ := CvGetAccessPathForMediaAgent(strconv.Itoa(d.Get("storagepoolid").(int)), strconv.Itoa(d.Get("backuplocationid").(int)), *req.MediaAgents[0].Id)
	return strconv.Itoa(id)
}

func RetrieveBucketAccessPathId(req MsgCreateAccessPathForBucketOfCloudStorageRequest, resp *MsgCreateAccessPathForBucketOfCloudStorageResponse, d *schema.ResourceData, m interface{}) string {
	id, _ := CvGetCloudAccessPathForMediaAgent(strconv.Itoa(d.Get("cloudstorageid").(int)), strconv.Itoa(d.Get("bucketid").(int)), *req.MediaAgent.Id)
	return strconv.Itoa(id)
}

func ConfigureCredential_AWSWithRoleArn(req *MsgCreateCredentialAWSWithRoleArnRequest, d *schema.ResourceData, m interface{}) error {
	accountType := "CLOUD_ACCOUNT"
	req.AccountType = new(string)
	req.AccountType = &accountType

	vendorType := "AMAZON"
	req.VendorType = new(string)
	req.VendorType = &vendorType

	password := ""
	req.Password = new(string)
	req.Password = &password

	return nil
}

func UpdateCredential_AWSWithRoleArn(req *MsgUpdateCredentialAWSWithRoleArnRequest, d *schema.ResourceData, m interface{}) error {
	password := ""
	req.Password = new(string)
	req.Password = &password

	return nil
}

func ConfigureCredential_AWS(req *MsgCreateCredentialAWSRequest, d *schema.ResourceData, m interface{}) error {
	accountType := "CLOUD_ACCOUNT"
	req.AccountType = new(string)
	req.AccountType = &accountType

	vendorType := "AMAZON"
	req.VendorType = new(string)
	req.VendorType = &vendorType

	return nil
}

func UpdateCredential_AWS(req *MsgUpdateCredentialAWSRequest, d *schema.ResourceData, m interface{}) error {
	return nil
}

func ConfigureCredential_Azure(req *MsgCreateCredentialAzureRequest, d *schema.ResourceData, m interface{}) error {
	accountType := "CLOUD_ACCOUNT"
	req.AccountType = new(string)
	req.AccountType = &accountType

	vendorType := "MICROSOFT_AZURE_TYPE"
	req.VendorType = new(string)
	req.VendorType = &vendorType

	return nil
}

func UpdateCredential_Azure(req *MsgUpdateCredentialAzureRequest, d *schema.ResourceData, m interface{}) error {
	return nil
}

func ConfigureCredential_AzureWithTenantId(req *MsgCreateCredentialAzureWithTenantIdRequest, d *schema.ResourceData, m interface{}) error {
	accountType := "CLOUD_ACCOUNT"
	req.AccountType = new(string)
	req.AccountType = &accountType

	vendorType := "MICROSOFT_AZURE_TYPE"
	req.VendorType = new(string)
	req.VendorType = &vendorType

	req.Endpoints = build_endpoints()

	return nil
}

func UpdateCredential_AzureWithTenantId(req *MsgUpdateCredentialAzureWithTenantIdRequest, d *schema.ResourceData, m interface{}) error {
	return nil
}

func build_endpoints() *MsgAzureEndpoints {
	t_storage := "blob.core.windows.net"
	t_activedirectory := "https://login.microsoftonline.com/"
	t_resourcemanager := "https://management.azure.com/"
	return &MsgAzureEndpoints{Storage: &t_storage, ActiveDirectory: &t_activedirectory, ResourceManager: &t_resourcemanager}
}

func UpdateBackupDestinations(req *MsgCreateServerPlanRequest, d *schema.ResourceData, m interface{}) error {

	// if req.BackupDestinations != nil {
	// 	for i := range req.BackupDestinations {
	// 		req.BackupDestinations[i].BackupDestinationName = req.BackupDestinations[i].PlanBackupDestination.Name
	// 	}
	// }

	return nil
}

func UpdateUserRequest(req *MsgModifyUserRequest, d *schema.ResourceData, m interface{}) error {
	if d.HasChange("password") {
		var val string
		if os.Getenv("CV_TER_PASSWORD") != "" {
			val = os.Getenv("CV_TER_PASSWORD")
		} else {
			val = os.Getenv("CV_PASSWORD")
		}
		if val == "" {
			return fmt.Errorf("cannot change password without a provider password or the environment variable CV_TER_PASSWORD")
		}
		req.ValidationPassword = new(string)
		req.ValidationPassword = &val
	}
	return nil
}

func accessNodeStateCopy(d *schema.ResourceData) ([]map[string]interface{}, bool) {
	nodes := d.Get("accessnodes").(*schema.Set).List()

	val := make([]map[string]interface{}, len(nodes))

	for i := range nodes {
		val[i] = nodes[i].(map[string]interface{})
	}

	return val, true
}

func GetAccessNodes(d *schema.ResourceData, model *MsgaccessNodeListModel) ([]map[string]interface{}, bool) {

	if model == nil {
		return accessNodeStateCopy(d)
	}

	data := model.AccessNode

	if data == nil {
		return nil, false
	}

	val := make([]map[string]interface{}, 0)
	for i := range data {
		tmp := make(map[string]interface{})
		added := false
		if data[i].Id != nil {
			tmp["id"] = data[i].Id
			added = true
		}
		if data[i].Type != nil {
			tmp["type"] = data[i].Type
			added = true
		}
		if added {
			val = append(val, tmp)
		}
	}

	return val, true
}

func GetConsoleTypes(d *schema.ResourceData, values []MsgRestrictedConsoleTypesSet) ([]map[string]interface{}, bool) {
	items := make([]string, len(values))
	for a, raw_a := range values {
		items[a] = *raw_a.ConsoleType
	}

	data := make([]map[string]interface{}, 1)
	data[0] = map[string]interface{}{
		"consoletype": items,
	}

	return data, true
}

func SortPlanSchedules(d *schema.ResourceData, data []MsgPlanScheduleSet) []MsgPlanScheduleSet {
	if len(data) == 0 {
		return data
	}

	rpo, _ := d.Get("rpo").([]interface{})
	if len(rpo) == 0 {
		return data
	}
	t_rpo := rpo[0].(map[string]interface{})

	backupfrequency, h_backupfrequency := t_rpo["backupfrequency"].([]interface{})
	if !h_backupfrequency || len(backupfrequency) == 0 {
		return data
	}
	t_backupfrequency := backupfrequency[0].(map[string]interface{})

	schedules, h_schedules := t_backupfrequency["schedules"].([]interface{})
	if !h_schedules || len(schedules) == 0 {
		return data
	}

	curr_data := make([]MsgPlanScheduleSet, 0)
	missing_data := make([]string, 0)

	for _, iter_a := range schedules {
		raw_a := iter_a.(map[string]interface{})
		sched_name := raw_a["schedulename"].(string)
		tmp := nextPlanSchedules(sched_name, data)
		if tmp != nil {
			curr_data = append(curr_data, *tmp)
		} else {
			missing_data = append(missing_data, sched_name)
		}
	}

	for _, sched_name := range missing_data {
		tmp := nextPlanSchedules(sched_name, data)
		if tmp != nil {
			curr_data = append(curr_data, *tmp)
		}
	}

	return curr_data
}

func nextPlanSchedules(scheduleName string, data []MsgPlanScheduleSet) *MsgPlanScheduleSet {
	for _, iter_a := range data {
		if iter_a.ScheduleName != nil && *iter_a.ScheduleName == scheduleName {
			return &iter_a
		}
	}
	return nil
}

type BackupDestinationKey struct {
	Destination string
	Region      string
}

func GetBackupDestinationKey(data map[string]interface{}) BackupDestinationKey {
	var key BackupDestinationKey

	//key.Destination = data["backupdestinationname"].(string)

	backupdestination, h_backupdestination := data["planbackupdestination"].([]interface{})
	if !h_backupdestination || len(backupdestination) == 0 {
		return key
	}
	t_backupdestination := backupdestination[0].(map[string]interface{})

	key.Destination = t_backupdestination["name"].(string)

	region, h_region := data["region"].([]interface{})
	if !h_region || len(region) == 0 {
		return key
	}
	t_region := region[0].(map[string]interface{})

	key.Region = t_region["name"].(string)

	return key
}

func SortBackupDestinations(d *schema.ResourceData, data []MsgPlanBackupDestinationSet) []MsgPlanBackupDestinationSet {
	if len(data) == 0 {
		return data
	}

	destinations, _ := d.Get("backupdestinations").([]interface{})
	if len(destinations) == 0 {
		return data
	}

	curr_data := make([]MsgPlanBackupDestinationSet, 0)

	for _, iter_a := range destinations {
		raw_a := iter_a.(map[string]interface{})
		key := GetBackupDestinationKey(raw_a)
		tmp := nextBackupDestination(key, data)
		if tmp != nil {
			curr_data = append(curr_data, *tmp)
		}
	}

	// for _, iter_a := range data {
	// 	if !hasBackupDestination(iter_a, curr_data) {
	// 		curr_data = append(curr_data, iter_a)
	// 	}
	// }

	return curr_data
}

// func hasBackupDestination(destination MsgPlanBackupDestinationSet, data []MsgPlanBackupDestinationSet) bool {
// 	for _, iter_a := range data {
// 		if iter_a.PlanBackupDestination.Name == destination.PlanBackupDestination.Name {
// 			if iter_a.Region != nil && *iter_a.Region.Name != "" {
// 				//has region
// 				if destination.Region != nil && *destination.Region.Name != "" {
// 					return iter_a.Region.Name == destination.Region.Name
// 				}
// 			} else {
// 				//no region
// 				if destination.Region == nil || *destination.Region.Name == "" {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

func nextBackupDestination(key BackupDestinationKey, data []MsgPlanBackupDestinationSet) *MsgPlanBackupDestinationSet {
	for _, iter_a := range data {
		if key.Region == "" {
			if iter_a.PlanBackupDestination != nil && *iter_a.PlanBackupDestination.Name == key.Destination {
				return &iter_a
			}
		} else {
			if iter_a.PlanBackupDestination != nil && *iter_a.PlanBackupDestination.Name == key.Destination {
				if iter_a.Region != nil && *iter_a.Region.Name == key.Region {
					return &iter_a
				}
			}
		}
	}
	return nil
}
