---
page_title: "Commvault: commvault_oracle_subclient Resource"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_subclient resource to create, update, or delete Oracle subclients for backup operations.
---

# commvault_oracle_subclient (Resource)

Use the commvault_oracle_subclient resource to manage Oracle subclients. Subclients allow you to organize and configure backup content, storage policies, and backup options for Oracle databases.

## Syntax

```hcl
resource "commvault_oracle_subclient" "<local_name>" {
  subclient_name                  = "<Subclient Name>"
  client_name                     = "<Client Name>"
  instance_name                   = "<Oracle SID>"
  enable_backup                   = <true|false>
  description                     = "<Description>"
  backup_archive_log              = <true|false>
  backup_sp_file                  = <true|false>
  backup_control_file             = <true|false>
  delete_archive_log_after_backup = <true|false>

  storage_policy {
    name = "<Storage Policy Name>"
  }

  log_storage_policy {
    name = "<Log Storage Policy Name>"
  }
}
```

## Example Usage

### Basic Oracle Subclient

```hcl
resource "commvault_oracle_subclient" "default_subclient" {
  subclient_name = "default_backup"
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  enable_backup  = true
}
```

### Oracle Subclient with Full Configuration

```hcl
resource "commvault_oracle_subclient" "full_backup" {
  subclient_name                  = "full_backup_subclient"
  client_name                     = "oracle-server-01"
  instance_name                   = "PRODDB"
  enable_backup                   = true
  description                     = "Full database backup subclient"
  backup_archive_log              = true
  backup_sp_file                  = true
  backup_control_file             = true
  delete_archive_log_after_backup = true

  storage_policy {
    name = "Oracle_DataBackup_Policy"
  }

  log_storage_policy {
    name = "Oracle_LogBackup_Policy"
  }
}
```

## Argument Reference

### Required

- **subclient_name** (String) - Name of the subclient.
- **client_name** (String) - Name of the client where the Oracle instance is configured.
- **instance_name** (String) - Name of the Oracle instance (SID).

### Optional

- **enable_backup** (Boolean) - Whether backup is enabled for this subclient. Default: `true`.
- **description** (String) - Description of the subclient.
- **backup_archive_log** (Boolean) - Whether to backup archive logs. Default: `false`.
- **backup_sp_file** (Boolean) - Whether to backup the SPFILE. Default: `false`.
- **backup_control_file** (Boolean) - Whether to backup the control file. Default: `false`.
- **delete_archive_log_after_backup** (Boolean) - Whether to delete archive logs after successful backup. Default: `false`.
- **storage_policy** (Block) - Data backup storage policy configuration.
  - **name** (String) - Name of the storage policy.
  - **id** (Number) - ID of the storage policy.
- **log_storage_policy** (Block) - Log backup storage policy configuration.
  - **name** (String) - Name of the log storage policy.
  - **id** (Number) - ID of the log storage policy.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The unique identifier of the Oracle subclient in Commvault.

## Import

Oracle subclients can be imported using the subclient ID:

```shell
terraform import commvault_oracle_subclient.full_backup <subclient_id>
```
