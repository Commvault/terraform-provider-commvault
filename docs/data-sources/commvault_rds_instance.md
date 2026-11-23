---
page_title: "commvault_rds_instance Data Source"
subcategory: "Amazon RDS"
description: |-
  Use the commvault_rds_instance data source to retrieve an Amazon RDS or Aurora snapshot registration.
---

# commvault_rds_instance (Data Source)

Reads an existing Commvault RDS or Aurora snapshot registration by its numeric
Cloud Apps instance ID. This lookup does not create or manage the registration.

## Example Usage

```hcl
data "commvault_rds_instance" "existing" {
  instance_id = 123
}

output "rds_registration_name" {
  value = data.commvault_rds_instance.existing.instance_name
}
```

## Argument Reference

### Required

- **instance_id** (Number) - Positive Commvault Cloud Apps instance ID, not an AWS database or client ID.

## Attribute Reference

- **id** (String) - Commvault Cloud Apps instance ID.
- **client_id** (Number) - Parent Commvault cloud account client ID.
- **client_name** (String) - Parent cloud account name.
- **instance_name** (String) - Commvault registration name.
- **plan_name** (String) - Assigned backup plan name.
- **access_node** (Block List) - Access nodes, each containing `client_id`, `client_name`, and `type`.
- **credential_name** (String) - Associated credential name, when available.
- **use_iam_role** (Boolean) - IAM-role setting, when available.
- **number_of_backup_streams** (Number) - Configured backup stream count.
- **region_end_points** (String) - Region endpoint settings.
- **point_in_time_restore_enabled** (Boolean) - PITR protection setting.
- **point_in_time_restore_retention_period** (Number) - Automated-backup retention period in days.
- **point_in_time_restore_backup_period** (String) - Automated-backup window.
- **point_in_time_restore_apply_immediately** (Boolean) - Immediate-application setting for automated-backup changes.

Field meanings correspond to the [resource](../resources/rds_instance.md).
Creation-only `default_subclient_id` and `use_resource_pool_info` are not exposed.
Optional properties are populated when returned by Commvault.

## Usage Notes

A missing instance, a non-RDS workload, or an invalid response fails lookup.
Aurora snapshot registrations use this same data source.
