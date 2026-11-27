---
page_title: "Commvault: commvault_rds_instance Resource"
subcategory: "Amazon RDS"
description: |-
  Use the commvault_rds_instance resource to manage Amazon RDS and Aurora snapshot registrations.
---

# commvault_rds_instance (Resource)

Use the commvault_rds_instance resource to register Amazon RDS for protection
under an existing Commvault cloud account. Aurora snapshot protection uses the
same resource. This manages the Commvault registration, not the AWS database
or cluster.

## Example Usage

```hcl
data "commvault_plan" "protection" {
  name = "Database protection"
}

resource "commvault_rds_instance" "rds" {
  client_name   = "existing-aws-account"
  instance_name = "production-rds"
  plan_name     = data.commvault_plan.protection.name

  access_node {
    client_id   = 123
    client_name = "existing-backup-gateway"
    type        = "CLIENT_ENTITY"
  }

  lifecycle {
    prevent_destroy = true
  }
}
```

Use the [commvault_plan data source](../data-sources/commvault_plan.md) to look
up an existing compatible plan by name and pass its `name` to `plan_name`.
A literal plan name is also accepted.

Use existing compatible access nodes. Omit `credential_name`
to inherit the cloud account credential. To manage a different existing
Commvault credential, set `credential_name`; set `use_iam_role = true` only
when using a compatible IAM-role credential.

For selected databases or Aurora clusters, create an
[RDS instance group](rds_instance_group.md) under the registration. Configure
cross-account snapshot sharing/copying and subclient-specific cross-region
snapshot replica mappings on that group, not on this registration.

## Argument Reference

### Required

- **client_name** (String) - Existing Commvault cloud account name. Changing this requires replacement.
- **instance_name** (String) - Commvault registration name, not an AWS database identifier. Supports in-place rename.
- **plan_name** (String) - Name of an existing compatible backup plan. Supports in-place update.

### Optional

- **access_node** (Block List) - Existing access nodes or client groups. Supports in-place update; clearing all configured access nodes is unsupported.
- **credential_name** (String) - Existing Commvault credential to associate with the registration. Omission inherits the cloud account credential on creation. Clearing a managed credential association is unsupported.
- **use_iam_role** (Boolean) - Whether the managed credential uses an IAM role. Default: `false`. Requires `credential_name` when `true`.
- **use_resource_pool_info** (Boolean) - Whether to use resource-pool information. Default: `false`. Changing this requires replacement.
- **number_of_backup_streams** (Number) - Positive backup stream count. Omission retains the server setting.
- **region_end_points** (String) - Region endpoint settings in the format supported by your Commvault deployment. Omission retains the server setting.
- **point_in_time_restore_enabled** (Boolean) - Whether to enable point-in-time restore protection. Omission retains the server setting.
- **point_in_time_restore_retention_period** (Number) - Automated-backup retention period in days, from `0` to `35`. Omission retains the server setting.
- **point_in_time_restore_backup_period** (String) - Automated-backup window in the format supported by your Commvault deployment. Omission retains the server setting.
- **point_in_time_restore_apply_immediately** (Boolean) - Whether to apply changes to the source database's automated-backup settings immediately. Omission retains the server setting.

### Nested `access_node` Arguments

- **client_id** (Number, Required) - Positive Commvault access-node client or client-group ID.
- **client_name** (String, Required) - Matching access-node client or client-group name.
- **type** (String, Optional) - `CLIENT_ENTITY` or `CLIENTGROUP_ENTITY`. Default: `CLIENT_ENTITY`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- **id** (String) - Commvault Cloud Apps instance ID.
- **instance_id** (Number) - Commvault Cloud Apps instance ID, not an AWS database ID.
- **client_id** (Number) - Parent Commvault cloud account client ID.
- **default_subclient_id** (Number) - Default subclient ID returned during creation. It is not rediscovered during import.

## Usage Notes

Creation also creates a default subclient. Its default content protects all
instances in all regions; adding a narrower managed group does not narrow the
default subclient's content.

PITR settings can change automated backups on the source AWS database and
require workload-compatible settings and permissions. Configuring them does
not submit a restore job. Applying this resource does not submit an immediate
backup or verify recoverability.

Destroy deletes the Commvault registration, not the AWS database or cluster.
Remove dependent managed groups first. The example's `prevent_destroy` guard
blocks both deletion and replacement until intentionally removed. To disable
protection explicitly, use
[commvault_cloud_instance_deconfigure](cloud_instance_deconfigure.md).
Deleting a registration does not guarantee erasure of retained backups or
Commvault metadata.

## Import

RDS and Aurora snapshot registrations can be imported using their numeric
Commvault instance ID:

```shell
terraform import commvault_rds_instance.rds 123
```

Create a matching resource configuration before importing. Import into an
untracked Terraform address and do not manage the same registration in multiple
states. If refresh does not return `use_resource_pool_info`, its local value is
retained; configure it to match the existing registration.
