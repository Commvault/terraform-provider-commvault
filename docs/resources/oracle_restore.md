---
page_title: "Commvault: commvault_oracle_restore Resource"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_restore resource to perform Oracle database restore and recovery operations.
---

# commvault_oracle_restore (Resource)

Use the commvault_oracle_restore resource to perform Oracle database restore and recovery operations. This resource supports various restore scenarios including in-place restore, out-of-place restore, point-in-time recovery, and SCN-based recovery.

## Syntax

```hcl
resource "commvault_oracle_restore" "<local_name>" {
  # Source
  source_client_name   = "<Source Client Name>"
  source_instance_name = "<Source Oracle SID>"

  # Destination
  dest_client_name   = "<Destination Client Name>"
  dest_instance_name = "<Destination Oracle SID>"

  # Restore Options
  restore_type = "<in_place|out_of_place|duplicate>"
  
  # Recovery Options
  recover          = <true|false>
  recover_to       = "<current_time|point_in_time|scn>"
  point_in_time    = <epoch_timestamp>
  scn              = "<SCN Value>"
}
```

## Example Usage

### In-Place Restore with Point-in-Time Recovery

```hcl
resource "commvault_oracle_restore" "pit_restore" {
  source_client_name   = "oracle-server-01"
  source_instance_name = "PRODDB"
  
  dest_client_name   = "oracle-server-01"
  dest_instance_name = "PRODDB"
  
  restore_type = "in_place"
  
  restore_control_file = true
  restore_data         = true
  
  recover    = true
  recover_to = "point_in_time"
  point_in_time = 1692939777
  
  reset_logs           = true
  switch_database_mode = true
}
```

### Out-of-Place Restore

```hcl
resource "commvault_oracle_restore" "cross_machine" {
  source_client_name   = "oracle-server-01"
  source_instance_name = "PRODDB"
  
  dest_client_name   = "oracle-server-02"
  dest_instance_name = "DEVDB"
  
  restore_type = "out_of_place"
  
  restore_control_file = true
  restore_data         = true
  
  recover    = true
  recover_to = "scn"
  scn        = "12954669"
  
  redirect_all_tablespaces = true
  redirect_path            = "/u01/oradata/DEVDB"
  
  reset_logs           = true
  switch_database_mode = true
  no_catalog           = true
}
```

### SCN-Based Recovery

```hcl
resource "commvault_oracle_restore" "scn_restore" {
  source_client_name   = "oracle-server-01"
  source_instance_name = "PRODDB"
  
  dest_client_name   = "oracle-server-01"
  dest_instance_name = "PRODDB"
  
  restore_type = "in_place"
  
  restore_control_file = true
  restore_data         = true
  
  recover    = true
  recover_to = "scn"
  scn        = "12954669"
  
  reset_logs = true
}
```

### Archive Log Restore

```hcl
resource "commvault_oracle_restore" "archive_restore" {
  source_client_name   = "oracle-server-01"
  source_instance_name = "PRODDB"
  
  dest_client_name   = "oracle-server-01"
  dest_instance_name = "PRODDB"
  
  restore_archive_log = true
  archive_log_dest    = "/u01/archive_restore"
  
  archive_log_by   = "BYLSN"
  start_lsn        = "1"
  end_lsn          = "100"
}
```

## Argument Reference

### Required

- **source_client_name** (String) - Name of the source client.
- **source_instance_name** (String) - Name of the source Oracle instance (SID).
- **dest_client_name** (String) - Name of the destination client.
- **dest_instance_name** (String) - Name of the destination Oracle instance (SID).

### Optional - Restore Options

- **restore_type** (String) - Type of restore operation. Valid values: `in_place`, `out_of_place`, `duplicate`. Default: `in_place`.
- **restore_control_file** (Boolean) - Whether to restore the control file. Default: `false`.
- **restore_sp_file** (Boolean) - Whether to restore the SPFILE. Default: `false`.
- **restore_data** (Boolean) - Whether to restore database datafiles. Default: `true`.
- **restore_streams** (Number) - Number of RMAN channels/streams for restore. Default: `1`.
- **no_catalog** (Boolean) - Skip RMAN catalog usage during restore. Default: `false`.

### Optional - Recovery Options

- **recover** (Boolean) - Whether to recover the database after restore. Default: `false`.
- **recover_to** (String) - Recovery target type. Valid values: `current_time`, `point_in_time`, `scn`, `most_recent`. Default: `current_time`.
- **point_in_time** (Number) - Epoch timestamp for point-in-time recovery.
- **scn** (String) - System Change Number for SCN-based recovery.
- **reset_logs** (Boolean) - Open database with RESETLOGS after recovery. Default: `false`.
- **open_database** (Boolean) - Open the database after restore/recovery. Default: `true`.
- **switch_database_mode** (Boolean) - Automatically change database mode for restore. Default: `false`.

### Optional - Redirect Options

- **redirect_all_tablespaces** (Boolean) - Redirect all datafiles to a single path. Default: `false`.
- **redirect_path** (String) - Destination path for redirected datafiles.

### Optional - Archive Log Restore

- **restore_archive_log** (Boolean) - Restore archive logs instead of database. Default: `false`.
- **archive_log_dest** (String) - Target path for archive log restore.
- **archive_log_by** (String) - Archive log selection criteria. Valid values: `BYLSN`, `BYTIME`.
- **start_lsn** (String) - Start log sequence number for archive log range.
- **end_lsn** (String) - End log sequence number for archive log range.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The ID of this resource.
- **job_id** (Number) - The job ID of the restore operation.
- **task_id** (Number) - The task ID of the restore operation.

## Notes

- This resource triggers a restore job when created.
- For cross-machine restores, ensure the destination client has the Oracle agent installed.
- When using `reset_logs = true`, a new incarnation of the database is created.
- The `recover` and `restore_archive_log` options are mutually exclusive.
