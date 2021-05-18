package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAWSStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreateAWSStorage,
		Read:   resourceReadAWSStorage,
		Update: resourceUpdateAWSStorage,
		Delete: resourceDeleteAWSStorage,

		Schema: map[string]*schema.Schema{
			"storage_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mediaagent": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"secret_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"credentials_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ddb_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"company_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceCreateAWSStorage(d *schema.ResourceData, m interface{}) error {
	var createStorageReq handler.CreateStorageReq
	var subStoreList handler.SubStoreList
	var maInfoList handler.MaInfoList
	createStorageReq.StoragePolicyName = d.Get("storage_name").(string)
	createStorageReq.CopyName = d.Get("storage_name").(string) + "_Primary"
	createStorageReq.Type = "CVA_REGULAR_SP"
	createStorageReq.NumberOfCopies = 1
	createStorageReq.StoragePolicyCopyInfo.CopyType = "SYNCHRONOUS"
	createStorageReq.StoragePolicyCopyInfo.IsDefault = "SET_TRUE"
	createStorageReq.StoragePolicyCopyInfo.Active = "SET_TRUE"
	createStorageReq.StoragePolicyCopyInfo.Library.LibraryName = d.Get("bucket").(string)
	if d.Get("ddb_location").(string) == "" {
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.GlobalStoragePolicy = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.CopyFlags.PreserveEncryptionModeAsInSource = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.ExtendedFlags.GlobalStoragePolicy = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.EnableDeduplication = "SET_FALSE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.EnableDASHFull = "SET_FALSE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.HostGlobalDedupStore = "SET_FALSE"
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.BlockLevelDedup = "SET_FALSE"
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.EnableGlobalDeduplication = "SET_FALSE"
		maInfoList.SubStoreList = append(maInfoList.SubStoreList, subStoreList)
		maInfoList.MediaAgent.MediaAgentName = d.Get("mediaagent").(string)
		createStorageReq.StoragePolicyCopyInfo.DDBPartitionInfo.MaInfoList = append(createStorageReq.StoragePolicyCopyInfo.DDBPartitionInfo.MaInfoList, maInfoList)
	} else {
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.GlobalStoragePolicy = "SET_FALSE"
		createStorageReq.StoragePolicyCopyInfo.CopyFlags.PreserveEncryptionModeAsInSource = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.ExtendedFlags.GlobalStoragePolicy = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.BlockLevelDedup = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.StoragePolicyFlags.EnableGlobalDeduplication = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.EnableDeduplication = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.EnableDASHFull = "SET_TRUE"
		createStorageReq.StoragePolicyCopyInfo.DedupeFlags.HostGlobalDedupStore = "SET_TRUE"
		subStoreList.AccessPath.Path = d.Get("ddb_location").(string)
		subStoreList.DiskFreeThresholdMB = 5120
		subStoreList.DiskFreeWarningThreshholdMB = 10240
		maInfoList.SubStoreList = append(maInfoList.SubStoreList, subStoreList)
		maInfoList.MediaAgent.MediaAgentName = d.Get("mediaagent").(string)
		createStorageReq.StoragePolicyCopyInfo.DDBPartitionInfo.MaInfoList = append(createStorageReq.StoragePolicyCopyInfo.DDBPartitionInfo.MaInfoList, maInfoList)
		createStorageReq.StoragePolicyCopyInfo.DDBPartitionInfo.SidbStoreInfo.NumSIDBStore = 2
	}
	createStorageReq.StoragePolicyCopyInfo.MediaAgent.MediaAgentName = d.Get("mediaagent").(string)
	createStorageReq.StoragePolicyCopyInfo.RetentionRules.RetentionFlags.EnableDataAging = "SET_TRUE"
	createStorageReq.StoragePolicyCopyInfo.IsFromGui = false
	createStorageReq.StoragePolicyCopyInfo.RetentionRules.RetainArchiverDataForDays = -1
	createStorageReq.StoragePolicyCopyInfo.RetentionRules.RetainBackupDataForCycles = -1
	createStorageReq.StoragePolicyCopyInfo.RetentionRules.RetainBackupDataForDays = -1
	createStorageReq.StoragePolicyCopyInfo.NumberOfStreamsToCombine = 1

	var storage handler.Storage
	storage.MediaAgent.MediaAgentName = d.Get("mediaagent").(string)
	storage.Path = d.Get("bucket").(string)
	storage.DeviceType = 2
	storage.MetallicStorageInfo.StorageClass = append(storage.MetallicStorageInfo.StorageClass, "CONTAINER_DEFAULT")
	storage.MetallicStorageInfo.Replication = append(storage.MetallicStorageInfo.Replication, "NONE")
	if d.Get("credentials_name").(string) != "" {
		storage.Credentials.UserName = d.Get("service_host").(string) + "@0//__CVCRED__"
		storage.Credentials.Password = "OTg3NjU0MzIx"
		storage.SavedCredential.CredentialName = d.Get("credentials_name").(string)
	} else if d.Get("access_key_id").(string) != "" && d.Get("secret_access_key").(string) != "" {
		storage.Credentials.UserName = d.Get("service_host").(string) + "@0//" + d.Get("access_key_id").(string)
		storage.Credentials.Password = d.Get("secret_access_key").(string)
		storage.SavedCredential.CredentialName = ""
	} else {
		return fmt.Errorf("Saved Credentails names or access key, secrect key are missing")
	}
	createStorageReq.Storage = append(createStorageReq.Storage, storage)
	storageResp := handler.CreateStorage(createStorageReq, d.Get("company_id").(int))
	errorCode := storageResp.Error.ErrorCode
	if errorCode != 0 {
		return fmt.Errorf("AWS Storage creation failed")
	}
	storageID := strconv.Itoa(storageResp.ArchiveGroupCopy.StoragePolicyID)
	d.SetId(storageID)
	return resourceReadDiskStorage(d, m)
}

func resourceReadAWSStorage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceUpdateAWSStorage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeleteAWSStorage(d *schema.ResourceData, m interface{}) error {
	storageID := d.Id()
	genericResp := handler.DeleteStorage(storageID)
	if genericResp.ErrorCode != 0 {
		return fmt.Errorf("Error in deletion of AWS Storage")
	}
	d.SetId("")
	return nil
}
