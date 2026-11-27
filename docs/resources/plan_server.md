---
page_title: "Commvault: commvault_plan_server Resource"
subcategory: "Plans"
description: |-
  Use the commvault_plan_server resource type to create, update, or delete a server plan in the CommCell environment.

---

# commvault_plan_server (Resource)

Use the commvault_plan_server resource type to create, update, or delete a server plan in the CommCell environment.

## Backup Destinations

Define backup destinations **inline** using the `backupdestinations` block. The provider creates the plan with its primary destinations; snapshot replicas are added after their source copies exist.

```hcl
resource "commvault_plan_server" "my_plan" {
  planname = "MyBackupPlan"

  backupdestinations {
    backupdestinationname = "Primary-Destination"
    retentionperioddays   = 30
    storagepool {
      id = data.commvault_storagepool.my_pool.id
    }
  }
  # ... other plan settings
}
```

> **Note:** The `backupdestinationids` attribute (referencing pre-created `commvault_plan_backupdestination` resources) is not supported by the API. Always use inline `backupdestinations` blocks.

### Snapshot Replica Copies

Use the same inline destination configuration when creating a plan with a snapshot
replica or adding a replica to an existing Terraform-managed plan:

```hcl
data "commvault_storagepool" "pool" {
  name = "Snapshot-Capable-Pool"
}

resource "commvault_plan_server" "snapshot_plan" {
  planname = "Snapshot-Plan"

  # Used for subsequent destination additions, not for initial plan creation.
  backupdestinationsoperationtype = "ADD"

  backupdestinations {
    backupdestinationname = "Primary-Destination"
    retentionperioddays   = 14
    storagepool {
      id = data.commvault_storagepool.pool.id
    }
  }

  backupdestinations {
    backupdestinationname = "Primary-Snapshot"
    issnapcopy            = "true"
    storagepool {
      id = data.commvault_storagepool.pool.id
    }
  }

  backupdestinations {
    backupdestinationname = "Snapshot-Replica"
    issnapcopy            = "true"
    ismirrorcopy          = "false"
    retentionperioddays   = 14
    sourcecopy {
      name = "Primary-Snapshot"
    }
    storagepool {
      id = data.commvault_storagepool.pool.id
    }
  }
}
```

Creation with snapshot replicas is staged: create the plan and primary copies,
resolve each replica's source copy ID, then send `ADD` for only the replicas.
The create API does not reliably create multiple snapshot copies in one request.
If the follow-up addition fails, the provider reports the error and retains the
new plan's ID for refresh or cleanup; it does not claim success or delete the plan.

The example omits retention on the primary snapshot copy. On the tested SaaS
deployment, explicitly configured primary-snapshot `retentionperioddays` and
`snaprecoverypoints` were returned as zero during staged creation, causing
post-write confirmation to fail. This limitation remains unresolved; confirm
support in your deployment before managing these settings. Replica retention
is configured separately.

`issnapcopy = "true"` selects a snapshot copy; `ismirrorcopy = "false"` selects
Vault/Replica rather than a mirror copy. Configure any required `region` and
`mappings` blocks for your deployment. These flags do not enable snapshot support
on an incompatible workload or storage system.

For `AMAZON` and `AZURE` mappings, supply cloud region display names using
`source.name` and `target.name`. Use the `regionName` returned by
`GET /Cloud/Regions?vendor=4` for AWS, not its `regionCode` or IDs from `/Regions`.
For example, the following maps AWS `ap-south-1` to `us-east-1`:

```hcl
mappings {
  vendor = "AMAZON"
  source {
    name = "Asia Pacific (Mumbai)"
  }
  target {
    name = "US East (N. Virginia)"
  }
}
```

For an existing plan, retain its existing destination blocks unchanged and append
the replica block. `ADD` sends only the new destination; it does not resend existing
copies. Every configured destination must have a unique, non-empty
`backupdestinationname`.

Set `backupdestinationsoperationtype` explicitly when changing existing destinations:

- `ADD`: add new destinations while retaining existing blocks unchanged.
- `MODIFY`: change existing destinations without adding, removing, or renaming them.
- `DELETE`: remove destination blocks while retaining the remaining blocks unchanged.
- `OVERWRITE`: submit the complete desired list. This can remove existing copies;
  do not use it as an alternative to `ADD`.
