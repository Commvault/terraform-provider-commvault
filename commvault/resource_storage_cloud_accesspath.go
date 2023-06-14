package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorage_Cloud_AccessPath() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateStorage_Cloud_AccessPath,
        Read:   resourceReadStorage_Cloud_AccessPath,
        Update: resourceUpdateStorage_Cloud_AccessPath,
        Delete: resourceDeleteStorage_Cloud_AccessPath,

        Schema: map[string]*schema.Schema{
            "cloudstorageid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "Id of cloud Storage",
            },
            "bucketid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "Id of Bucket",
            },
            "mediaagent": {
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
        },
    }
}

func resourceCreateStorage_Cloud_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath
    var response_id = strconv.Itoa(0)
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagent = build_storage_cloud_accesspath_msgidname(d, val.([]interface{}))
    }
    var req = handler.MsgCreateAccessPathForBucketOfCloudStorageRequest{MediaAgent:t_mediaagent}
    resp, err := handler.CvCreateAccessPathForBucketOfCloudStorage(req, strconv.Itoa(d.Get("cloudstorageid").(int)), strconv.Itoa(d.Get("bucketid").(int)))
    if err != nil {
        return fmt.Errorf("operation [CreateAccessPathForBucketOfCloudStorage] failed, Error %s", err)
    }
    response_id = handler.RetrieveBucketAccessPathId(req, resp, d, m)
    if response_id == "0" {
        return fmt.Errorf("operation [CreateAccessPathForBucketOfCloudStorage] failed")
    } else {
        d.SetId(response_id)
        return resourceReadStorage_Cloud_AccessPath(d, m)
    }
}

func resourceReadStorage_Cloud_AccessPath(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceUpdateStorage_Cloud_AccessPath(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceDeleteStorage_Cloud_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath/{accessPathId}
    _, err := handler.CvDeleteAccessPathForBucketOfCloudStorage(strconv.Itoa(d.Get("cloudstorageid").(int)), strconv.Itoa(d.Get("bucketid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteAccessPathForBucketOfCloudStorage] failed, Error %s", err)
    }
    return nil
}

func build_storage_cloud_accesspath_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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
