package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAWSProtectionGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateAWSProtectionGroup,
        Read:   resourceReadAWSProtectionGroup,
        Update: resourceUpdateAWSProtectionGroup,
        Delete: resourceDeleteAWSProtectionGroup,

        Schema: map[string]*schema.Schema{
            "cloudconnection": {
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
            "workloads": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "workloadoptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "islongtermretention": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "credential": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: "",
                                                },
                                                "id": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "israpidrecovery": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "solution": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "workload": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "allcloudaccounts": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "type": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "[NONE, VMW, MICROSOFT, XEN_SERVER, AMAZON, AZURE, REDHAT, AZURE_V2, SCVMM, NUTANIX, ORACLE_VM, DOCKER, OPENSTACK, ORACLE_CLOUD, FUSIONCOMPUTE, VCLOUD, GOOGLE_CLOUD, AZURE_STACK, ALIBABA_CLOUD, ORACLE_CLOUD_INFRASTRUCTURE, KUBERNETES, REDHAT_OPENSHIFT, MONGODB_ATLAS, PROXMOX, KUBERNETES_AKS, AZURE_STACK_HCI, KUBERNETES_EKS, MORPHEUS, KUBERNETES_GKE, KUBERNETES_OKE, NUTANIX_PRISM_CENTRAL, APACHE_CLOUDSTACK, VMWARE_CLOUD_FOUNDATION]",
            },
            "content": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "guestcredentialassocid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Credential association ID given to link entity with credential id.",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "username": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "matchrule": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Description of the rule group",
                        },
                        "rules": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Operation type for VM rules/filters [CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, ENDS_WITH, EQUALS, STARTS_WITH]",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "name of the VM to be added as content",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME, VM_BADGE, REGION]",
                                    },
                                    "value": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "value for the few type of VM Content like powerstate",
                                    },
                                },
                            },
                        },
                        "existingcredential": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentialid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "credentialname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "contentfilter": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "guestcredentialassocid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Credential association ID given to link entity with credential id.",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "username": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "matchrule": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Description of the rule group",
                        },
                        "rules": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Operation type for VM rules/filters [CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, ENDS_WITH, EQUALS, STARTS_WITH]",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "name of the VM to be added as content",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME, VM_BADGE, REGION]",
                                    },
                                    "value": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "value for the few type of VM Content like powerstate",
                                    },
                                },
                            },
                        },
                        "existingcredential": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentialid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "credentialname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "cloudaccounts": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "List of cloud accounts",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "cloud account name",
                        },
                        "id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "cloud account id",
                        },
                        "uuid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The globally unique identifier for the account",
                        },
                        "email": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "email for cloud account",
                        },
                    },
                },
            },
            "awsaccountdiscoverydetails": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Particulars required to create hypervisors corresponding to accounts in the organization",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rolenamewithpath": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Name of IAM role to assume to access accounts in the organization",
                        },
                        "accountaccessnodes": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of servers and server groups to use to access the onboarded accounts",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "displayname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
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
                                    "accessnodetype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Type of access node item [SERVER, SERVER_GROUP]",
                                    },
                                },
                            },
                        },
                        "backupserviceaccountid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Id of the account that hosts the Commvault backup infrastructure",
                        },
                    },
                },
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of protection group",
            },
            "deleteactions": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "When a protection group is deleted, specify what actions should be taken regarding the entities that were created by the protection group",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "deconfigureresources": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Do you want to deconfigure resources created by protection group",
                        },
                        "decoupleresources": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Do you want to disassociate accounts from Protection Group and backup keep running",
                        },
                    },
                },
            },
            "filterworkloads": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "workloadoptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "islongtermretention": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "credential": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: "",
                                                },
                                                "id": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "israpidrecovery": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "solution": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "workload": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "plan": {
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
            "backuponpostcreation": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "contentoperationtype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Specifies the operation to be performed on ruleGroups list [OVERWRITE, ADD, MODIFY, DELETE]",
            },
            "filtercontent": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "guestcredentialassocid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Credential association ID given to link entity with credential id.",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "username": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "matchrule": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Description of the rule group",
                        },
                        "rules": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Operation type for VM rules/filters [CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, ENDS_WITH, EQUALS, STARTS_WITH]",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "name of the VM to be added as content",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME, VM_BADGE, REGION]",
                                    },
                                    "value": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "value for the few type of VM Content like powerstate",
                                    },
                                },
                            },
                        },
                        "existingcredential": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentialid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "credentialname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "cloudaccountsoperationtype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Specifies the operation to be performed on solution list [OVERWRITE, ADD, MODIFY, DELETE]",
            },
            "protectiongroup": {
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
            "filtercontentoperationtype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Specifies the operation to be performed on filter ruleGroups list [OVERWRITE, ADD, MODIFY, DELETE]",
            },
        },
    }
}

