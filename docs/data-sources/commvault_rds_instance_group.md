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

```hcl
data "commvault_rds_instance_group" "existing" {
  subclient_id = 456
}

output "rds_group_content" {
  value = data.commvault_rds_instance_group.existing.content
}
```

## Argument Reference

### Required

- **subclient_id** (Number) - Positive Commvault subclient ID.

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
- **content** (Block List) - Flat selections, each containing `display_name`, `name`, `path`, `type`, `update_type`, `value`, `all_or_any_children`, and `negation`.

Field meanings and content type mappings correspond to the
[resource](../resources/rds_instance_group.md). Optional properties are
populated when returned by Commvault.

## Usage Notes

A missing group, a non-RDS workload, or unsupported content fails lookup.
Only flat REGION/Instance/Cluster selections are supported; nested selections
and default all-database content are not supported. Aurora snapshot groups
use this same data source with `Cluster` content.