- `NONE` (or an omitted operation): does not authorize a destination change.

Refresh reads destinations from the server, including snapshot/replica flags,
storage, source copy, retention, region, and mappings. Server-populated names do
not create a difference when an entity was configured only by ID. Omitting
false snapshot/mirror flags in a server response is interpreted as `false`.
When an optimized V4 plan response omits a secondary snapshot copy's mappings,
refresh retrieves them from `GET /V5/ServerPlan/{planId}/BackupDestination/{id}`.
The server must support and permit that endpoint; read errors are not hidden.
Omitting `backupdestinations` leaves the server's destination list computed; declaring
blocks manages the complete list. Missing copies and changed configured fields
are reflected in subsequent plans.

The current update contract does not expose `isconfiguredforreplication`,
`backupstarttime`, `optimizeforinstantclone`, or `storagepolicy`. The provider
rejects these fields on destination additions/modifications rather than silently
dropping them. Their creation support is unchanged.
`optimizeforinstantclone` and `storagepolicy` are not included in the current
plan-read model, so their configured values cannot be verified by refresh.
The current SDK also cannot reliably diff changes to the optional/computed
`isdefault`, `copytype`, and `copyprecedence` fields within an existing destination.
Such changes are rejected during planning. Set them during plan creation or when
adding a new destination instead. Their server values are still read back.

An HTTP-success response is not sufficient by itself: component/server errors
and destination settings not confirmed by readback cause an explicit error.
If creation partially succeeds, the plan ID is retained in state; inspect and
refresh the existing plan before retrying.

### RDS and Aurora Group Overrides

Plan `mappings` define copy-level defaults. Use
[commvault_rds_instance_group](rds_instance_group.md) to configure cross-account
settings and subclient-specific region overrides on an existing secondary
snapshot replica. Select its copy ID from
`backupDestinations[].planBackupDestination.id` in
`GET /V4/ServerPlan/{planId}`. The inline Terraform destination blocks do not
export individual copy IDs.

Removing a group's overrides restores inheritance from these defaults without
changing the plan mappings. Applying a plan or group configures protection; it
does not submit an immediate backup or snapshot replication job or verify AWS
permissions and recoverability.

## Example Usage

**Configure commvault server plan with required fields**
```hcl
data "commvault_storagepool" "storagepool1" {
  name = "Test-Terraform-Pool"
}

resource "commvault_plan_server" "Terraform-Plan1" {
  planname = "Terraform-Test-Plan"
  backupdestinations {
    backupdestinationname = "Primary-Test-BkpDst"
    storagepool {
      id = data.commvault_storagepool.storagepool1.id
    }
  }
}
```

