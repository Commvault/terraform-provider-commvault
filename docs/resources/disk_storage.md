---
page_title: " Commvault : commvault_disk_storage Resource"
subcategory: "Deprecated"
description: |-
  Use the commvault_disk_storage resource type to create or delete a Disk Storage in the Commcell environment.
---

# commvault_disk_storage (Resource)

Use the commvault_disk_storage resource type to create or delete a Disk Storage in the Commcell environment. This is deprecated in the latest version. Starting SP32, use commvault_storage_disk resource.



## Syntax

```
resource "commvault_disk_storage" "<local name>"{
	storage_name = "<Storage Name>"
	mediaagent = "<Media Agent Name>"
	backup_location = "<Backup Location>"
	ddb_location = "<DDB Location>"
	company_id = <Company ID>
}

```

## Example Usage

```
resource "commvault_disk_storage" "DSD"{
	storage_name = "DemoDiskStorage-Dedup"
	mediaagent = "MediaAgent1"
	backup_location = "c:\\DemoDiskStorage-Dedup"
	ddb_location = "c:\\DemoDiskStorage-Dedup-DDB"
	company_id = 22
}

```

### Required

- **storage_name** (String) Specifies the Name of the Disk Storage.
- **mediaagent** (String) Specifies the Media agent used for the Disk Storage.
- **backup_location** (String) Specifies the full path to the storage location.

### Optional

- **ddb_location** (String) Specifies the Deduplication path for the storage
- **company_id** (Number) Specifies the company id to which the created disk storage should be associated with.
- **id** (String) The ID of this resource.


