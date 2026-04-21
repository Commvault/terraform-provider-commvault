---
page_title: "Commvault: commvault_oracle_backup Resource"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_backup resource to trigger Oracle database backup operations.
---

# commvault_oracle_backup (Resource)

Use the commvault_oracle_backup resource to trigger Oracle database backup jobs. This resource supports full and incremental backups with configurable RMAN options.

## Syntax

```hcl
resource "commvault_oracle_backup" "<local_name>" {
  client_name    = "<Client Name>"
  instance_name  = "<Oracle SID>"
  subclient_name = "<Subclient Name>"
  backup_type    = "<full|incremental>"
}
```

## Example Usage

### Full Database Backup

```hcl
resource "commvault_oracle_backup" "full_backup" {
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  subclient_name = "default"
  backup_type    = "full"
}
```

### Incremental Backup

```hcl
resource "commvault_oracle_backup" "incr_backup" {
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  subclient_name = "default"
  backup_type    = "incremental"
}
```

### Backup with Specific Subclient

```hcl
resource "commvault_oracle_backup" "subclient_backup" {
  client_name    = "oracle-server-01"
  instance_name  = "PRODDB"
  subclient_name = "archive_logs"
  backup_type    = "full"
}
```

## Argument Reference

### Required

- **client_name** (String) - Name of the client where the Oracle instance is configured.
- **instance_name** (String) - Name of the Oracle instance (SID).

### Optional

- **subclient_name** (String) - Name of the subclient to backup. If not specified, the default subclient is used.
- **backup_type** (String) - Type of backup to perform. Valid values: `full`, `incremental`. Default: `full`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The ID of this resource.
- **job_id** (Number) - The job ID of the backup operation.
- **task_id** (Number) - The task ID of the backup operation.

## Notes

- This resource triggers a backup job when created. The resource will wait for the job to be submitted (not completed).
- To monitor backup job status, use the Commvault Command Center or API.
- Destroying this resource does not affect any completed backup jobs or backup data.
