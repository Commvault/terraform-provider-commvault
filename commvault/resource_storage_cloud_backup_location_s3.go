package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorage_Cloud_Backup_Location_S3() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateStorage_Cloud_Backup_Location_S3,
        Read:   resourceReadStorage_Cloud_Backup_Location_S3,
        Update: resourceUpdateStorage_Cloud_Backup_Location_S3,
        Delete: resourceDeleteStorage_Cloud_Backup_Location_S3,

        Schema: map[string]*schema.Schema{
            "cloudstorageid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "",
            },
            "bucket": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of bucket",
            },
            "mediaagent": {
                Type:        schema.TypeList,
                Required:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "storageclass": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Appropriate storage class for your account [Standard, Reduced Redundancy Storage, Standard - Infrequent access, One zone - Infrequent access, Intelligent tiering, S3 Glacier, Standard/Glacier (Combined Storage Tiers), Standard-IA/Glacier (Combined Storage Tiers), One Zone-IA/Glacier (Combined Storage Tiers), Intelligent-Tiering/Glacier (Combined Storage Tiers), S3 Glacier Deep Archive, Standard/Deep Archive (Combined Storage Tiers), Standard-IA/Deep Archive (Combined Storage Tiers), One Zone-IA/Deep Archive (Combined Storage Tiers), Intelligent-Tiering/Deep Archive (Combined Storage Tiers)]",
            },
            "servicehost": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "IP address or fully qualified domain name or URL for the cloud library based on cloud vendor",
            },
            "credentials": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "cloudtype": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of cloud vendor [Amazon S3]",
            },
            "arnrole": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Needed for AWS STS assume role and AWS STS assume role with IAM role policy",
            },
            "authentication": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Authentication type for the cloud storage server [Access and secret keys, AWS IAM role policy, AWS STS assume role, AWS STS assume role with IAM role policy]",
            },
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Password for proxy configuration (Should be in Base64 format)",
            },
            "port": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "Port for proxy configuration",
            },
            "proxyaddress": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "If the MediaAgent accesses the mount path using a proxy then proxy server address needs to be provided. If you want to remove proxy information, pass empty string in proxyAddress.",
            },
            "username": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Username for proxy configuration",
            },
            "enable": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "When true, means mount path is enabled",
            },
            "prepareforretirement": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "When true, the deduplicated blocks in the mount path will not be referenced when there are multiple mount paths in the library.",
            },
            "storageacceleratorcredentials": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "preventnewdatawritestobackuplocation": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "When true, prevents new data writes to backup location by changing number of writers to zero",
            },
        },
    }
}

