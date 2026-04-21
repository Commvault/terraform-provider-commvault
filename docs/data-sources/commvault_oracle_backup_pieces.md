---
page_title: "commvault_oracle_backup_pieces Data Source"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_backup_pieces data source to retrieve Oracle RMAN backup pieces information.
---

# commvault_oracle_backup_pieces (Data Source)

Use the commvault_oracle_backup_pieces data source to retrieve information about Oracle RMAN backup pieces for a given instance and time range.

## Example Usage

### Get All Backup Pieces for an Instance

```hcl
data "commvault_oracle_backup_pieces" "recent" {
  instance_id = "208"
}

output "backup_pieces" {
  value = data.commvault_oracle_backup_pieces.recent.pieces
}
```

### Get Backup Pieces for a Time Range

```hcl
data "commvault_oracle_backup_pieces" "last_week" {
  instance_id = "208"
  from_time   = 1692334800  # 7 days ago (epoch)
  to_time     = 1692939600  # now (epoch)
}

output "backup_count" {
  value = length(data.commvault_oracle_backup_pieces.last_week.pieces)
}
```

### Filter by Backup Type

```hcl
data "commvault_oracle_backup_pieces" "full_backups" {
  instance_id = "208"
  from_time   = 1692334800
}

output "full_backup_pieces" {
  value = [for p in data.commvault_oracle_backup_pieces.full_backups.pieces : p if p.backup_type == "FULL"]
}
```

## Argument Reference

### Required

- **instance_id** (String) - The ID of the Oracle instance.

### Optional

- **from_time** (Number) - Start time filter in epoch format.
- **to_time** (Number) - End time filter in epoch format.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The ID of this data source.
- **pieces** (List) - List of backup pieces. Each piece contains:
  - **backup_piece_name** (String) - Name of the backup piece.
  - **tag** (String) - RMAN tag associated with the backup piece.
  - **start_time** (Number) - Start time of the backup in epoch format.
  - **end_time** (Number) - End time of the backup in epoch format.
  - **backup_type** (String) - Type of backup (FULL, INCREMENTAL, ARCHIVELOG).
  - **size** (Number) - Size of the backup piece in bytes.