func resourceCreateAWSProtectionGroup(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/ProtectionGroup/AWS
    var response_id = strconv.Itoa(0)
    var t_cloudconnection *handler.MsgIdName
    if val, ok := d.GetOk("cloudconnection"); ok {
        t_cloudconnection = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
    }
    var t_workloads []handler.MsgProtectionGroupWorkloadSet
    if val, ok := d.GetOk("workloads"); ok {
        t_workloads = build_awsprotectiongroup_msgprotectiongroupworkloadset_array(d, val.(*schema.Set).List())
    }
    var t_allcloudaccounts *bool
    if val, ok := d.GetOk("allcloudaccounts"); ok {
        t_allcloudaccounts = handler.ToBooleanValue(val, false)
    }
    var t_type *string
    if val, ok := d.GetOk("type"); ok {
        t_type = handler.ToStringValue(val, false)
    }
    var t_content []handler.MsgRuleGroupContentSet
    if val, ok := d.GetOk("content"); ok {
        t_content = build_awsprotectiongroup_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
    }
    var t_contentfilter []handler.MsgRuleGroupContentSet
    if val, ok := d.GetOk("contentfilter"); ok {
        t_contentfilter = build_awsprotectiongroup_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
    }
    var t_cloudaccounts []handler.MsgCloudAccountSet
    if val, ok := d.GetOk("cloudaccounts"); ok {
        t_cloudaccounts = build_awsprotectiongroup_msgcloudaccountset_array(d, val.(*schema.Set).List())
    }
    var t_awsaccountdiscoverydetails *handler.MsgAWSOrganizationAccountDiscoveryDetails
    if val, ok := d.GetOk("awsaccountdiscoverydetails"); ok {
        t_awsaccountdiscoverydetails = build_awsprotectiongroup_msgawsorganizationaccountdiscoverydetails(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_deleteactions *handler.MsgWorkloadProtectionGroupDeleteSettings
    if val, ok := d.GetOk("deleteactions"); ok {
        t_deleteactions = build_awsprotectiongroup_msgworkloadprotectiongroupdeletesettings(d, val.([]interface{}))
    }
    var t_filterworkloads []handler.MsgProtectionGroupWorkloadSet
    if val, ok := d.GetOk("filterworkloads"); ok {
        t_filterworkloads = build_awsprotectiongroup_msgprotectiongroupworkloadset_array(d, val.(*schema.Set).List())
    }
    var t_plan *handler.MsgIdName
    if val, ok := d.GetOk("plan"); ok {
        t_plan = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
    }
    var t_backuponpostcreation *bool
    if val, ok := d.GetOk("backuponpostcreation"); ok {
        t_backuponpostcreation = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgCreateWorkloadProtectionGroupAWSRequest{CloudConnection:t_cloudconnection, Workloads:t_workloads, AllCloudAccounts:t_allcloudaccounts, Type:t_type, Content:t_content, ContentFilter:t_contentfilter, CloudAccounts:t_cloudaccounts, AwsAccountDiscoveryDetails:t_awsaccountdiscoverydetails, Name:t_name, DeleteActions:t_deleteactions, FilterWorkloads:t_filterworkloads, Plan:t_plan, BackupOnPostCreation:t_backuponpostcreation}
    resp, err := handler.CvCreateWorkloadProtectionGroupAWS(req)
    if err != nil {
        return fmt.Errorf("operation [CreateWorkloadProtectionGroupAWS] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateWorkloadProtectionGroupAWS] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateAWSProtectionGroup(d, m)
    }
}

func resourceReadAWSProtectionGroup(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/ProtectionGroup/{protectionGroupId}
    _, err := handler.CvgetProtectionGroupProperties(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [getProtectionGroupProperties] failed, Error %s", err)
    }
    return nil
}

func resourceUpdateAWSProtectionGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ProtectionGroup/{protectionGroupId}
    var t_contentoperationtype *string
    if d.HasChange("contentoperationtype") {
        val := d.Get("contentoperationtype")
        t_contentoperationtype = handler.ToStringValue(val, false)
    }
    var t_filtercontent []handler.MsgRuleGroupContentSet
    if d.HasChange("filtercontent") {
        val := d.Get("filtercontent")
        t_filtercontent = build_awsprotectiongroup_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
    }
    var t_workloads []handler.MsgProtectionGroupWorkloadSet
    if d.HasChange("workloads") {
        val := d.Get("workloads")
        t_workloads = build_awsprotectiongroup_msgprotectiongroupworkloadset_array(d, val.(*schema.Set).List())
    }
    var t_allcloudaccounts *bool
    if d.HasChange("allcloudaccounts") {
        val := d.Get("allcloudaccounts")
        t_allcloudaccounts = handler.ToBooleanValue(val, false)
    }
    var t_cloudaccountsoperationtype *string
    if d.HasChange("cloudaccountsoperationtype") {
        val := d.Get("cloudaccountsoperationtype")
        t_cloudaccountsoperationtype = handler.ToStringValue(val, false)
    }
    var t_type *string
    if d.HasChange("type") {
        val := d.Get("type")
        t_type = handler.ToStringValue(val, false)
    }
    var t_content []handler.MsgRuleGroupContentSet
    if d.HasChange("content") {
        val := d.Get("content")
        t_content = build_awsprotectiongroup_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
    }
    var t_protectiongroup *handler.MsgIdName
    if d.HasChange("protectiongroup") {
        val := d.Get("protectiongroup")
        t_protectiongroup = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
    }
    var t_cloudaccounts []handler.MsgCloudAccountSet
    if d.HasChange("cloudaccounts") {
        val := d.Get("cloudaccounts")
        t_cloudaccounts = build_awsprotectiongroup_msgcloudaccountset_array(d, val.(*schema.Set).List())
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_awsaccountdiscoverydetails *handler.MsgAWSOrganizationAccountDiscoveryDetails
    if d.HasChange("awsaccountdiscoverydetails") {
        val := d.Get("awsaccountdiscoverydetails")
        t_awsaccountdiscoverydetails = build_awsprotectiongroup_msgawsorganizationaccountdiscoverydetails(d, val.([]interface{}))
    }
    var t_deleteactions *handler.MsgWorkloadProtectionGroupDeleteSettings
    if d.HasChange("deleteactions") {
        val := d.Get("deleteactions")
        t_deleteactions = build_awsprotectiongroup_msgworkloadprotectiongroupdeletesettings(d, val.([]interface{}))
    }
    var t_filtercontentoperationtype *string
    if d.HasChange("filtercontentoperationtype") {
        val := d.Get("filtercontentoperationtype")
        t_filtercontentoperationtype = handler.ToStringValue(val, false)
    }
    var t_filterworkloads []handler.MsgProtectionGroupWorkloadSet
    if d.HasChange("filterworkloads") {
        val := d.Get("filterworkloads")
        t_filterworkloads = build_awsprotectiongroup_msgprotectiongroupworkloadset_array(d, val.(*schema.Set).List())
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan")
        t_plan = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
    }
    var req = handler.MsgUpdateWorkloadProtectionGroupRequest{ContentOperationType:t_contentoperationtype, FilterContent:t_filtercontent, Workloads:t_workloads, AllCloudAccounts:t_allcloudaccounts, CloudAccountsOperationType:t_cloudaccountsoperationtype, Type:t_type, Content:t_content, ProtectionGroup:t_protectiongroup, CloudAccounts:t_cloudaccounts, NewName:t_newname, AwsAccountDiscoveryDetails:t_awsaccountdiscoverydetails, DeleteActions:t_deleteactions, FilterContentOperationType:t_filtercontentoperationtype, FilterWorkloads:t_filterworkloads, Plan:t_plan}
    _, err := handler.CvUpdateWorkloadProtectionGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateWorkloadProtectionGroup] failed, Error %s", err)
    }
    return resourceReadAWSProtectionGroup(d, m)
}

func resourceCreateUpdateAWSProtectionGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ProtectionGroup/{protectionGroupId}
    var execUpdate bool = false
    var t_contentoperationtype *string
    if val, ok := d.GetOk("contentoperationtype"); ok {
        t_contentoperationtype = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_filtercontent []handler.MsgRuleGroupContentSet
    if val, ok := d.GetOk("filtercontent"); ok {
        t_filtercontent = build_awsprotectiongroup_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_cloudaccountsoperationtype *string
    if val, ok := d.GetOk("cloudaccountsoperationtype"); ok {
        t_cloudaccountsoperationtype = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_protectiongroup *handler.MsgIdName
    if val, ok := d.GetOk("protectiongroup"); ok {
        t_protectiongroup = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_filtercontentoperationtype *string
    if val, ok := d.GetOk("filtercontentoperationtype"); ok {
        t_filtercontentoperationtype = handler.ToStringValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateWorkloadProtectionGroupRequest{ContentOperationType:t_contentoperationtype, FilterContent:t_filtercontent, CloudAccountsOperationType:t_cloudaccountsoperationtype, ProtectionGroup:t_protectiongroup, FilterContentOperationType:t_filtercontentoperationtype}
        _, err := handler.CvUpdateWorkloadProtectionGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateWorkloadProtectionGroup] failed, Error %s", err)
        }
    }
    return resourceReadAWSProtectionGroup(d, m)
}

func resourceDeleteAWSProtectionGroup(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/ProtectionGroup/{protectionGroupId}
    _, err := handler.CvDeleteWorkloadProtectionGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteWorkloadProtectionGroup] failed, Error %s", err)
    }
    return nil
}

func build_awsprotectiongroup_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_awsprotectiongroup_msgrulegroupcontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleGroupContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleGroupContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_guestcredentialassocid *int
            if val, ok := raw_a["guestcredentialassocid"]; ok {
                t_guestcredentialassocid = handler.ToIntValue(val, true)
            }
            var t_guestcredentials *handler.MsgVMGuestCredentials
            if val, ok := raw_a["guestcredentials"]; ok {
                t_guestcredentials = build_awsprotectiongroup_msgvmguestcredentials(d, val.([]interface{}))
            }
            var t_matchrule *string
            if val, ok := raw_a["matchrule"]; ok {
                t_matchrule = handler.ToStringValue(val, true)
            }
            var t_description *string
            if val, ok := raw_a["description"]; ok {
                t_description = handler.ToStringValue(val, true)
            }
            var t_rules []handler.MsgRuleContentSet
            if val, ok := raw_a["rules"]; ok {
                t_rules = build_awsprotectiongroup_msgrulecontentset_array(d, val.(*schema.Set).List())
            }
            var t_existingcredential *handler.MsgVMExistingCredential
            if val, ok := raw_a["existingcredential"]; ok {
                t_existingcredential = build_awsprotectiongroup_msgvmexistingcredential(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgRuleGroupContentSet{GuestCredentialAssocId:t_guestcredentialassocid, GuestCredentials:t_guestcredentials, MatchRule:t_matchrule, Description:t_description, Rules:t_rules, ExistingCredential:t_existingcredential}
        }
        return tmp
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgvmexistingcredential(d *schema.ResourceData, r []interface{}) *handler.MsgVMExistingCredential {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentialid *int
        if val, ok := tmp["credentialid"]; ok {
            t_credentialid = handler.ToIntValue(val, true)
        }
        var t_credentialname *string
        if val, ok := tmp["credentialname"]; ok {
            t_credentialname = handler.ToStringValue(val, true)
        }
        return &handler.MsgVMExistingCredential{CredentialId:t_credentialid, CredentialName:t_credentialname}
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgrulecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_condition *string
            if val, ok := raw_a["condition"]; ok {
                t_condition = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            var t_value *string
            if val, ok := raw_a["value"]; ok {
                t_value = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgRuleContentSet{Condition:t_condition, Name:t_name, Type:t_type, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgvmguestcredentials(d *schema.ResourceData, r []interface{}) *handler.MsgVMGuestCredentials {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"]; ok {
            t_password = handler.ToStringValue(val, true)
        }
        var t_username *string
        if val, ok := tmp["username"]; ok {
            t_username = handler.ToStringValue(val, true)
        }
        return &handler.MsgVMGuestCredentials{Password:t_password, UserName:t_username}
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgprotectiongroupworkloadset_array(d *schema.ResourceData, r []interface{}) []handler.MsgProtectionGroupWorkloadSet {
    if r != nil {
        tmp := make([]handler.MsgProtectionGroupWorkloadSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_workloadoptions *handler.MsgWorkloadOptions
            if val, ok := raw_a["workloadoptions"]; ok {
                t_workloadoptions = build_awsprotectiongroup_msgworkloadoptions(d, val.([]interface{}))
            }
            var t_solution *handler.MsgIdName
            if val, ok := raw_a["solution"]; ok {
                t_solution = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
            }
            var t_workload *handler.MsgIdName
            if val, ok := raw_a["workload"]; ok {
                t_workload = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgProtectionGroupWorkloadSet{WorkloadOptions:t_workloadoptions, Solution:t_solution, Workload:t_workload}
        }
        return tmp
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgworkloadoptions(d *schema.ResourceData, r []interface{}) *handler.MsgWorkloadOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_islongtermretention *bool
        if val, ok := tmp["islongtermretention"]; ok {
            t_islongtermretention = handler.ToBooleanValue(val, true)
        }
        var t_credential *handler.MsgIdName
        if val, ok := tmp["credential"]; ok {
            t_credential = build_awsprotectiongroup_msgidname(d, val.([]interface{}))
        }
        var t_israpidrecovery *bool
        if val, ok := tmp["israpidrecovery"]; ok {
            t_israpidrecovery = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgWorkloadOptions{IsLongTermRetention:t_islongtermretention, Credential:t_credential, IsRapidRecovery:t_israpidrecovery}
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgworkloadprotectiongroupdeletesettings(d *schema.ResourceData, r []interface{}) *handler.MsgWorkloadProtectionGroupDeleteSettings {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_deconfigureresources *bool
        if val, ok := tmp["deconfigureresources"]; ok {
            t_deconfigureresources = handler.ToBooleanValue(val, true)
        }
        var t_decoupleresources *bool
        if val, ok := tmp["decoupleresources"]; ok {
            t_decoupleresources = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgWorkloadProtectionGroupDeleteSettings{DeconfigureResources:t_deconfigureresources, DecoupleResources:t_decoupleresources}
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgawsorganizationaccountdiscoverydetails(d *schema.ResourceData, r []interface{}) *handler.MsgAWSOrganizationAccountDiscoveryDetails {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_rolenamewithpath *string
        if val, ok := tmp["rolenamewithpath"]; ok {
            t_rolenamewithpath = handler.ToStringValue(val, true)
        }
        var t_accountaccessnodes []handler.MsgAccessNodeItemSet
        if val, ok := tmp["accountaccessnodes"]; ok {
            t_accountaccessnodes = build_awsprotectiongroup_msgaccessnodeitemset_array(d, val.(*schema.Set).List())
        }
        var t_backupserviceaccountid *string
        if val, ok := tmp["backupserviceaccountid"]; ok {
            t_backupserviceaccountid = handler.ToStringValue(val, true)
        }
        return &handler.MsgAWSOrganizationAccountDiscoveryDetails{RoleNameWithPath:t_rolenamewithpath, AccountAccessNodes:t_accountaccessnodes, BackupServiceAccountId:t_backupserviceaccountid}
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgaccessnodeitemset_array(d *schema.ResourceData, r []interface{}) []handler.MsgAccessNodeItemSet {
    if r != nil {
        tmp := make([]handler.MsgAccessNodeItemSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_displayname *string
            if val, ok := raw_a["displayname"]; ok {
                t_displayname = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            var t_accessnodetype *string
            if val, ok := raw_a["accessnodetype"]; ok {
                t_accessnodetype = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgAccessNodeItemSet{DisplayName:t_displayname, Name:t_name, Id:t_id, AccessNodeType:t_accessnodetype}
        }
        return tmp
    } else {
        return nil
    }
}

func build_awsprotectiongroup_msgcloudaccountset_array(d *schema.ResourceData, r []interface{}) []handler.MsgCloudAccountSet {
    if r != nil {
        tmp := make([]handler.MsgCloudAccountSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_id *string
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToStringValue(val, true)
            }
            var t_uuid *string
            if val, ok := raw_a["uuid"]; ok {
                t_uuid = handler.ToStringValue(val, true)
            }
            var t_email *string
            if val, ok := raw_a["email"]; ok {
                t_email = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgCloudAccountSet{Name:t_name, Id:t_id, Uuid:t_uuid, Email:t_email}
        }
        return tmp
    } else {
        return nil
    }
}