**Configure commvault server plan with custom fields**
```hcl
data "commvault_storagepool" "storagepool1" {
  name = "Test-Terraform-Pool"
}

data "commvault_region" "region1" {
    name = "us-east-2"
}

data "commvault_storagepool" "storagepool2" {
  name = "Test-Terraform-Pool-2"
}

resource "commvault_plan_server" "Terraform-Plan3" {
  planname = "Terraform-Test-Plan-Custom-2"
  backupdestinations {
    backupdestinationname = "Primary-Test-BkpDst-custom"
    storagepool {
      id = data.commvault_storagepool.storagepool1.id
    }
    retentionperioddays = 7
    useextendedretentionrules = "true"
    retentionruletype = "RETENTION_PERIOD"
    extendedretentionrules {
      firstextendedretentionrule {
        type                = "HALF_YEARLY_FULLS"
        retentionperioddays = 90
      }
    }
    overrideretentionsettings = "true"
   region {
      id = data.commvault_region.region1.id
    }
  }
  backupdestinations {
    backupdestinationname = "Secondary-Test-BkpDst-custom"
    storagepool {
      id = data.commvault_storagepool.storagepool2.id
    }
    backupstocopy = "All_JOBS"
    sourcecopy {
      name = "Primary-Test-BkpDst-custom"
    }
    region {
      id = data.commvault_region.region1.id
    }
  }
  backupcontent {
    windowsincludedpaths = ["Music", "Pictures"]
    unixincludedpaths    = ["Dropbox", "OneDrive", "Content Edge Drive", "iCloud Drive", "Google Drive", "Desktop"]
    unixexcludedpaths                   = ["Documents", "Downloads"]
    backupsystemstate                   = "true"
    backupsystemstateonlywithfullbackup = "true"
    usevssforsystemstate                = "true"
    windowsexcludedpaths                = ["Videos", "home"]
    windowsfiltertoexcludepaths         = ["C:\\test"]
    unixfiltertoexcludepaths = ["/var"]
  }
  allowplanoverride = "true"
  overriderestrictions {
    rpo           = "OPTIONAL"
    backupcontent = "MUST"
    storagepool   = "NOT_ALLOWED"
  }
  snapshotoptions {
    enablebackupcopy  = "true"
    backupcopyrpomins = 360
  }
  rpo {
    fullbackupwindow {
      dayofweek = ["THURSDAY", "FRIDAY", "SATURDAY"]
      starttime = 0
      endtime   = 86340
    }
    fullbackupwindow {
      dayofweek = ["SUNDAY"]
      starttime = 72000
      endtime   = 86340
    }
    fullbackupwindow {
      dayofweek = ["MONDAY", "TUESDAY"]
      starttime = 3600
      endtime   = 86340
    }
    sla {
      usesystemdefaultsla = "false"
      excludefromsla = "false"
      slaperiod      = 28
    }
    backupfrequency {
      schedules {
        schedulename     = "terraform-sched-custom-1"
        fordatabasesonly = "false"
        schedulepattern {
          frequency             = 2
          schedulefrequencytype = "DAILY"
          starttime             = 82740
        }
        backuptype = "DIFFERENTIAL"
      }
      schedules {
        schedulename     = "terraform-sched-custom-2"
        backuptype       = "FULL"
        fordatabasesonly = "false"
        schedulepattern {
          schedulefrequencytype = "DAILY"
          #starttime             = 75600
          starttime             = 82740
          exceptions {
            ondates          = [4, 6]
          }
          frequency = 1
          timezone {
            id = 67
          }
        }
      }
      schedules {
        schedulename     = "terraform-sched-custom-4"
        backuptype       = "INCREMENTAL"
        fordatabasesonly = "false"
        schedulepattern {
          schedulefrequencytype = "WEEKLY"
          starttime             = 43260
          frequency             = 1
          weeklydays            = ["SUNDAY", "MONDAY", "TUESDAY"]
          timezone {
            id = 67
          }
        }
      }
    }
    backupwindow {
      dayofweek = ["MONDAY", "THURSDAY", "SUNDAY"]
      starttime = 3600
      endtime   = 86340
    }
    backupwindow {
      dayofweek = ["TUESDAY", "WEDNESDAY", "FRIDAY", "SATURDAY"]
      starttime = 0
      endtime   = 86340
    }
  }
}
```

**Configure commvault server plan with inline backupdestinations (primary/secondary copy designation)**
```hcl
data "commvault_storagepool" "primary_pool" {
  name = "Primary-Storage-Pool"
}

data "commvault_storagepool" "secondary_pool" {
  name = "Secondary-Storage-Pool"
}

resource "commvault_plan_server" "plan_with_copies" {
  planname = "Multi-Copy-Plan"
  
  backupdestinations {
    backupdestinationname = "Primary-Copy"
    retentionperioddays   = 30
    storagepool { 
      id = data.commvault_storagepool.primary_pool.id 
    }
    isdefault      = true     # Mark as primary/default copy
    copytype       = "PRIMARY" # Explicit primary designation
    copyprecedence = 0         # Highest priority (primary)
  }
  
  backupdestinations {
    backupdestinationname = "Secondary-Copy"
    retentionperioddays   = 30
    storagepool { 
      id = data.commvault_storagepool.secondary_pool.id 
    }
    backupstocopy = "All_JOBS"           # Synchronous copy type
    backupstarttime = -1                 # Use default start time
    sourcecopy { name = "Primary-Copy" } # Copy from primary
    isdefault      = false               # Not the primary copy
    copytype       = "AUXILIARY"         # Auxiliary/secondary designation
    copyprecedence = 1                   # Lower priority (secondary)
  }
}
```