func resourceCreateStorage_Cloud_Backup_Location_S3(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Cloud/{cloudStorageId}/Bucket
    var response_id = strconv.Itoa(0)
    var t_bucket *string
    if val, ok := d.GetOk("bucket"); ok {
        t_bucket = handler.ToStringValue(val, false)
    }
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagent = build_storage_cloud_backup_location_s3_msgidname(d, val.([]interface{}))
    }
    var t_storageclass *string
    if val, ok := d.GetOk("storageclass"); ok {
        t_storageclass = handler.ToStringValue(val, false)
    }
    var t_servicehost *string
    if val, ok := d.GetOk("servicehost"); ok {
        t_servicehost = handler.ToStringValue(val, false)
    }
    var t_credentials *handler.MsgIdName
    if val, ok := d.GetOk("credentials"); ok {
        t_credentials = build_storage_cloud_backup_location_s3_msgidname(d, val.([]interface{}))
    }
    var t_cloudtype *string
    if val, ok := d.GetOk("cloudtype"); ok {
        t_cloudtype = handler.ToStringValue(val, false)
    }
    var t_arnrole *string
    if val, ok := d.GetOk("arnrole"); ok {
        t_arnrole = handler.ToStringValue(val, false)
    }
    var t_authentication *string
    if val, ok := d.GetOk("authentication"); ok {
        t_authentication = handler.ToStringValue(val, false)
    }
    var t_password *string
    if val, ok := d.GetOk("password"); ok {
        t_password = handler.ToStringValue(val, false)
    }
    var t_port *int
    if val, ok := d.GetOk("port"); ok {
        t_port = handler.ToIntValue(val, false)
    }
    var t_proxyaddress *string
    if val, ok := d.GetOk("proxyaddress"); ok {
        t_proxyaddress = handler.ToStringValue(val, false)
    }
    var t_username *string
    if val, ok := d.GetOk("username"); ok {
        t_username = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateBucketforCloudStorageS3Request{Bucket:t_bucket, MediaAgent:t_mediaagent, StorageClass:t_storageclass, ServiceHost:t_servicehost, Credentials:t_credentials, CloudType:t_cloudtype, ArnRole:t_arnrole, Authentication:t_authentication, Password:t_password, Port:t_port, ProxyAddress:t_proxyaddress, Username:t_username}
    resp, err := handler.CvCreateBucketforCloudStorageS3(req, strconv.Itoa(d.Get("cloudstorageid").(int)))
    if err != nil {
        return fmt.Errorf("operation [CreateBucketforCloudStorageS3] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateBucketforCloudStorageS3] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateStorage_Cloud_Backup_Location_S3(d, m)
    }
}

func resourceReadStorage_Cloud_Backup_Location_S3(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    resp, err := handler.CvGetBucketDetailsOfCloudStorage(strconv.Itoa(d.Get("cloudstorageid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetBucketDetailsOfCloudStorage] failed, Error %s", err)
    }
    if rtn, ok := serialize_storage_cloud_backup_location_s3_msgidname(d, resp.Bucket); ok {
        d.Set("bucket", rtn)
    } else {
        d.Set("bucket", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateStorage_Cloud_Backup_Location_S3(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    var t_enable *bool
    if d.HasChange("enable") {
        val := d.Get("enable")
        t_enable = handler.ToBooleanValue(val, false)
    }
    var t_prepareforretirement *bool
    if d.HasChange("prepareforretirement") {
        val := d.Get("prepareforretirement")
        t_prepareforretirement = handler.ToBooleanValue(val, false)
    }
    var t_storageacceleratorcredentials *handler.MsgIdName
    if d.HasChange("storageacceleratorcredentials") {
        val := d.Get("storageacceleratorcredentials")
        t_storageacceleratorcredentials = build_storage_cloud_backup_location_s3_msgidname(d, val.([]interface{}))
    }
    var t_preventnewdatawritestobackuplocation *bool
    if d.HasChange("preventnewdatawritestobackuplocation") {
        val := d.Get("preventnewdatawritestobackuplocation")
        t_preventnewdatawritestobackuplocation = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgModifyBucketOfCloudStorageRequest{Enable:t_enable, PrepareForRetirement:t_prepareforretirement, StorageAcceleratorCredentials:t_storageacceleratorcredentials, PreventNewDataWritesToBackupLocation:t_preventnewdatawritestobackuplocation}
    _, err := handler.CvModifyBucketOfCloudStorage(req, strconv.Itoa(d.Get("cloudstorageid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyBucketOfCloudStorage] failed, Error %s", err)
    }
    return resourceReadStorage_Cloud_Backup_Location_S3(d, m)
}

func resourceCreateUpdateStorage_Cloud_Backup_Location_S3(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    var execUpdate bool = false
    var t_enable *bool
    if val, ok := d.GetOk("enable"); ok {
        t_enable = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_prepareforretirement *bool
    if val, ok := d.GetOk("prepareforretirement"); ok {
        t_prepareforretirement = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_storageacceleratorcredentials *handler.MsgIdName
    if val, ok := d.GetOk("storageacceleratorcredentials"); ok {
        t_storageacceleratorcredentials = build_storage_cloud_backup_location_s3_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_preventnewdatawritestobackuplocation *bool
    if val, ok := d.GetOk("preventnewdatawritestobackuplocation"); ok {
        t_preventnewdatawritestobackuplocation = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyBucketOfCloudStorageRequest{Enable:t_enable, PrepareForRetirement:t_prepareforretirement, StorageAcceleratorCredentials:t_storageacceleratorcredentials, PreventNewDataWritesToBackupLocation:t_preventnewdatawritestobackuplocation}
        _, err := handler.CvModifyBucketOfCloudStorage(req, strconv.Itoa(d.Get("cloudstorageid").(int)), d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyBucketOfCloudStorage] failed, Error %s", err)
        }
    }
    return resourceReadStorage_Cloud_Backup_Location_S3(d, m)
}

func resourceDeleteStorage_Cloud_Backup_Location_S3(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    _, err := handler.CvDeleteBucketOfCloudStorage(strconv.Itoa(d.Get("cloudstorageid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteBucketOfCloudStorage] failed, Error %s", err)
    }
    return nil
}

func build_storage_cloud_backup_location_s3_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsgIdName{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func serialize_storage_cloud_backup_location_s3_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //Msgnull
    //MsgIdName
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if added {
        return val, true
    } else {
        return nil, false
    }
}
