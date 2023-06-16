---
page_title: " Commvault : commvault_plan (deprecated) Resource "
subcategory: "Deprecated"
description: |-
  Use the commvault_plan resource type to create or delete a Plans in the Commcell environment.
---

# commvault_plan (Resource)

 Use the commvault_plan resource type to create or delete a Plans in the Commcell environment. This is deprecated in the latest version. Starting SP32, use commvault_plan_server resource.

## Syntax

```
resource "commvault_plan" "<local name>"{
	plan_name = "<Plan Name>"
	retention_period_days = <Retention Period>
	backup_destination_name = "<Backup Destination Name>"
	backup_destination_storage = "<Backup Destination Storage>"
	company_id = <Company ID>
	rpo_in_days = <RPO in days>
}
```

## Example Usage

```
resource "commvault_plan" "Plan1"{
	plan_name = "Plan1"
	retention_period_days = 15
	backup_destination_name = "Plan1Dest"
	backup_destination_storage = "storagePool1"
	company_id = 22
	rpo_in_days = 5
}
```

### Required

- **plan_name** (String) Specifies the Plan name used for creation of the plan.
- **retention_period_days** (Number) Specifies the number of days that the software retains the data.
- **backup_destination_name** (String) Specifies the destination name for the backup.
- **backup_destination_storage** (String) Specifies the backup destination storage used for the plan.

### Optional

- **company_id** (Number) Specifies the companyid to which the created plan needs to be associated with.
- **rpo_in_days** (Number) Specifies the rpo in Days for created plan. default value is 1.
- **id** (String) The ID of this resource.



