---
page_title: " Commvault : commvault_plan_server Resource "
subcategory: "Plans"
description: |-
  Use the commvault_plan_server resource type to create or delete server plan in the CommCell environment.

---

# commvault_plan_server (Resource)

Use the commvault_plan_server resource type to create or delete server plan in the CommCell environment.

## Example Usage

**Configure commvault server plan with required fields**
```hcl
data "commvault_storagepool" "storagepool1" {
  name = "Test-Terraform-Pool"
}

resource "commvault_plan_backupdestination" "plan_backupdestination1" {
  name = "Primary-Test-BkpDst"
   storagepool {
    id = data.commvault_storagepool.storagepool1.id
   }
}

resource "commvault_plan_server" "Terraform-Plan1" {
  planname = "Terraform-Test-Plan"
  backupdestinationids = [commvault_plan_backupdestination.plan_backupdestination1.id]
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

resource "commvault_plan_backupdestination" "plan_backupdestination2" {
   name = "Primary-Test-BkpDst-custom"
   storagepool {
    id = data.commvault_storagepool.storagepool1.id
   }
   retentionperioddays       = 7
   enabledataaging = "true"
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

resource "commvault_plan_backupdestination" "plan_backupdestination3" {
   name = "Secondary-Test-BkpDst-custom"
   storagepool {
    id = data.commvault_storagepool.storagepool2.id
   }
   backupstocopy = "All_JOBS"
   backupstarttime = -1
   sourcecopy {
    id = commvault_plan_backupdestination.plan_backupdestination2.id
   }
   region {
      id = data.commvault_region.region1.id
    }
}

resource "commvault_plan_server" "Terraform-Plan3" {
  planname = "Terraform-Test-Plan-Custom-2"
  backupdestinationids = [commvault_plan_backupdestination.plan_backupdestination2.id, commvault_plan_backupdestination.plan_backupdestination3.id]
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


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `planname` (String) Name of the new plan
- `backupdestinationids` (Set of Number) Primary Backup Destination Ids (which were created before plan creation). This is only considered when backupDestinations array object is not defined.

### Optional

- `allowplanoverride` (String) Flag to enable overriding of plan. Plan cannot be overriden by default.
- `backupcontent` (Block List) This feature applies only to file system agents (see [below for nested schema](#nestedblock--backupcontent))
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


