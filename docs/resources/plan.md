---
page_title: " Commvault : commvault_plan Resource "
subcategory: "Plans"
description: |-
  Use the commvault_plan resource type to create or delete a Plans in the Commcell environment.
---

# commvault_plan (Resource)

 Use the commvault_plan resource type to create or delete a Plans in the Commcell environment.



##Syntax

```
resource "commvault_plan" "<local name>"{
	plan_name = "<Plan Name>"
	retention_period_days = <Retention Period>
	backup_destination_name = "<Backup Destination Name>"
	backup_destination_storage = "<Backup Destination Storage>"
	company_id = <Company ID>
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
}
```

### Required

- **plan_name** (String) Specifies the Plan name used for creation of the plan.
- **retention_period_days** (Number) Specifies the number of days that the software retains the data.
- **backup_destination_name** (String) Specifies the destination name for the backup.
- **backup_destination_storage** (String) Specifies the backup destination storage used for the plan.

### Optional

- **company_id** (Number) Specifies the companyid to which the created plan needs to be associated with.
- **id** (String) The ID of this resource.



