---
page_title: "commvault_oracle_subclient Data Source"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_subclient data source to retrieve information about Oracle subclients.
---

# commvault_oracle_subclient (Data Source)

Use the commvault_oracle_subclient data source to retrieve information about Oracle subclients configured in Commvault.

## Example Usage

### Lookup by Client, Instance, and Subclient Name

```hcl
data "commvault_oracle_subclient" "default" {
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  subclient_name = "default"
}

output "subclient_id" {
  value = data.commvault_oracle_subclient.default.id
}
```

### Get Subclient Properties

```hcl
data "commvault_oracle_subclient" "backup" {
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  subclient_name = "full_backup"
}

output "storage_policy" {
  value = data.commvault_oracle_subclient.backup.storage_policy_name
}
```

## Argument Reference

### Required

- **client_name** (String) - Name of the client where the Oracle instance is configured.
- **instance_name** (String) - Name of the Oracle instance (SID).
- **subclient_name** (String) - Name of the subclient.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The unique identifier of the Oracle subclient.
- **subclient_id** (Number) - The ID of the subclient.
- **client_id** (Number) - The ID of the client.
- **instance_id** (Number) - The ID of the instance.
- **enable_backup** (Boolean) - Whether backup is enabled for this subclient.
- **description** (String) - Description of the subclient.
- **backup_archive_log** (Boolean) - Whether archive log backup is enabled.
- **backup_sp_file** (Boolean) - Whether SPFILE backup is enabled.
- **backup_control_file** (Boolean) - Whether control file backup is enabled.
- **delete_archive_log_after_backup** (Boolean) - Whether archive logs are deleted after backup.
- **storage_policy_name** (String) - Name of the data backup storage policy.
- **log_storage_policy_name** (String) - Name of the log backup storage policy.
