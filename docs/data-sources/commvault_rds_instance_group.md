---
page_title: "commvault_rds_instance_group Data Source"
subcategory: "Amazon RDS"
description: |-
  Use the commvault_rds_instance_group data source to retrieve an Amazon RDS or Aurora snapshot protection group.
---

# commvault_rds_instance_group (Data Source)

Reads an existing RDS or Aurora snapshot content group by its numeric Commvault
subclient ID. This lookup does not create or manage the group.

## Example Usage

### Group Lookup

```hcl
data "commvault_rds_instance_group" "existing" {
  subclient_id = 456
}

output "rds_group_content" {
  value = data.commvault_rds_instance_group.existing.content
}
```

### Group Lookup with Snapshot Replica Mappings

```hcl
data "commvault_rds_instance_group" "with_replica" {
  subclient_id = 456

  snapshot_replica {
    copy_id = 2860
  }
}
```

## Argument Reference

### Required

- **subclient_id** (Number) - Positive Commvault subclient ID.

### Optional

- **snapshot_replica** (Block List, Max: 1) - Explicit replica selection for reading effective subclient region mappings. Omission skips replica lookup.

### Nested `snapshot_replica` Arguments

- **copy_id** (Number, Required) - Positive copy ID of a secondary snapshot replica on the group's assigned plan. Disk, primary snapshot, and mirror copies are rejected.

## Attribute Reference

- **id** (String) - Commvault subclient ID.
- **instance_id** (Number) - Parent Commvault RDS instance ID.
- **client_id** (Number) - Parent cloud account client ID.
- **client_name** (String) - Parent cloud account name.
- **instance_name** (String) - Current parent registration name.
- **display_name** (String) - Cloud account display name.
- **subclient_name** (String) - Group name.
- **plan_name** (String) - Assigned backup plan name.
- **number_of_backup_streams** (Number) - Configured backup stream count.
- **job_start_time** (Number) - Job start time in seconds since midnight.
- **auto_update_list** (Boolean) - Snapshot-list automatic-update setting.
- **enable_external_snapshot_discovery** (Boolean) - External snapshot discovery setting.
- **enable_snapshot_backup_copy** (Boolean) - Snapshot backup-copy setting.
- **wait_for_snapshot** (Boolean) - Snapshot completion wait setting.
- **is_cross_account_enabled** (Boolean) - Cross-account share-only setting.
- **is_cross_account_copy_enabled** (Boolean) - Cross-account full-copy setting.
- **destination_account** (Block List) - Destination Commvault cloud account reference, when configured.
- **snapshot_replica.region_mapping** (Block Set) - Effective region mappings for the explicitly selected copy and subclient.
- **content** (Block List) - Flat selections, each containing `display_name`, `name`, `path`, `type`, `update_type`, `value`, `all_or_any_children`, and `negation`.

Field meanings and content type mappings correspond to the
[resource](../resources/rds_instance_group.md). Optional properties are
populated when returned by Commvault.

### Nested `destination_account` Attributes

- **client_id** (Number) - Commvault destination cloud account client ID, not an AWS account number.
- **client_name** (String) - Destination cloud account name, when returned by Commvault.

### Nested `snapshot_replica.region_mapping` Attributes

- **source_region** (String) - Source cloud region display name.
- **destination_region** (String) - Destination cloud region display name.

## Usage Notes

A missing group, a non-RDS workload, or unsupported content fails lookup.
Only flat REGION/Instance/Cluster selections are supported; nested selections
and default all-database content are not supported. Aurora snapshot groups
use this same data source with `Cluster` content.

Cross-account settings are read without a replica selection. Region mappings
require an explicit `snapshot_replica` block and read access to
`GET /V5/ServerPlan/{planId}/BackupDestination/{copyId}` and
`GET /Replication/{copyId}/RegionMapping?subclientId={id}`. Reads inherit
copy-level mappings when no subclient overrides exist; the API does not expose
mapping provenance. This data source does not change mappings or verify
snapshot replication.