**Note:** When using inline `backupdestinations` blocks, use `isdefault`, `copytype`, and `copyprecedence` attributes to explicitly control which copy serves as the primary destination and in what order secondary copies should be used. This ensures Commvault applies the correct primary/auxiliary roles based on your Terraform configuration, rather than relying on storage pool capabilities alone.


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `planname` (String) Name of the new plan

### Optional

- `allowplanoverride` (String) Flag to enable overriding of plan. Plan cannot be overriden by default.
- `backupcontent` (Block List) This feature applies only to file system agents (see [below for nested schema](#nestedblock--backupcontent))
- `backupdestinationids` (Set of Number) Legacy destination references; not supported by the server-plan API used here. Use inline `backupdestinations` instead.
- `backupdestinations` (Block Set) Complete managed destination list. Omission leaves the server list computed. Each configured destination must have a unique, non-empty name (see [below for nested schema](#nestedblock--backupdestinations)).
- `backupdestinationsoperationtype` (String) Operation for changing destinations on an existing plan: `NONE`, `ADD`, `MODIFY`, `DELETE`, or `OVERWRITE`. Required only when changing the destination list or its settings. `ADD` sends only new destinations; `OVERWRITE` can remove existing copies. This setting is not sent during initial plan creation.
- `databaseoptions` (Block List) This feature applies only to database agents (see [below for nested schema](#nestedblock--databaseoptions))
- `filesystemaddon` (String) flag to enable backup content association for applicable file system workload.
- `overrideinheritsettings` (Block List) This feature applies to derived plans when inherit mode is optional.Provides user to set entity preference between parent and derived plan. (see [below for nested schema](#nestedblock--overrideinheritsettings))
- `overriderestrictions` (Block List) To allow the derived plans that use this plan as the base plan to override the settings, property allowPlanOverride must be true, and then select one of the options for Storage pool, RPO and backup Content. (see [below for nested schema](#nestedblock--overriderestrictions))
- `parentplan` (Block List) (see [below for nested schema](#nestedblock--parentplan))
- `regiontoconfigure` (Block List) (see [below for nested schema](#nestedblock--regiontoconfigure))
- `rpo` (Block List) Recovery Point Objective (RPO) is the maximum amount of time that data can be lost during a service disruption. Your RPO determines the frequency of your backup jobs. (see [below for nested schema](#nestedblock--rpo))
- `settings` (Block List) (see [below for nested schema](#nestedblock--settings))
- `snapshotoptions` (Block List) This feature applies only to File System Agents (see [below for nested schema](#nestedblock--snapshotoptions))
- `workload` (Block List) (see [below for nested schema](#nestedblock--workload))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--backupcontent"></a>
### Nested Schema for `backupcontent`

Optional:

- `backupsystemstate` (String) Do you want to back up the system state? Applicable only for Windows
- `backupsystemstateonlywithfullbackup` (String) Do you want to back up system state only with full backup? Applicable only if the value of backupSystemState is true
- `forceupdateproperties` (String) Do you want to sync properties on associated subclients even if properties are overriden at subclient level?
- `macexcludedpaths` (Set of String) Paths to exclude for Mac
- `macfiltertoexcludepaths` (Set of String) Paths that are exception to excluded paths for Mac
- `macincludedpaths` (Set of String) Paths to include for Mac
- `macnumberofdatareaders` (Block List) (see [below for nested schema](#nestedblock--backupcontent--macnumberofdatareaders))
- `unixexcludedpaths` (Set of String) Paths to exclude for UNIX
- `unixfiltertoexcludepaths` (Set of String) Paths that are exception to excluded paths for Unix
- `unixincludedpaths` (Set of String) Paths to include for UNIX
- `unixnumberofdatareaders` (Block List) (see [below for nested schema](#nestedblock--backupcontent--unixnumberofdatareaders))
- `usevssforsystemstate` (String) Do you want to back up system state with VSS? Applicable only if the value of backupSystemState is true
- `windowsexcludedpaths` (Set of String) Paths to exclude for Windows
- `windowsfiltertoexcludepaths` (Set of String) Paths that are exception to excluded paths for Windows
- `windowsincludedpaths` (Set of String) Paths to include for Windows
- `windowsnumberofdatareaders` (Block List) (see [below for nested schema](#nestedblock--backupcontent--windowsnumberofdatareaders))

<a id="nestedblock--backupcontent--macnumberofdatareaders"></a>
### Nested Schema for `backupcontent.macnumberofdatareaders`

Optional:

- `count` (Number) Number of data readers.
- `useoptimal` (String) Set optimal number of data readers. if it is set to true, count will be ignored.


<a id="nestedblock--backupcontent--unixnumberofdatareaders"></a>
### Nested Schema for `backupcontent.unixnumberofdatareaders`

Optional:

- `count` (Number) Number of data readers.
- `useoptimal` (String) Set optimal number of data readers. if it is set to true, count will be ignored.


<a id="nestedblock--backupcontent--windowsnumberofdatareaders"></a>
### Nested Schema for `backupcontent.windowsnumberofdatareaders`

Optional:

- `count` (Number) Number of data readers.
- `useoptimal` (String) Set optimal number of data readers. if it is set to true, count will be ignored.









<a id="nestedblock--backupdestinations"></a>
### Nested Schema for `backupdestinations`

Optional:

- `backupdestinationname` (String) Unique, non-empty destination name. Required for every configured destination; used to resolve server copy identity on updates.
- `backupstarttime` (Int) Backup start time in seconds. The time is provided in unix time format.
- `backupstocopy` (String) All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED, MANUAL_JOBS]
- `copytype` (String) Type of copy for this backup destination [PRIMARY, AUXILIARY, etc.] **Use to explicitly designate primary vs auxiliary copies.**
- `copyprecedence` (Int) Precedence/priority of this backup destination copy. Lower numbers = higher precedence. **Use to control copy priority.**
- `isdefault` (Boolean) Set to true to designate this backup destination as the primary/default copy. **Use to explicitly mark the primary copy.**
- `ismirrorcopy` (String) `"true"` for a mirror copy; `"false"` for a Vault/Replica copy. Only considered when `issnapcopy` is `"true"`.
- `issnapcopy` (String) `"true"` for a snapshot copy. Set `sourcecopy` to identify the primary snapshot source when configuring a replica.
- `isconfiguredforreplication` (String) Used if the copy is used for replication group
- `fullbackuptypestocopy` (String) Which type of backup type should be copied for the given backup destination when backup type is not all jobs. Default is LAST while adding new backup destination. [FIRST, LAST]
- `extendedretentionrules` (Block List) If you want to update, specify the whole object. Extended retention rules should be bigger than retention period.
- `netappcloudtarget` (String) Only for snap copy. Enabling this changes SVM Mapping to NetApp cloud targets only.
- `optimizeforinstantclone` (String) Flag to specify if primary storage is copy data management enabled.
- `overrideretentionsettings` (String) Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy.
- `retentionperioddays` (Int) Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes precedence.
- `retentionruletype` (String) Which type of retention rule should be used for the given backup destination [RETENTION_PERIOD, SNAP_RECOVERY_POINTS]
- `snaprecoverypoints` (Int) Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy.
- `storagetype` (String) [ALL, DISK, CLOUD, HYPERSCALE, TAPE]
- `useextendedretentionrules` (String) Use extended retention rules
- `mappings` (Block Set) Snapshot copy mapping configurations (see [below for nested schema](#nestedblock--backupdestinations--mappings)).
- `region` (Block List) Region for the backup destination (see [below for nested schema](#nestedblock--backupdestinations--region)).
- `sourcecopy` (Block List) Source copy details for secondary/auxiliary copies (see [below for nested schema](#nestedblock--backupdestinations--sourcecopy)).
- `storagepolicy` (Block List) Storage policy details
- `storagepool` (Block List) Storage pool details (see [below for nested schema](#nestedblock--backupdestinations--storagepool)).
- `storagetemplatetags` (Block Set) Storage template tags

The fields `backupstarttime`, `isconfiguredforreplication`,
`optimizeforinstantclone`, and `storagepolicy` are not supported when adding or
modifying destinations through the update API, including staged replica
creation. Configure `isdefault`, `copytype`, and `copyprecedence` only when
creating or adding a destination; changes on existing destinations are rejected.

<a id="nestedblock--backupdestinations--mappings"></a>
### Nested Schema for `backupdestinations.mappings`

Optional:

- `mappingtype` (String) Mapping type for Pure Storage replication: `DEFAULT_MAPPING` or `SUBCLIENT_MAPPING`.
- `source` (Block List) Source mapping entity (see [below for nested schema](#nestedblock--backupdestinations--mappings--source)).
- `sourcevendor` (Block List) Source vendor reference with optional `id` (Number) and `name` (String).
- `subclients` (Block Set) Subclient references, each with optional `id` (Number).
- `target` (Block List) Target mapping entity (see [below for nested schema](#nestedblock--backupdestinations--mappings--target)).
- `targetvendor` (Block List) Target vendor reference with optional `id` (Number) and `name` (String).
- `vendor` (String) Snapshot mapping vendor: `NETAPP`, `AMAZON`, `PURE`, or `AZURE`.

For `AMAZON` and `AZURE`, non-empty `source.name` and `target.name` cloud region
display names are required. Numeric region IDs alone are rejected.

<a id="nestedblock--backupdestinations--mappings--source"></a>
### Nested Schema for `backupdestinations.mappings.source`

Optional:

- `id` (Number) Source mapping entity ID.
- `name` (String) Source mapping entity name. Required for cloud mappings, such as `Asia Pacific (Mumbai)` for AWS.

<a id="nestedblock--backupdestinations--mappings--target"></a>
### Nested Schema for `backupdestinations.mappings.target`

Optional:

- `id` (Number) Target mapping entity ID.
- `name` (String) Target mapping entity name. Required for cloud mappings, such as `US East (N. Virginia)` for AWS.

<a id="nestedblock--backupdestinations--region"></a>
### Nested Schema for `backupdestinations.region`

Optional:

- `guid` (String) Region GUID in the creation schema.
- `id` (Number) Destination region ID.
- `name` (String) Destination region name.

<a id="nestedblock--backupdestinations--sourcecopy"></a>
### Nested Schema for `backupdestinations.sourcecopy`

Optional:

- `id` (Number) Existing source copy ID.
- `name` (String) Source copy name. The provider resolves it against the current plan before adding or modifying a destination when `id` is omitted.

<a id="nestedblock--backupdestinations--storagepool"></a>
### Nested Schema for `backupdestinations.storagepool`

Optional:

- `id` (Number) Existing storage pool ID.
- `name` (String) Existing storage pool name.

<a id="nestedblock--databaseoptions"></a>
### Nested Schema for `databaseoptions`

Optional:

- `commitfrequencyinhours` (Number) Commit frequency in hours
- `logbackuprpomins` (Number) Log backup RPO in minutes
- `runfullbackupevery` (Number) Full backup frequency in days
- `usediskcacheforlogbackups` (String) Use disk cache for log backups


<a id="nestedblock--overrideinheritsettings"></a>
### Nested Schema for `overrideinheritsettings`

Optional:

- `backupcontent` (String) Flag to specify if parent or derived plan backupContent should be used when inherit mode is optional. True - derived, False - Base.
- `backupdestination` (String) Flag to specify if parent or derived plan backupDestination should be used when inherit mode is optional. True - derived, False - Base.
- `rpo` (String) Flag to specify if parent or derived plan rpo should be used when inherit mode is optional. True - derived, False - Base.


<a id="nestedblock--overriderestrictions"></a>
### Nested Schema for `overriderestrictions`

Optional:

- `backupcontent` (String) [OPTIONAL, MUST, NOT_ALLOWED]
- `rpo` (String) [OPTIONAL, MUST, NOT_ALLOWED]
- `storagepool` (String) [OPTIONAL, MUST, NOT_ALLOWED]


<a id="nestedblock--parentplan"></a>
### Nested Schema for `parentplan`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--regiontoconfigure"></a>
### Nested Schema for `regiontoconfigure`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--rpo"></a>
### Nested Schema for `rpo`

Optional:

- `backupfrequency` (Block List) (see [below for nested schema](#nestedblock--rpo--backupfrequency))
- `backupwindow` (Block Set) Backup window for incremental backup (see [below for nested schema](#nestedblock--rpo--backupwindow))
- `fullbackupwindow` (Block Set) Backup window for full backup (see [below for nested schema](#nestedblock--rpo--fullbackupwindow))
- `sla` (Block List) A server meets SLA (Service Level Agreement) when all of its subclients have at least one successful backup during the number of days specified at the CommCell, Server Group or plan level. (see [below for nested schema](#nestedblock--rpo--sla))

<a id="nestedblock--rpo--backupfrequency"></a>
### Nested Schema for `rpo.backupfrequency`

Optional:

- `schedules` (Block List) (see [below for nested schema](#nestedblock--rpo--backupfrequency--schedules))

<a id="nestedblock--rpo--backupfrequency--schedules"></a>
### Nested Schema for `rpo.backupfrequency.schedules`

Required:

- `backuptype` (String) Schedule Backup level [FULL, INCREMENTAL, DIFFERENTIAL, SYNTHETICFULL, TRANSACTIONLOG]
- `schedulename` (String) Name of the schedule, for modify
- `schedulepattern` (Block List, Min: 1) Used to describe when the schedule runs (see [below for nested schema](#nestedblock--rpo--backupfrequency--schedules--schedulepattern))

Optional:

- `fordatabasesonly` (String) Boolean to indicate if schedule is for database agents
- `scheduleoption` (Block List) Specific options to be set on schedules (see [below for nested schema](#nestedblock--rpo--backupfrequency--schedules--scheduleoption))
- `vmoperationtype` (String) Type of DR operation (only applicable for Failover groups) [PLANNED_FAILOVER, TEST_BOOT]

<a id="nestedblock--rpo--backupfrequency--schedules--schedulepattern"></a>
### Nested Schema for `rpo.backupfrequency.schedules.schedulepattern`

Required:

- `schedulefrequencytype` (String) schedule frequency type [MINUTES, DAILY, WEEKLY, MONTHLY, YEARLY, AUTOMATIC]

Optional:

- `dayofmonth` (Number) Day on which to run the schedule, applicable for monthly, yearly
- `dayofweek` (String) [SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, DAY, WEEKDAY, WEEKEND_DAYS]
- `daysbetweensyntheticfulls` (Number) No of days between two synthetic full jobs
- `enddate` (Number) Schedule end date in epoch format
- `exceptions` (Block Set) Exceptions to when a schedule should not run, either in dates or week of month and days (see [below for nested schema](#nestedblock--rpo--backupfrequency--schedules--schedulepattern--exceptions))
- `frequency` (Number) Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months
- `maxbackupintervalinmins` (Number) The number of mins to force a backup on automatic schedule.
- `monthofyear` (String) [JANUARY, FEBRUARY, MARCH, APRIL, MAY, JUNE, JULY, AUGUST, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER]
- `nooftimes` (Number) The number of times you want the schedule to run.
- `repeatintervalinminutes` (Number) How often in minutes in a day the schedule runs, applicable for daily, weekly, monthly and yearly frequency types.
- `repeatuntiltime` (Number) Until what time to repeat the schedule in a day, requires repeatIntervalInMinutes
- `startdate` (Number) start date of schedule in epoch format
- `starttime` (Number) start time of schedule in seconds
- `timezone` (Block List) (see [below for nested schema](#nestedblock--rpo--backupfrequency--schedules--schedulepattern--timezone))
- `weeklydays` (Set of String) Days of the week for weekly frequency
- `weekofmonth` (String) Specific week of a month [FIRST, SECOND, THIRD, FOURTH, LAST]

<a id="nestedblock--rpo--backupfrequency--schedules--schedulepattern--exceptions"></a>
### Nested Schema for `rpo.backupfrequency.schedules.schedulepattern.exceptions`

Optional:

- `ondates` (Set of Number) list of dates in a month. For ex: 1, 20
- `ondayoftheweek` (Set of String) On which days, for ex: MONDAY, FRIDAY
- `onweekofthemonth` (Set of String) On which week of month, for ex: FIRST, LAST


<a id="nestedblock--rpo--backupfrequency--schedules--schedulepattern--timezone"></a>
### Nested Schema for `rpo.backupfrequency.schedules.schedulepattern.timezone`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.



<a id="nestedblock--rpo--backupfrequency--schedules--scheduleoption"></a>
### Nested Schema for `rpo.backupfrequency.schedules.scheduleoption`

Optional:

- `commitfrequencyinhours` (Number) Commit frequency in hours for disk cache backups from automatic schedules
- `daysbetweenautoconvert` (Number) Number of days between auto conversion of backup level applicable for databases on incremental and differential schedules of server plan
- `jobrunningtimeinmins` (Number) total job running time in minutes
- `o365itemselectionoption` (String) item backup option for O365 V2 backup jobs [SELECT_ALL, SELECT_NEVER_PROCESSED, SELECT_MEETING_SLA, SELECT_NOT_MEETING_SLA_PROCESSED_ATLEAST_ONCE, SELECT_FAILED_LAST_ATTEMPT, SELECT_PROCESSED_ATLEAST_ONCE, SELECT_NOT_MEETING_SLA, SELECT_MEETING_SLA_NOT_RECENTLY_BACKED_UP]
- `usediskcacheforlogbackups` (String) Used to enable disk caching feature on databases for automatic schedules on server plan




<a id="nestedblock--rpo--backupwindow"></a>
### Nested Schema for `rpo.backupwindow`

Optional:

- `dayofweek` (Set of String)
- `endtime` (Number) Time in seconds since the beginning of the day
- `starttime` (Number) Time in seconds since the beginning of the day


<a id="nestedblock--rpo--fullbackupwindow"></a>
### Nested Schema for `rpo.fullbackupwindow`

Optional:

- `dayofweek` (Set of String)
- `endtime` (Number) Time in seconds since the beginning of the day
- `starttime` (Number) Time in seconds since the beginning of the day


<a id="nestedblock--rpo--sla"></a>
### Nested Schema for `rpo.sla`

Optional:

- `enableafterdelay` (Number) Time provided in Unix format. Give 0 to reset any existing delay.
- `excludefromsla` (String) Flag to set to exclude plan from SLA
- `exclusionreason` (String) Reason for exclusion from SLA
- `slaperiod` (Number) SLA Period in Days
- `usesystemdefaultsla` (String) Flag to set to use System Default Service Level Agreement



<a id="nestedblock--settings"></a>
### Nested Schema for `settings`

Optional:

- `enableadvancedview` (String) Setting to suggest plan has some advanced settings present. Setting is OEM specific and not applicable for all cases.
- `filesearch` (Block List) This feature applies to file servers and virtualization. Enabling this feature allows you to search for backed-up files using the global search bar, and creates resource pools with required infrastructure entities. (see [below for nested schema](#nestedblock--settings--filesearch))

<a id="nestedblock--settings--filesearch"></a>
### Nested Schema for `settings.filesearch`

Optional:

- `enabled` (String) Flag for enabling indexing
- `status` (String) Type of indexing status. [NOT_APPLICABLE, ENABLED, SETUP_IN_PROGRESS]
- `statusmessage` (String) Tells what is happening behind the scene, so that user can knows why indexing is not enabled or if its in progress



<a id="nestedblock--snapshotoptions"></a>
### Nested Schema for `snapshotoptions`

Optional:

- `backupcopyrpomins` (Number) Backup copy RPO in minutes
- `enablebackupcopy` (String) Flag to enable backup copy


<a id="nestedblock--workload"></a>
### Nested Schema for `workload`

Optional:

- `solutions` (Block Set) (see [below for nested schema](#nestedblock--workload--solutions))
- `workloadgrouptypes` (Set of String)
- `workloadtypes` (Block Set) (see [below for nested schema](#nestedblock--workload--workloadtypes))

<a id="nestedblock--workload--solutions"></a>
### Nested Schema for `workload.solutions`

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--workload--workloadtypes"></a>
### Nested Schema for `workload.workloadtypes`

Read-Only:

- `id` (Number) The ID of this resource.


