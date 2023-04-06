package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVMGroup_V2() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateVMGroup_V2,
        Read:   resourceReadVMGroup_V2,
        Update: resourceUpdateVMGroup_V2,
        Delete: resourceDeleteVMGroup_V2,

        Schema: map[string]*schema.Schema{
            "meditech": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "systemname": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech system name",
                        },
                        "listenerip": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener IP of FQDN name",
                        },
                        "useraccount": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Sensitive:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "listenerport": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener Port",
                        },
                        "mbftimeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "MBF timeout (in seconds)",
                        },
                    },
                },
            },
            "hypervisor": {
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
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "subclient name ",
            },
            "storage": {
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
            "plan": {
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
            "content": {
                Type:        schema.TypeList,
                Required:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rulegroups": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "matchrule": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
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
                                                    Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME]",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "overwrite": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "activitycontrol": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restoreactivitycontroloptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "delaytime": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "timezone": {
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
                                                "time": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "delay time in unix timestamp",
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "actual delay time value in string format according to the timezone",
                                                },
                                            },
                                        },
                                    },
                                    "activitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]",
                                    },
                                    "enableafteradelay": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "enableactivitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity type is enabled",
                                    },
                                },
                            },
                        },
                        "backupactivitycontroloptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "delaytime": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "timezone": {
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
                                                "time": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "delay time in unix timestamp",
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "actual delay time value in string format according to the timezone",
                                                },
                                            },
                                        },
                                    },
                                    "activitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]",
                                    },
                                    "enableafteradelay": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "enableactivitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity type is enabled",
                                    },
                                },
                            },
                        },
                        "enablebackup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Backup is enabled",
                        },
                        "enablerestore": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Restore is enabled",
                        },
                    },
                },
            },
            "settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "autodetectvmowner": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if auto detect VM Owner enabled",
                        },
                        "collectfiledetailsforgranularrecovery": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if metadata collection is enabled. Only applicable for Indexing v1",
                        },
                        "noofreaders": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of readers for backup",
                        },
                        "usechangedblocktrackingonvm": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if Changed Block Tracking is enabled",
                        },
                        "jobstarttime": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Start Time for the VM Group Job",
                        },
                        "usevmcheckpointsetting": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if use VM CheckPoint setting is enabled",
                        },
                        "customsnapshotresourcegroup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Custom snapshot resource group name for Azure",
                        },
                        "regionalsnapshot": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True when snapshot storage location is regional",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentials": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Sensitive:    true,
                                                    Description: "password to access the network path",
                                                },
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "username to access the network path",
                                                },
                                            },
                                        },
                                    },
                                    "savedcredentials": {
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
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "vmbackuptype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "[APPLICATION_AWARE, FILE_SYSTEM_AND_APPLICATION_CONSISTENT, CRASH_CONSISTENT, APP_BASED_BACKUP, INHERITED]",
                        },
                        "isvmgroupdiskfiltersincluded": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Is VM group disk filters included in VM instance disk filters",
                        },
                        "datastorefreespacecheck": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if Datastore Free space check is enabled",
                        },
                        "allowemptysubclient": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if empty subclient is allowed",
                        },
                        "datastorefreespacerequired": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "precentage of datastore free space check value",
                        },
                        "customsnapshottags": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "represents custom tags to be set on snapshots",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "represents name of the tag",
                                    },
                                    "value": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "represents value of the tag",
                                    },
                                },
                            },
                        },
                        "isapplicationaware": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Is the VM App Aware",
                        },
                        "transportmode": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "transport mode based on environment. Values are case sensitive [AUTO, SAN, HOT_ADD, NAS, NBD_SSL, NBD]",
                        },
                        "collectfiledetailsfromsnapshotcopy": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if metadata collection is enabled for intellisnap jobs. Only applicable for Indexing v1",
                        },
                        "crossaccount": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "shareonly": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if replicate and copy or sharing of amazon snapshot to different amazon account in same or different geographic location is enabled",
                                    },
                                    "fullcopy": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if full copy of amazon snapshot to different amazon account is enabled",
                                    },
                                    "destinationaccount": {
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
                            },
                        },
                    },
                },
            },
            "diskfilters": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rules": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Operation type for VM rules/filters [CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, ENDS_WITH, EQUALS, STARTS_WITH]",
                                    },
                                    "vmname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "VM Name of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "The string to be filtered",
                                    },
                                    "filtertype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "[NONE, DISK_PATH, DISK_PATTERN, DISK_VIRTUAL_DEVICE_NODE, DISK_DATASTORE, DISK_LABEL, DISK_TYPE, DISK_ADDRESS, CONTAINER_PATTERN, DISK_TAG]",
                                    },
                                    "overwrite": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if disk filter in vmgroup has to be overwritten, by default it will append the content",
                                    },
                                    "value": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "The value string to be filtered, in case of disk tag , value of tag to be filtered",
                                    },
                                    "vmguid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "VM Guid of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered",
                                    },
                                },
                            },
                        },
                        "overwrite": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "securityassociations": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "role": {
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
                        "iscreatorassociation": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                        "externalusergroup": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "External User Group Entity",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "providerid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "Provider id",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "External Group Name",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "User Group Id",
                                    },
                                    "providername": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Provider Name",
                                    },
                                },
                            },
                        },
                        "permissionlist": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "permissionid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "exclude": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Flag to specify if this is included permission or excluded permission.",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Returns the type of association. [ALL_CATEGORIES, CATEGORY_ENTITY, PERMISSION_ENTITY]",
                                    },
                                    "categoryname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "categoryid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "permissionname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "user": {
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
                        "usergroup": {
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
            "timezone": {
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
            "filters": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rulegroups": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "matchrule": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
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
                                                    Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME]",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "overwrite": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "accessnode": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "",
                        },
                    },
                },
            },
            "snapshotmanagement": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "useseparateproxyforsnaptotape": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if separate proxy client is used for snap to tape",
                        },
                        "snapengine": {
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
                        "isindependentdisksenabled": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if independent disk option is enabled",
                        },
                        "backupcopyinterface": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "[FILE_SYSTEM, RMAN, VOLUME_COPY]",
                        },
                        "enablehardwaresnapshot": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if hardware snapshot is enabled",
                        },
                        "snapmountproxy": {
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
                        "vmapplicationusername": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Virtual machine application user name",
                        },
                        "snapmountesxhost": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Name of ESX Host",
                        },
                        "israwdevicemapsenabled": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if raw device maps option is enabled",
                        },
                    },
                },
            },
            "enablefileindexing": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "True if file indexing needs to be enabled",
            },
            "applicationvalidation": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "recoverytarget": {
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
                        "schedule": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Schedule for application validation for VM Group",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isscheduleenabled": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if application validation schedule is enabled",
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Description for validation schedule",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "taskid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Job Id for the application validation task. 0 if schedule is disabled",
                                    },
                                },
                            },
                        },
                        "maximumnoofthreads": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of backup Validation Threads",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentials": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Sensitive:    true,
                                                    Description: "password to access the network path",
                                                },
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "username to access the network path",
                                                },
                                            },
                                        },
                                    },
                                    "savedcredentials": {
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
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "keepvalidatedvmsrunning": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "If true then validated VMs will be available until expiration time set on the recovery target",
                        },
                        "validatevmbackups": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if VM Backup validation is enabled",
                        },
                        "usesourcevmesxtomount": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Use Source VM ESX To Mount",
                        },
                        "customvalidationscript": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Custom validation script to be used during VM backup validation",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "windows": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "path": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Path for the validation script",
                                                },
                                                "unccredentials": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "password": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Sensitive:    true,
                                                                Description: "password to access the network path",
                                                            },
                                                            "name": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "username to access the network path",
                                                            },
                                                        },
                                                    },
                                                },
                                                "uncsavedcredentials": {
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
                                                "arguments": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Arguments for the script",
                                                },
                                                "isdisabled": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Is the script disabled",
                                                },
                                                "islocal": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "True if the script is local",
                                                },
                                            },
                                        },
                                    },
                                    "unix": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "path": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Path for the validation script",
                                                },
                                                "unccredentials": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "password": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Sensitive:    true,
                                                                Description: "password to access the network path",
                                                            },
                                                            "name": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "username to access the network path",
                                                            },
                                                        },
                                                    },
                                                },
                                                "uncsavedcredentials": {
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
                                                "arguments": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Arguments for the script",
                                                },
                                                "isdisabled": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Is the script disabled",
                                                },
                                                "islocal": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "True if the script is local",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "copy": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ismirrorcopy": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a mirror copy?",
                                    },
                                    "snapcopytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "[DEFAULT_MIRROR, DEFAULT_VAULT_REPLICA, MIRROR, VAULT_REPLICA, SNAPSHOT_PRIMARY]",
                                    },
                                    "isdefault": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a default backup destination?",
                                    },
                                    "copyprecedence": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Order of backup destinaion copy created in storage policy",
                                    },
                                    "issnapcopy": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a snap copy?",
                                    },
                                    "copytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "[SYNCHRONOUS, SELECTIVE]",
                                    },
                                    "defaultreplicacopy": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a default replica copy?",
                                    },
                                    "isactive": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this an active backup destination?",
                                    },
                                    "arrayreplicacopy": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this an array replica copy?",
                                    },
                                    "backupdestination": {
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
                            },
                        },
                    },
                },
            },
            "meditechsystems": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "systemname": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech system name",
                        },
                        "listenerip": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener IP of FQDN name",
                        },
                        "useraccount": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Sensitive:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "listenerport": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener Port",
                        },
                        "mbftimeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "MBF timeout (in seconds)",
                        },
                    },
                },
            },
        },
    }
}

func resourceCreateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/VMGroup
    var response_id = strconv.Itoa(0)
    var t_meditech *handler.MsgmeditechPropResp
    if val, ok := d.GetOk("meditech"); ok {
        t_meditech = build_vmgroup_v2_msgmeditechpropresp(d, val.([]interface{}))
    }
    var t_hypervisor *handler.MsgIdName
    if val, ok := d.GetOk("hypervisor"); ok {
        t_hypervisor = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_storage *handler.MsgIdName
    if val, ok := d.GetOk("storage"); ok {
        t_storage = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_plan *handler.MsgIdName
    if val, ok := d.GetOk("plan"); ok {
        t_plan = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_content *handler.MsgvmContent
    if val, ok := d.GetOk("content"); ok {
        t_content = build_vmgroup_v2_msgvmcontent(d, val.([]interface{}))
    }
    var req = handler.MsgCreateVMGroupRequest{Meditech:t_meditech, Hypervisor:t_hypervisor, Name:t_name, Storage:t_storage, Plan:t_plan, Content:t_content}
    resp, err := handler.CvCreateVMGroup(req)
    if err != nil {
        return fmt.Errorf("operation [CreateVMGroup] failed, Error %s", err)
    }
    if resp.SubclientId != nil {
        response_id = strconv.Itoa(*resp.SubclientId)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateVMGroup] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateVMGroup_V2(d, m)
    }
}

func resourceReadVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/VmGroup/{VmGroupId}
    resp, err := handler.CvGetVMGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetVMGroup] failed, Error %s", err)
    }
    if rtn, ok := serialize_vmgroup_v2_msgactivitycontroloptions(d, resp.ActivityControl); ok {
        d.Set("activitycontrol", rtn)
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgvmgroupsettings(d, resp.Settings); ok {
        d.Set("settings", rtn)
    } else {
        d.Set("settings", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgvmdiskfilterpropset_array(d, resp.DiskFilters); ok {
        d.Set("diskfilters", rtn)
    } else {
        d.Set("diskfilters", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgsecurityassocset_array(d, resp.SecurityAssociations); ok {
        d.Set("securityassociations", rtn)
    } else {
        d.Set("securityassociations", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgvmcontentset_array(d, resp.Filters); ok {
        d.Set("filters", rtn)
    } else {
        d.Set("filters", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgvmcontentset_array(d, resp.Content); ok {
        d.Set("content", rtn)
    } else {
        d.Set("content", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgsnapcopyinfo(d, resp.SnapshotManagement); ok {
        d.Set("snapshotmanagement", rtn)
    } else {
        d.Set("snapshotmanagement", make([]map[string]interface{}, 0))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_vmgroup_v2_msgvmappvalidation(d, resp.ApplicationValidation); ok {
        d.Set("applicationvalidation", rtn)
    } else {
        d.Set("applicationvalidation", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgmeditechpropresp(d, resp.MeditechSystems); ok {
        d.Set("meditechsystems", rtn)
    } else {
        d.Set("meditechsystems", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, resp.CommonProperties.Hypervisor); ok {
        d.Set("hypervisor", rtn)
    } else {
        d.Set("hypervisor", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, resp.Summary.Plan); ok {
        d.Set("plan", rtn)
    } else {
        d.Set("plan", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/VmGroup/{VmGroupId}
    var t_activitycontrol *handler.MsgActivityControlOptions
    if d.HasChange("activitycontrol") {
        val := d.Get("activitycontrol")
        t_activitycontrol = build_vmgroup_v2_msgactivitycontroloptions(d, val.([]interface{}))
    }
    var t_settings *handler.MsgvmGroupSettings
    if d.HasChange("settings") {
        val := d.Get("settings")
        t_settings = build_vmgroup_v2_msgvmgroupsettings(d, val.([]interface{}))
    }
    var t_diskfilters *handler.MsgvmDiskFilterProp
    if d.HasChange("diskfilters") {
        val := d.Get("diskfilters")
        t_diskfilters = build_vmgroup_v2_msgvmdiskfilterprop(d, val.([]interface{}))
    }
    var t_securityassociations []handler.MsgSecurityAssocSet
    if d.HasChange("securityassociations") {
        val := d.Get("securityassociations")
        t_securityassociations = build_vmgroup_v2_msgsecurityassocset_array(d, val.(*schema.Set).List())
    }
    var t_timezone *handler.MsgIdName
    if d.HasChange("timezone") {
        val := d.Get("timezone")
        t_timezone = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_storage *handler.MsgIdName
    if d.HasChange("storage") {
        val := d.Get("storage")
        t_storage = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_filters *handler.MsgvmContent
    if d.HasChange("filters") {
        val := d.Get("filters")
        t_filters = build_vmgroup_v2_msgvmcontent(d, val.([]interface{}))
    }
    var t_accessnode []handler.MsgIdNameSet
    if d.HasChange("accessnode") {
        val := d.Get("accessnode")
        t_accessnode = build_vmgroup_v2_msgidnameset_array(d, val.(*schema.Set).List())
    }
    var t_content *handler.MsgvmContent
    if d.HasChange("content") {
        val := d.Get("content")
        t_content = build_vmgroup_v2_msgvmcontent(d, val.([]interface{}))
    }
    var t_snapshotmanagement *handler.MsgsnapCopyInfo
    if d.HasChange("snapshotmanagement") {
        val := d.Get("snapshotmanagement")
        t_snapshotmanagement = build_vmgroup_v2_msgsnapcopyinfo(d, val.([]interface{}))
    }
    var t_enablefileindexing *bool
    if d.HasChange("enablefileindexing") {
        val := d.Get("enablefileindexing")
        t_enablefileindexing = handler.ToBooleanValue(val, false)
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_applicationvalidation *handler.MsgvmAppValidation
    if d.HasChange("applicationvalidation") {
        val := d.Get("applicationvalidation")
        t_applicationvalidation = build_vmgroup_v2_msgvmappvalidation(d, val.([]interface{}))
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan")
        t_plan = build_vmgroup_v2_msgidname(d, val.([]interface{}))
    }
    var t_meditechsystems *handler.MsgmeditechPropResp
    if d.HasChange("meditechsystems") {
        val := d.Get("meditechsystems")
        t_meditechsystems = build_vmgroup_v2_msgmeditechpropresp(d, val.([]interface{}))
    }
    var req = handler.MsgUpdateVMGroupRequest{ActivityControl:t_activitycontrol, Settings:t_settings, DiskFilters:t_diskfilters, SecurityAssociations:t_securityassociations, TimeZone:t_timezone, Storage:t_storage, Filters:t_filters, AccessNode:t_accessnode, Content:t_content, SnapshotManagement:t_snapshotmanagement, EnableFileIndexing:t_enablefileindexing, NewName:t_newname, ApplicationValidation:t_applicationvalidation, Plan:t_plan, MeditechSystems:t_meditechsystems}
    _, err := handler.CvUpdateVMGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateVMGroup] failed, Error %s", err)
    }
    return resourceReadVMGroup_V2(d, m)
}

func resourceCreateUpdateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/VmGroup/{VmGroupId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgActivityControlOptions
    if val, ok := d.GetOk("activitycontrol"); ok {
        t_activitycontrol = build_vmgroup_v2_msgactivitycontroloptions(d, val.([]interface{}))
        execUpdate = true
    }
    var t_settings *handler.MsgvmGroupSettings
    if val, ok := d.GetOk("settings"); ok {
        t_settings = build_vmgroup_v2_msgvmgroupsettings(d, val.([]interface{}))
        execUpdate = true
    }
    var t_diskfilters *handler.MsgvmDiskFilterProp
    if val, ok := d.GetOk("diskfilters"); ok {
        t_diskfilters = build_vmgroup_v2_msgvmdiskfilterprop(d, val.([]interface{}))
        execUpdate = true
    }
    var t_securityassociations []handler.MsgSecurityAssocSet
    if val, ok := d.GetOk("securityassociations"); ok {
        t_securityassociations = build_vmgroup_v2_msgsecurityassocset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_timezone *handler.MsgIdName
    if val, ok := d.GetOk("timezone"); ok {
        t_timezone = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_filters *handler.MsgvmContent
    if val, ok := d.GetOk("filters"); ok {
        t_filters = build_vmgroup_v2_msgvmcontent(d, val.([]interface{}))
        execUpdate = true
    }
    var t_accessnode []handler.MsgIdNameSet
    if val, ok := d.GetOk("accessnode"); ok {
        t_accessnode = build_vmgroup_v2_msgidnameset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_snapshotmanagement *handler.MsgsnapCopyInfo
    if val, ok := d.GetOk("snapshotmanagement"); ok {
        t_snapshotmanagement = build_vmgroup_v2_msgsnapcopyinfo(d, val.([]interface{}))
        execUpdate = true
    }
    var t_enablefileindexing *bool
    if val, ok := d.GetOk("enablefileindexing"); ok {
        t_enablefileindexing = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_applicationvalidation *handler.MsgvmAppValidation
    if val, ok := d.GetOk("applicationvalidation"); ok {
        t_applicationvalidation = build_vmgroup_v2_msgvmappvalidation(d, val.([]interface{}))
        execUpdate = true
    }
    var t_meditechsystems *handler.MsgmeditechPropResp
    if val, ok := d.GetOk("meditechsystems"); ok {
        t_meditechsystems = build_vmgroup_v2_msgmeditechpropresp(d, val.([]interface{}))
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateVMGroupRequest{ActivityControl:t_activitycontrol, Settings:t_settings, DiskFilters:t_diskfilters, SecurityAssociations:t_securityassociations, TimeZone:t_timezone, Filters:t_filters, AccessNode:t_accessnode, SnapshotManagement:t_snapshotmanagement, EnableFileIndexing:t_enablefileindexing, ApplicationValidation:t_applicationvalidation, MeditechSystems:t_meditechsystems}
        _, err := handler.CvUpdateVMGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateVMGroup] failed, Error %s", err)
        }
    }
    return resourceReadVMGroup_V2(d, m)
}

func resourceDeleteVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/VmGroup/{VmGroupId}
    _, err := handler.CvDeleteVMGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteVMGroup] failed, Error %s", err)
    }
    return nil
}

func build_vmgroup_v2_msgmeditechpropresp(d *schema.ResourceData, r []interface{}) *handler.MsgmeditechPropResp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_systemname *string
        if val, ok := tmp["systemname"]; ok {
            t_systemname = handler.ToStringValue(val, true)
        }
        var t_listenerip *string
        if val, ok := tmp["listenerip"]; ok {
            t_listenerip = handler.ToStringValue(val, true)
        }
        var t_useraccount *handler.MsgUserNamePassword
        if val, ok := tmp["useraccount"]; ok {
            t_useraccount = build_vmgroup_v2_msgusernamepassword(d, val.([]interface{}))
        }
        var t_listenerport *int
        if val, ok := tmp["listenerport"]; ok {
            t_listenerport = handler.ToIntValue(val, true)
        }
        var t_mbftimeout *int
        if val, ok := tmp["mbftimeout"]; ok {
            t_mbftimeout = handler.ToIntValue(val, true)
        }
        return &handler.MsgmeditechPropResp{SystemName:t_systemname, ListenerIP:t_listenerip, UserAccount:t_useraccount, ListenerPort:t_listenerport, MBFtimeout:t_mbftimeout}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"]; ok {
            t_password = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        return &handler.MsgUserNamePassword{Password:t_password, Name:t_name}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmappvalidation(d *schema.ResourceData, r []interface{}) *handler.MsgvmAppValidation {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_recoverytarget *handler.MsgIdName
        if val, ok := tmp["recoverytarget"]; ok {
            t_recoverytarget = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        var t_schedule *handler.MsgValidationScheduleObject
        if val, ok := tmp["schedule"]; ok {
            t_schedule = build_vmgroup_v2_msgvalidationscheduleobject(d, val.([]interface{}))
        }
        var t_maximumnoofthreads *int
        if val, ok := tmp["maximumnoofthreads"]; ok {
            t_maximumnoofthreads = handler.ToIntValue(val, true)
        }
        var t_guestcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["guestcredentials"]; ok {
            t_guestcredentials = build_vmgroup_v2_msgguestcredentialinfo(d, val.([]interface{}))
        }
        var t_keepvalidatedvmsrunning *bool
        if val, ok := tmp["keepvalidatedvmsrunning"]; ok {
            t_keepvalidatedvmsrunning = handler.ToBooleanValue(val, true)
        }
        var t_validatevmbackups *bool
        if val, ok := tmp["validatevmbackups"]; ok {
            t_validatevmbackups = handler.ToBooleanValue(val, true)
        }
        var t_usesourcevmesxtomount *bool
        if val, ok := tmp["usesourcevmesxtomount"]; ok {
            t_usesourcevmesxtomount = handler.ToBooleanValue(val, true)
        }
        var t_customvalidationscript *handler.MsgappValidationScript
        if val, ok := tmp["customvalidationscript"]; ok {
            t_customvalidationscript = build_vmgroup_v2_msgappvalidationscript(d, val.([]interface{}))
        }
        var t_copy *handler.MsgPlanSourceCopy
        if val, ok := tmp["copy"]; ok {
            t_copy = build_vmgroup_v2_msgplansourcecopy(d, val.([]interface{}))
        }
        return &handler.MsgvmAppValidation{RecoveryTarget:t_recoverytarget, Schedule:t_schedule, MaximumNoOfThreads:t_maximumnoofthreads, GuestCredentials:t_guestcredentials, KeepValidatedVMsRunning:t_keepvalidatedvmsrunning, ValidateVMBackups:t_validatevmbackups, UseSourceVmESXToMount:t_usesourcevmesxtomount, CustomValidationScript:t_customvalidationscript, Copy:t_copy}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgplansourcecopy(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSourceCopy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_ismirrorcopy *bool
        if val, ok := tmp["ismirrorcopy"]; ok {
            t_ismirrorcopy = handler.ToBooleanValue(val, true)
        }
        var t_snapcopytype *string
        if val, ok := tmp["snapcopytype"]; ok {
            t_snapcopytype = handler.ToStringValue(val, true)
        }
        var t_isdefault *bool
        if val, ok := tmp["isdefault"]; ok {
            t_isdefault = handler.ToBooleanValue(val, true)
        }
        var t_copyprecedence *int
        if val, ok := tmp["copyprecedence"]; ok {
            t_copyprecedence = handler.ToIntValue(val, true)
        }
        var t_issnapcopy *bool
        if val, ok := tmp["issnapcopy"]; ok {
            t_issnapcopy = handler.ToBooleanValue(val, true)
        }
        var t_copytype *string
        if val, ok := tmp["copytype"]; ok {
            t_copytype = handler.ToStringValue(val, true)
        }
        var t_defaultreplicacopy *bool
        if val, ok := tmp["defaultreplicacopy"]; ok {
            t_defaultreplicacopy = handler.ToBooleanValue(val, true)
        }
        var t_isactive *bool
        if val, ok := tmp["isactive"]; ok {
            t_isactive = handler.ToBooleanValue(val, true)
        }
        var t_arrayreplicacopy *bool
        if val, ok := tmp["arrayreplicacopy"]; ok {
            t_arrayreplicacopy = handler.ToBooleanValue(val, true)
        }
        var t_backupdestination *handler.MsgIdName
        if val, ok := tmp["backupdestination"]; ok {
            t_backupdestination = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgPlanSourceCopy{IsMirrorCopy:t_ismirrorcopy, SnapCopyType:t_snapcopytype, IsDefault:t_isdefault, CopyPrecedence:t_copyprecedence, IsSnapCopy:t_issnapcopy, CopyType:t_copytype, DefaultReplicaCopy:t_defaultreplicacopy, IsActive:t_isactive, ArrayReplicaCopy:t_arrayreplicacopy, BackupDestination:t_backupdestination}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_vmgroup_v2_msgappvalidationscript(d *schema.ResourceData, r []interface{}) *handler.MsgappValidationScript {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_windows *handler.MsgValidationScript
        if val, ok := tmp["windows"]; ok {
            t_windows = build_vmgroup_v2_msgvalidationscript(d, val.([]interface{}))
        }
        var t_unix *handler.MsgValidationScript
        if val, ok := tmp["unix"]; ok {
            t_unix = build_vmgroup_v2_msgvalidationscript(d, val.([]interface{}))
        }
        return &handler.MsgappValidationScript{Windows:t_windows, Unix:t_unix}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvalidationscript(d *schema.ResourceData, r []interface{}) *handler.MsgValidationScript {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_path *string
        if val, ok := tmp["path"]; ok {
            t_path = handler.ToStringValue(val, true)
        }
        var t_unccredentials *handler.MsgUserNamePassword
        if val, ok := tmp["unccredentials"]; ok {
            t_unccredentials = build_vmgroup_v2_msgusernamepassword(d, val.([]interface{}))
        }
        var t_uncsavedcredentials *handler.MsgIdName
        if val, ok := tmp["uncsavedcredentials"]; ok {
            t_uncsavedcredentials = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        var t_arguments *string
        if val, ok := tmp["arguments"]; ok {
            t_arguments = handler.ToStringValue(val, true)
        }
        var t_isdisabled *bool
        if val, ok := tmp["isdisabled"]; ok {
            t_isdisabled = handler.ToBooleanValue(val, true)
        }
        var t_islocal *bool
        if val, ok := tmp["islocal"]; ok {
            t_islocal = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgValidationScript{Path:t_path, UNCCredentials:t_unccredentials, UNCSavedCredentials:t_uncsavedcredentials, Arguments:t_arguments, IsDisabled:t_isdisabled, IsLocal:t_islocal}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgguestcredentialinfo(d *schema.ResourceData, r []interface{}) *handler.MsgguestCredentialInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentials *handler.MsgUserNamePassword
        if val, ok := tmp["credentials"]; ok {
            t_credentials = build_vmgroup_v2_msgusernamepassword(d, val.([]interface{}))
        }
        var t_savedcredentials *handler.MsgIdName
        if val, ok := tmp["savedcredentials"]; ok {
            t_savedcredentials = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgguestCredentialInfo{Credentials:t_credentials, SavedCredentials:t_savedcredentials}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvalidationscheduleobject(d *schema.ResourceData, r []interface{}) *handler.MsgValidationScheduleObject {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_isscheduleenabled *bool
        if val, ok := tmp["isscheduleenabled"]; ok {
            t_isscheduleenabled = handler.ToBooleanValue(val, true)
        }
        var t_description *string
        if val, ok := tmp["description"]; ok {
            t_description = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        var t_taskid *int
        if val, ok := tmp["taskid"]; ok {
            t_taskid = handler.ToIntValue(val, true)
        }
        return &handler.MsgValidationScheduleObject{IsScheduleEnabled:t_isscheduleenabled, Description:t_description, Id:t_id, TaskId:t_taskid}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgsnapcopyinfo(d *schema.ResourceData, r []interface{}) *handler.MsgsnapCopyInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_useseparateproxyforsnaptotape *bool
        if val, ok := tmp["useseparateproxyforsnaptotape"]; ok {
            t_useseparateproxyforsnaptotape = handler.ToBooleanValue(val, true)
        }
        var t_snapengine *handler.MsgIdName
        if val, ok := tmp["snapengine"]; ok {
            t_snapengine = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        var t_isindependentdisksenabled *bool
        if val, ok := tmp["isindependentdisksenabled"]; ok {
            t_isindependentdisksenabled = handler.ToBooleanValue(val, true)
        }
        var t_backupcopyinterface *string
        if val, ok := tmp["backupcopyinterface"]; ok {
            t_backupcopyinterface = handler.ToStringValue(val, true)
        }
        var t_enablehardwaresnapshot *bool
        if val, ok := tmp["enablehardwaresnapshot"]; ok {
            t_enablehardwaresnapshot = handler.ToBooleanValue(val, true)
        }
        var t_snapmountproxy *handler.MsgIdName
        if val, ok := tmp["snapmountproxy"]; ok {
            t_snapmountproxy = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        var t_vmapplicationusername *string
        if val, ok := tmp["vmapplicationusername"]; ok {
            t_vmapplicationusername = handler.ToStringValue(val, true)
        }
        var t_snapmountesxhost *string
        if val, ok := tmp["snapmountesxhost"]; ok {
            t_snapmountesxhost = handler.ToStringValue(val, true)
        }
        var t_israwdevicemapsenabled *bool
        if val, ok := tmp["israwdevicemapsenabled"]; ok {
            t_israwdevicemapsenabled = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgsnapCopyInfo{UseSeparateProxyForSnapToTape:t_useseparateproxyforsnaptotape, SnapEngine:t_snapengine, IsIndependentDisksEnabled:t_isindependentdisksenabled, BackupCopyInterface:t_backupcopyinterface, EnableHardwareSnapshot:t_enablehardwaresnapshot, SnapMountProxy:t_snapmountproxy, VmApplicationUserName:t_vmapplicationusername, SnapMountESXHost:t_snapmountesxhost, IsRawDeviceMapsEnabled:t_israwdevicemapsenabled}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
    if r != nil {
        tmp := make([]handler.MsgIdNameSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgIdNameSet{Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmcontent(d *schema.ResourceData, r []interface{}) *handler.MsgvmContent {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_rulegroups []handler.MsgRuleGroupContentSet
        if val, ok := tmp["rulegroups"]; ok {
            t_rulegroups = build_vmgroup_v2_msgrulegroupcontentset_array(d, val.(*schema.Set).List())
        }
        var t_overwrite *bool
        if val, ok := tmp["overwrite"]; ok {
            t_overwrite = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgvmContent{RuleGroups:t_rulegroups, Overwrite:t_overwrite}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgrulegroupcontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleGroupContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleGroupContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_matchrule *string
            if val, ok := raw_a["matchrule"]; ok {
                t_matchrule = handler.ToStringValue(val, true)
            }
            var t_rules []handler.MsgRuleContentSet
            if val, ok := raw_a["rules"]; ok {
                t_rules = build_vmgroup_v2_msgrulecontentset_array(d, val.(*schema.Set).List())
            }
            tmp[a] = handler.MsgRuleGroupContentSet{MatchRule:t_matchrule, Rules:t_rules}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgrulecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleContentSet {
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
            tmp[a] = handler.MsgRuleContentSet{Condition:t_condition, Name:t_name, Type:t_type}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgsecurityassocset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSecurityAssocSet {
    if r != nil {
        tmp := make([]handler.MsgSecurityAssocSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_role *handler.MsgIdName
            if val, ok := raw_a["role"]; ok {
                t_role = build_vmgroup_v2_msgidname(d, val.([]interface{}))
            }
            var t_iscreatorassociation *bool
            if val, ok := raw_a["iscreatorassociation"]; ok {
                t_iscreatorassociation = handler.ToBooleanValue(val, true)
            }
            var t_externalusergroup *handler.MsgexternalUserGroup
            if val, ok := raw_a["externalusergroup"]; ok {
                t_externalusergroup = build_vmgroup_v2_msgexternalusergroup(d, val.([]interface{}))
            }
            var t_permissionlist []handler.MsgPermissionRespSet
            if val, ok := raw_a["permissionlist"]; ok {
                t_permissionlist = build_vmgroup_v2_msgpermissionrespset_array(d, val.(*schema.Set).List())
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"]; ok {
                t_user = build_vmgroup_v2_msgidname(d, val.([]interface{}))
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"]; ok {
                t_usergroup = build_vmgroup_v2_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgSecurityAssocSet{Role:t_role, IsCreatorAssociation:t_iscreatorassociation, ExternalUserGroup:t_externalusergroup, PermissionList:t_permissionlist, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgpermissionrespset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPermissionRespSet {
    if r != nil {
        tmp := make([]handler.MsgPermissionRespSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_permissionid *int
            if val, ok := raw_a["permissionid"]; ok {
                t_permissionid = handler.ToIntValue(val, true)
            }
            var t_exclude *bool
            if val, ok := raw_a["exclude"]; ok {
                t_exclude = handler.ToBooleanValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            var t_categoryname *string
            if val, ok := raw_a["categoryname"]; ok {
                t_categoryname = handler.ToStringValue(val, true)
            }
            var t_categoryid *int
            if val, ok := raw_a["categoryid"]; ok {
                t_categoryid = handler.ToIntValue(val, true)
            }
            var t_permissionname *string
            if val, ok := raw_a["permissionname"]; ok {
                t_permissionname = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgPermissionRespSet{PermissionId:t_permissionid, Exclude:t_exclude, Type:t_type, CategoryName:t_categoryname, CategoryId:t_categoryid, PermissionName:t_permissionname}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgexternalusergroup(d *schema.ResourceData, r []interface{}) *handler.MsgexternalUserGroup {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_providerid *int
        if val, ok := tmp["providerid"]; ok {
            t_providerid = handler.ToIntValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        var t_providername *string
        if val, ok := tmp["providername"]; ok {
            t_providername = handler.ToStringValue(val, true)
        }
        return &handler.MsgexternalUserGroup{ProviderId:t_providerid, Name:t_name, Id:t_id, ProviderName:t_providername}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmdiskfilterprop(d *schema.ResourceData, r []interface{}) *handler.MsgvmDiskFilterProp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_rules []handler.MsgvmDiskFilterSet
        if val, ok := tmp["rules"]; ok {
            t_rules = build_vmgroup_v2_msgvmdiskfilterset_array(d, val.(*schema.Set).List())
        }
        var t_overwrite *bool
        if val, ok := tmp["overwrite"]; ok {
            t_overwrite = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgvmDiskFilterProp{Rules:t_rules, Overwrite:t_overwrite}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmdiskfilterset_array(d *schema.ResourceData, r []interface{}) []handler.MsgvmDiskFilterSet {
    if r != nil {
        tmp := make([]handler.MsgvmDiskFilterSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_condition *string
            if val, ok := raw_a["condition"]; ok {
                t_condition = handler.ToStringValue(val, true)
            }
            var t_vmname *string
            if val, ok := raw_a["vmname"]; ok {
                t_vmname = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_filtertype *string
            if val, ok := raw_a["filtertype"]; ok {
                t_filtertype = handler.ToStringValue(val, true)
            }
            var t_overwrite *bool
            if val, ok := raw_a["overwrite"]; ok {
                t_overwrite = handler.ToBooleanValue(val, true)
            }
            var t_value *string
            if val, ok := raw_a["value"]; ok {
                t_value = handler.ToStringValue(val, true)
            }
            var t_vmguid *string
            if val, ok := raw_a["vmguid"]; ok {
                t_vmguid = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgvmDiskFilterSet{Condition:t_condition, VmName:t_vmname, Name:t_name, FilterType:t_filtertype, Overwrite:t_overwrite, Value:t_value, VmGuid:t_vmguid}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmgroupsettings(d *schema.ResourceData, r []interface{}) *handler.MsgvmGroupSettings {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_autodetectvmowner *bool
        if val, ok := tmp["autodetectvmowner"]; ok {
            t_autodetectvmowner = handler.ToBooleanValue(val, true)
        }
        var t_collectfiledetailsforgranularrecovery *bool
        if val, ok := tmp["collectfiledetailsforgranularrecovery"]; ok {
            t_collectfiledetailsforgranularrecovery = handler.ToBooleanValue(val, true)
        }
        var t_noofreaders *int
        if val, ok := tmp["noofreaders"]; ok {
            t_noofreaders = handler.ToIntValue(val, true)
        }
        var t_usechangedblocktrackingonvm *bool
        if val, ok := tmp["usechangedblocktrackingonvm"]; ok {
            t_usechangedblocktrackingonvm = handler.ToBooleanValue(val, true)
        }
        var t_jobstarttime *int
        if val, ok := tmp["jobstarttime"]; ok {
            t_jobstarttime = handler.ToIntValue(val, true)
        }
        var t_usevmcheckpointsetting *bool
        if val, ok := tmp["usevmcheckpointsetting"]; ok {
            t_usevmcheckpointsetting = handler.ToBooleanValue(val, true)
        }
        var t_customsnapshotresourcegroup *string
        if val, ok := tmp["customsnapshotresourcegroup"]; ok {
            t_customsnapshotresourcegroup = handler.ToStringValue(val, true)
        }
        var t_regionalsnapshot *bool
        if val, ok := tmp["regionalsnapshot"]; ok {
            t_regionalsnapshot = handler.ToBooleanValue(val, true)
        }
        var t_guestcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["guestcredentials"]; ok {
            t_guestcredentials = build_vmgroup_v2_msgguestcredentialinfo(d, val.([]interface{}))
        }
        var t_vmbackuptype *string
        if val, ok := tmp["vmbackuptype"]; ok {
            t_vmbackuptype = handler.ToStringValue(val, true)
        }
        var t_isvmgroupdiskfiltersincluded *bool
        if val, ok := tmp["isvmgroupdiskfiltersincluded"]; ok {
            t_isvmgroupdiskfiltersincluded = handler.ToBooleanValue(val, true)
        }
        var t_datastorefreespacecheck *bool
        if val, ok := tmp["datastorefreespacecheck"]; ok {
            t_datastorefreespacecheck = handler.ToBooleanValue(val, true)
        }
        var t_allowemptysubclient *bool
        if val, ok := tmp["allowemptysubclient"]; ok {
            t_allowemptysubclient = handler.ToBooleanValue(val, true)
        }
        var t_datastorefreespacerequired *int
        if val, ok := tmp["datastorefreespacerequired"]; ok {
            t_datastorefreespacerequired = handler.ToIntValue(val, true)
        }
        var t_customsnapshottags []handler.MsgresourceTagSet
        if val, ok := tmp["customsnapshottags"]; ok {
            t_customsnapshottags = build_vmgroup_v2_msgresourcetagset_array(d, val.(*schema.Set).List())
        }
        var t_isapplicationaware *bool
        if val, ok := tmp["isapplicationaware"]; ok {
            t_isapplicationaware = handler.ToBooleanValue(val, true)
        }
        var t_transportmode *string
        if val, ok := tmp["transportmode"]; ok {
            t_transportmode = handler.ToStringValue(val, true)
        }
        var t_collectfiledetailsfromsnapshotcopy *bool
        if val, ok := tmp["collectfiledetailsfromsnapshotcopy"]; ok {
            t_collectfiledetailsfromsnapshotcopy = handler.ToBooleanValue(val, true)
        }
        var t_crossaccount *handler.MsgAmazonCrossAccount
        if val, ok := tmp["crossaccount"]; ok {
            t_crossaccount = build_vmgroup_v2_msgamazoncrossaccount(d, val.([]interface{}))
        }
        return &handler.MsgvmGroupSettings{AutoDetectVMOwner:t_autodetectvmowner, CollectFileDetailsforGranularRecovery:t_collectfiledetailsforgranularrecovery, NoOfReaders:t_noofreaders, UseChangedBlockTrackingOnVM:t_usechangedblocktrackingonvm, JobStartTime:t_jobstarttime, UseVMCheckpointSetting:t_usevmcheckpointsetting, CustomSnapshotResourceGroup:t_customsnapshotresourcegroup, RegionalSnapshot:t_regionalsnapshot, GuestCredentials:t_guestcredentials, VmBackupType:t_vmbackuptype, IsVMGroupDiskFiltersIncluded:t_isvmgroupdiskfiltersincluded, DatastoreFreespaceCheck:t_datastorefreespacecheck, AllowEmptySubclient:t_allowemptysubclient, DatastoreFreespaceRequired:t_datastorefreespacerequired, CustomSnapshotTags:t_customsnapshottags, IsApplicationAware:t_isapplicationaware, TransportMode:t_transportmode, CollectFileDetailsFromSnapshotCopy:t_collectfiledetailsfromsnapshotcopy, CrossAccount:t_crossaccount}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgamazoncrossaccount(d *schema.ResourceData, r []interface{}) *handler.MsgAmazonCrossAccount {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_shareonly *bool
        if val, ok := tmp["shareonly"]; ok {
            t_shareonly = handler.ToBooleanValue(val, true)
        }
        var t_fullcopy *bool
        if val, ok := tmp["fullcopy"]; ok {
            t_fullcopy = handler.ToBooleanValue(val, true)
        }
        var t_destinationaccount *handler.MsgIdName
        if val, ok := tmp["destinationaccount"]; ok {
            t_destinationaccount = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgAmazonCrossAccount{ShareOnly:t_shareonly, FullCopy:t_fullcopy, DestinationAccount:t_destinationaccount}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgresourcetagset_array(d *schema.ResourceData, r []interface{}) []handler.MsgresourceTagSet {
    if r != nil {
        tmp := make([]handler.MsgresourceTagSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_value *string
            if val, ok := raw_a["value"]; ok {
                t_value = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgresourceTagSet{Name:t_name, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgactivitycontroloptions(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_restoreactivitycontroloptions *handler.MsgActivityControlOptionsProp
        if val, ok := tmp["restoreactivitycontroloptions"]; ok {
            t_restoreactivitycontroloptions = build_vmgroup_v2_msgactivitycontroloptionsprop(d, val.([]interface{}))
        }
        var t_backupactivitycontroloptions *handler.MsgActivityControlOptionsProp
        if val, ok := tmp["backupactivitycontroloptions"]; ok {
            t_backupactivitycontroloptions = build_vmgroup_v2_msgactivitycontroloptionsprop(d, val.([]interface{}))
        }
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"]; ok {
            t_enablebackup = handler.ToBooleanValue(val, true)
        }
        var t_enablerestore *bool
        if val, ok := tmp["enablerestore"]; ok {
            t_enablerestore = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgActivityControlOptions{RestoreActivityControlOptions:t_restoreactivitycontroloptions, BackupActivityControlOptions:t_backupactivitycontroloptions, EnableBackup:t_enablebackup, EnableRestore:t_enablerestore}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgactivitycontroloptionsprop(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptionsProp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_delaytime *handler.MsgActivityControlTileDelayTime
        if val, ok := tmp["delaytime"]; ok {
            t_delaytime = build_vmgroup_v2_msgactivitycontroltiledelaytime(d, val.([]interface{}))
        }
        var t_activitytype *string
        if val, ok := tmp["activitytype"]; ok {
            t_activitytype = handler.ToStringValue(val, true)
        }
        var t_enableafteradelay *bool
        if val, ok := tmp["enableafteradelay"]; ok {
            t_enableafteradelay = handler.ToBooleanValue(val, true)
        }
        var t_enableactivitytype *bool
        if val, ok := tmp["enableactivitytype"]; ok {
            t_enableactivitytype = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgActivityControlOptionsProp{DelayTime:t_delaytime, ActivityType:t_activitytype, EnableAfterADelay:t_enableafteradelay, EnableActivityType:t_enableactivitytype}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgactivitycontroltiledelaytime(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlTileDelayTime {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"]; ok {
            t_timezone = build_vmgroup_v2_msgidname(d, val.([]interface{}))
        }
        var t_time *int
        if val, ok := tmp["time"]; ok {
            t_time = handler.ToIntValue(val, true)
        }
        var t_value *string
        if val, ok := tmp["value"]; ok {
            t_value = handler.ToStringValue(val, true)
        }
        return &handler.MsgActivityControlTileDelayTime{TimeZone:t_timezone, Time:t_time, Value:t_value}
    } else {
        return nil
    }
}

func serialize_vmgroup_v2_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgIdName
    //MsgIdName
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.Id != nil {
        val[0]["id"] = data.Id
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgmeditechpropresp(d *schema.ResourceData, data *handler.MsgmeditechPropResp) ([]map[string]interface{}, bool) {
    //MsgmeditechPropResp
    //MsgmeditechPropResp
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.SystemName != nil {
        val[0]["systemname"] = data.SystemName
        added = true
    }
    if data.ListenerIP != nil {
        val[0]["listenerip"] = data.ListenerIP
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgusernamepassword(d, data.UserAccount); ok {
        val[0]["useraccount"] = rtn
        added = true
    }
    if data.ListenerPort != nil {
        val[0]["listenerport"] = data.ListenerPort
        added = true
    }
    if data.MBFtimeout != nil {
        val[0]["mbftimeout"] = data.MBFtimeout
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgusernamepassword(d *schema.ResourceData, data *handler.MsgUserNamePassword) ([]map[string]interface{}, bool) {
    //MsgmeditechPropResp -> MsgUserNamePassword
    //MsgmeditechPropResp -> MsgUserNamePassword
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Password != nil {
        val[0]["password"] = data.Password
        added = true
    }
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgvmappvalidation(d *schema.ResourceData, data *handler.MsgvmAppValidation) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation
    //MsgvmAppValidation
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.RecoveryTarget); ok {
        val[0]["recoverytarget"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgvalidationscheduleobject(d, data.Schedule); ok {
        val[0]["schedule"] = rtn
        added = true
    }
    if data.MaximumNoOfThreads != nil {
        val[0]["maximumnoofthreads"] = data.MaximumNoOfThreads
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgguestcredentialinfo(d, data.GuestCredentials); ok {
        val[0]["guestcredentials"] = rtn
        added = true
    }
    if data.KeepValidatedVMsRunning != nil {
        val[0]["keepvalidatedvmsrunning"] = strconv.FormatBool(*data.KeepValidatedVMsRunning)
        added = true
    }
    if data.ValidateVMBackups != nil {
        val[0]["validatevmbackups"] = strconv.FormatBool(*data.ValidateVMBackups)
        added = true
    }
    if data.UseSourceVmESXToMount != nil {
        val[0]["usesourcevmesxtomount"] = strconv.FormatBool(*data.UseSourceVmESXToMount)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgappvalidationscript(d, data.CustomValidationScript); ok {
        val[0]["customvalidationscript"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgplansourcecopy(d, data.Copy); ok {
        val[0]["copy"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgplansourcecopy(d *schema.ResourceData, data *handler.MsgPlanSourceCopy) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation -> MsgPlanSourceCopy
    //MsgvmAppValidation -> MsgPlanSourceCopy
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.IsMirrorCopy != nil {
        val[0]["ismirrorcopy"] = strconv.FormatBool(*data.IsMirrorCopy)
        added = true
    }
    if data.SnapCopyType != nil {
        val[0]["snapcopytype"] = data.SnapCopyType
        added = true
    }
    if data.IsDefault != nil {
        val[0]["isdefault"] = strconv.FormatBool(*data.IsDefault)
        added = true
    }
    if data.CopyPrecedence != nil {
        val[0]["copyprecedence"] = data.CopyPrecedence
        added = true
    }
    if data.IsSnapCopy != nil {
        val[0]["issnapcopy"] = strconv.FormatBool(*data.IsSnapCopy)
        added = true
    }
    if data.CopyType != nil {
        val[0]["copytype"] = data.CopyType
        added = true
    }
    if data.DefaultReplicaCopy != nil {
        val[0]["defaultreplicacopy"] = strconv.FormatBool(*data.DefaultReplicaCopy)
        added = true
    }
    if data.IsActive != nil {
        val[0]["isactive"] = strconv.FormatBool(*data.IsActive)
        added = true
    }
    if data.ArrayReplicaCopy != nil {
        val[0]["arrayreplicacopy"] = strconv.FormatBool(*data.ArrayReplicaCopy)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.BackupDestination); ok {
        val[0]["backupdestination"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgappvalidationscript(d *schema.ResourceData, data *handler.MsgappValidationScript) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation -> MsgappValidationScript
    //MsgvmAppValidation -> MsgappValidationScript
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgvalidationscript(d, data.Windows); ok {
        val[0]["windows"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgvalidationscript(d, data.Unix); ok {
        val[0]["unix"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgvalidationscript(d *schema.ResourceData, data *handler.MsgValidationScript) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation -> MsgappValidationScript -> MsgValidationScript
    //MsgvmAppValidation -> MsgappValidationScript -> MsgValidationScript
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Path != nil {
        val[0]["path"] = data.Path
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgusernamepassword(d, data.UNCCredentials); ok {
        val[0]["unccredentials"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.UNCSavedCredentials); ok {
        val[0]["uncsavedcredentials"] = rtn
        added = true
    }
    if data.Arguments != nil {
        val[0]["arguments"] = data.Arguments
        added = true
    }
    if data.IsDisabled != nil {
        val[0]["isdisabled"] = strconv.FormatBool(*data.IsDisabled)
        added = true
    }
    if data.IsLocal != nil {
        val[0]["islocal"] = strconv.FormatBool(*data.IsLocal)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgguestcredentialinfo(d *schema.ResourceData, data *handler.MsgguestCredentialInfo) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation -> MsgguestCredentialInfo
    //MsgvmAppValidation -> MsgguestCredentialInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgusernamepassword(d, data.Credentials); ok {
        val[0]["credentials"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.SavedCredentials); ok {
        val[0]["savedcredentials"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgvalidationscheduleobject(d *schema.ResourceData, data *handler.MsgValidationScheduleObject) ([]map[string]interface{}, bool) {
    //MsgvmAppValidation -> MsgValidationScheduleObject
    //MsgvmAppValidation -> MsgValidationScheduleObject
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.IsScheduleEnabled != nil {
        val[0]["isscheduleenabled"] = strconv.FormatBool(*data.IsScheduleEnabled)
        added = true
    }
    if data.Description != nil {
        val[0]["description"] = data.Description
        added = true
    }
    if data.Id != nil {
        val[0]["id"] = data.Id
        added = true
    }
    if data.TaskId != nil {
        val[0]["taskid"] = data.TaskId
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgsnapcopyinfo(d *schema.ResourceData, data *handler.MsgsnapCopyInfo) ([]map[string]interface{}, bool) {
    //MsgsnapCopyInfo
    //MsgsnapCopyInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.UseSeparateProxyForSnapToTape != nil {
        val[0]["useseparateproxyforsnaptotape"] = strconv.FormatBool(*data.UseSeparateProxyForSnapToTape)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.SnapEngine); ok {
        val[0]["snapengine"] = rtn
        added = true
    }
    if data.IsIndependentDisksEnabled != nil {
        val[0]["isindependentdisksenabled"] = strconv.FormatBool(*data.IsIndependentDisksEnabled)
        added = true
    }
    if data.BackupCopyInterface != nil {
        val[0]["backupcopyinterface"] = data.BackupCopyInterface
        added = true
    }
    if data.EnableHardwareSnapshot != nil {
        val[0]["enablehardwaresnapshot"] = strconv.FormatBool(*data.EnableHardwareSnapshot)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.SnapMountProxy); ok {
        val[0]["snapmountproxy"] = rtn
        added = true
    }
    if data.VmApplicationUserName != nil {
        val[0]["vmapplicationusername"] = data.VmApplicationUserName
        added = true
    }
    if data.SnapMountESXHost != nil {
        val[0]["snapmountesxhost"] = data.SnapMountESXHost
        added = true
    }
    if data.IsRawDeviceMapsEnabled != nil {
        val[0]["israwdevicemapsenabled"] = strconv.FormatBool(*data.IsRawDeviceMapsEnabled)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgvmcontentset_array(d *schema.ResourceData, data []handler.MsgvmContentSet) ([]map[string]interface{}, bool) {
    //MsgvmContent
    //MsgvmContentSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if rtn, ok := serialize_vmgroup_v2_msgrulegroupcontentset_array(d, data[i].RuleGroups); ok {
            tmp["rulegroups"] = rtn
            added = true
        }
        if data[i].Overwrite != nil {
            tmp["overwrite"] = strconv.FormatBool(*data[i].Overwrite)
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgrulegroupcontentset_array(d *schema.ResourceData, data []handler.MsgRuleGroupContentSet) ([]map[string]interface{}, bool) {
    //MsgvmContent -> MsgRuleGroupContentSet
    //MsgvmContentSet -> MsgRuleGroupContentSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].MatchRule != nil {
            tmp["matchrule"] = data[i].MatchRule
            added = true
        }
        if rtn, ok := serialize_vmgroup_v2_msgrulecontentset_array(d, data[i].Rules); ok {
            tmp["rules"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgrulecontentset_array(d *schema.ResourceData, data []handler.MsgRuleContentSet) ([]map[string]interface{}, bool) {
    //MsgvmContent -> MsgRuleGroupContentSet -> MsgRuleContentSet
    //MsgvmContentSet -> MsgRuleGroupContentSet -> MsgRuleContentSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Condition != nil {
            tmp["condition"] = data[i].Condition
            added = true
        }
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
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

func serialize_vmgroup_v2_msgsecurityassocset_array(d *schema.ResourceData, data []handler.MsgSecurityAssocSet) ([]map[string]interface{}, bool) {
    //MsgSecurityAssocSet
    //MsgSecurityAssocSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if rtn, ok := serialize_vmgroup_v2_msgidname(d, data[i].Role); ok {
            tmp["role"] = rtn
            added = true
        }
        if data[i].IsCreatorAssociation != nil {
            tmp["iscreatorassociation"] = strconv.FormatBool(*data[i].IsCreatorAssociation)
            added = true
        }
        if rtn, ok := serialize_vmgroup_v2_msgexternalusergroup(d, data[i].ExternalUserGroup); ok {
            tmp["externalusergroup"] = rtn
            added = true
        }
        if rtn, ok := serialize_vmgroup_v2_msgpermissionrespset_array(d, data[i].PermissionList); ok {
            tmp["permissionlist"] = rtn
            added = true
        }
        if rtn, ok := serialize_vmgroup_v2_msgidname(d, data[i].User); ok {
            tmp["user"] = rtn
            added = true
        }
        if rtn, ok := serialize_vmgroup_v2_msgidname(d, data[i].UserGroup); ok {
            tmp["usergroup"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgpermissionrespset_array(d *schema.ResourceData, data []handler.MsgPermissionRespSet) ([]map[string]interface{}, bool) {
    //MsgSecurityAssocSet -> MsgPermissionRespSet
    //MsgSecurityAssocSet -> MsgPermissionRespSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].PermissionId != nil {
            tmp["permissionid"] = data[i].PermissionId
            added = true
        }
        if data[i].Exclude != nil {
            tmp["exclude"] = strconv.FormatBool(*data[i].Exclude)
            added = true
        }
        if data[i].Type != nil {
            tmp["type"] = data[i].Type
            added = true
        }
        if data[i].CategoryName != nil {
            tmp["categoryname"] = data[i].CategoryName
            added = true
        }
        if data[i].CategoryId != nil {
            tmp["categoryid"] = data[i].CategoryId
            added = true
        }
        if data[i].PermissionName != nil {
            tmp["permissionname"] = data[i].PermissionName
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgexternalusergroup(d *schema.ResourceData, data *handler.MsgexternalUserGroup) ([]map[string]interface{}, bool) {
    //MsgSecurityAssocSet -> MsgexternalUserGroup
    //MsgSecurityAssocSet -> MsgexternalUserGroup
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.ProviderId != nil {
        val[0]["providerid"] = data.ProviderId
        added = true
    }
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.Id != nil {
        val[0]["id"] = data.Id
        added = true
    }
    if data.ProviderName != nil {
        val[0]["providername"] = data.ProviderName
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgvmdiskfilterpropset_array(d *schema.ResourceData, data []handler.MsgvmDiskFilterPropSet) ([]map[string]interface{}, bool) {
    //MsgvmDiskFilterProp
    //MsgvmDiskFilterPropSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if rtn, ok := serialize_vmgroup_v2_msgvmdiskfilterset_array(d, data[i].Rules); ok {
            tmp["rules"] = rtn
            added = true
        }
        if data[i].Overwrite != nil {
            tmp["overwrite"] = strconv.FormatBool(*data[i].Overwrite)
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgvmdiskfilterset_array(d *schema.ResourceData, data []handler.MsgvmDiskFilterSet) ([]map[string]interface{}, bool) {
    //MsgvmDiskFilterProp -> MsgvmDiskFilterSet
    //MsgvmDiskFilterPropSet -> MsgvmDiskFilterSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Condition != nil {
            tmp["condition"] = data[i].Condition
            added = true
        }
        if data[i].VmName != nil {
            tmp["vmname"] = data[i].VmName
            added = true
        }
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].FilterType != nil {
            tmp["filtertype"] = data[i].FilterType
            added = true
        }
        if data[i].Overwrite != nil {
            tmp["overwrite"] = strconv.FormatBool(*data[i].Overwrite)
            added = true
        }
        if data[i].Value != nil {
            tmp["value"] = data[i].Value
            added = true
        }
        if data[i].VmGuid != nil {
            tmp["vmguid"] = data[i].VmGuid
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgvmgroupsettings(d *schema.ResourceData, data *handler.MsgvmGroupSettings) ([]map[string]interface{}, bool) {
    //MsgvmGroupSettings
    //MsgvmGroupSettings
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.AutoDetectVMOwner != nil {
        val[0]["autodetectvmowner"] = strconv.FormatBool(*data.AutoDetectVMOwner)
        added = true
    }
    if data.CollectFileDetailsforGranularRecovery != nil {
        val[0]["collectfiledetailsforgranularrecovery"] = strconv.FormatBool(*data.CollectFileDetailsforGranularRecovery)
        added = true
    }
    if data.NoOfReaders != nil {
        val[0]["noofreaders"] = data.NoOfReaders
        added = true
    }
    if data.UseChangedBlockTrackingOnVM != nil {
        val[0]["usechangedblocktrackingonvm"] = strconv.FormatBool(*data.UseChangedBlockTrackingOnVM)
        added = true
    }
    if data.JobStartTime != nil {
        val[0]["jobstarttime"] = data.JobStartTime
        added = true
    }
    if data.UseVMCheckpointSetting != nil {
        val[0]["usevmcheckpointsetting"] = strconv.FormatBool(*data.UseVMCheckpointSetting)
        added = true
    }
    if data.CustomSnapshotResourceGroup != nil {
        val[0]["customsnapshotresourcegroup"] = data.CustomSnapshotResourceGroup
        added = true
    }
    if data.RegionalSnapshot != nil {
        val[0]["regionalsnapshot"] = strconv.FormatBool(*data.RegionalSnapshot)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgguestcredentialinfo(d, data.GuestCredentials); ok {
        val[0]["guestcredentials"] = rtn
        added = true
    }
    if data.VmBackupType != nil {
        val[0]["vmbackuptype"] = data.VmBackupType
        added = true
    }
    if data.IsVMGroupDiskFiltersIncluded != nil {
        val[0]["isvmgroupdiskfiltersincluded"] = strconv.FormatBool(*data.IsVMGroupDiskFiltersIncluded)
        added = true
    }
    if data.DatastoreFreespaceCheck != nil {
        val[0]["datastorefreespacecheck"] = strconv.FormatBool(*data.DatastoreFreespaceCheck)
        added = true
    }
    if data.AllowEmptySubclient != nil {
        val[0]["allowemptysubclient"] = strconv.FormatBool(*data.AllowEmptySubclient)
        added = true
    }
    if data.DatastoreFreespaceRequired != nil {
        val[0]["datastorefreespacerequired"] = data.DatastoreFreespaceRequired
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgresourcetagset_array(d, data.CustomSnapshotTags); ok {
        val[0]["customsnapshottags"] = rtn
        added = true
    }
    if data.IsApplicationAware != nil {
        val[0]["isapplicationaware"] = strconv.FormatBool(*data.IsApplicationAware)
        added = true
    }
    if data.TransportMode != nil {
        val[0]["transportmode"] = data.TransportMode
        added = true
    }
    if data.CollectFileDetailsFromSnapshotCopy != nil {
        val[0]["collectfiledetailsfromsnapshotcopy"] = strconv.FormatBool(*data.CollectFileDetailsFromSnapshotCopy)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgamazoncrossaccount(d, data.CrossAccount); ok {
        val[0]["crossaccount"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgamazoncrossaccount(d *schema.ResourceData, data *handler.MsgAmazonCrossAccount) ([]map[string]interface{}, bool) {
    //MsgvmGroupSettings -> MsgAmazonCrossAccount
    //MsgvmGroupSettings -> MsgAmazonCrossAccount
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.ShareOnly != nil {
        val[0]["shareonly"] = strconv.FormatBool(*data.ShareOnly)
        added = true
    }
    if data.FullCopy != nil {
        val[0]["fullcopy"] = strconv.FormatBool(*data.FullCopy)
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.DestinationAccount); ok {
        val[0]["destinationaccount"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgresourcetagset_array(d *schema.ResourceData, data []handler.MsgresourceTagSet) ([]map[string]interface{}, bool) {
    //MsgvmGroupSettings -> MsgresourceTagSet
    //MsgvmGroupSettings -> MsgresourceTagSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Value != nil {
            tmp["value"] = data[i].Value
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_vmgroup_v2_msgactivitycontroloptions(d *schema.ResourceData, data *handler.MsgActivityControlOptions) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions
    //MsgActivityControlOptions
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgactivitycontroloptionsprop(d, data.RestoreActivityControlOptions); ok {
        val[0]["restoreactivitycontroloptions"] = rtn
        added = true
    }
    if rtn, ok := serialize_vmgroup_v2_msgactivitycontroloptionsprop(d, data.BackupActivityControlOptions); ok {
        val[0]["backupactivitycontroloptions"] = rtn
        added = true
    }
    if data.EnableBackup != nil {
        val[0]["enablebackup"] = strconv.FormatBool(*data.EnableBackup)
        added = true
    }
    if data.EnableRestore != nil {
        val[0]["enablerestore"] = strconv.FormatBool(*data.EnableRestore)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgactivitycontroloptionsprop(d *schema.ResourceData, data *handler.MsgActivityControlOptionsProp) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgactivitycontroltiledelaytime(d, data.DelayTime); ok {
        val[0]["delaytime"] = rtn
        added = true
    }
    if data.ActivityType != nil {
        val[0]["activitytype"] = data.ActivityType
        added = true
    }
    if data.EnableAfterADelay != nil {
        val[0]["enableafteradelay"] = strconv.FormatBool(*data.EnableAfterADelay)
        added = true
    }
    if data.EnableActivityType != nil {
        val[0]["enableactivitytype"] = strconv.FormatBool(*data.EnableActivityType)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_vmgroup_v2_msgactivitycontroltiledelaytime(d *schema.ResourceData, data *handler.MsgActivityControlTileDelayTime) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp -> MsgActivityControlTileDelayTime
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp -> MsgActivityControlTileDelayTime
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_vmgroup_v2_msgidname(d, data.TimeZone); ok {
        val[0]["timezone"] = rtn
        added = true
    }
    if data.Time != nil {
        val[0]["time"] = data.Time
        added = true
    }
    if data.Value != nil {
        val[0]["value"] = data.Value
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
