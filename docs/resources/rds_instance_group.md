---
page_title: "Commvault: commvault_rds_instance_group Resource"
subcategory: "Amazon RDS"
description: |-
  Use the commvault_rds_instance_group resource to manage Amazon RDS and Aurora snapshot protection groups.
---

# commvault_rds_instance_group (Resource)

Use the commvault_rds_instance_group resource to manage an explicitly created
content group (Commvault subclient) under an RDS registration. Select regions,
individual RDS databases, or Aurora clusters.

## Example Usage

### RDS Database Selection

The following example uses a registration defined with the
[RDS instance resource](rds_instance.md):

```hcl
data "commvault_plan" "protection" {
  name = "Database protection"
}

resource "commvault_rds_instance_group" "selected" {
  client_name    = commvault_rds_instance.rds.client_name
  instance_name  = commvault_rds_instance.rds.instance_name
  display_name   = "existing-aws-account"
  subclient_name = "production-databases"
  plan_name      = data.commvault_plan.protection.name

  content {
    display_name = "production-db"
    name         = "production-db"
    path         = "us-east-1"
    type         = "Instance"
    update_type  = 50
    value        = "us-east-1/production-db"
  }
}
```

### Aurora Cluster Selection

Aurora snapshot protection uses the same RDS registration and group resources,
with `Cluster` content:

```hcl
data "commvault_plan" "protection" {
  name = "Database protection"
}

resource "commvault_rds_instance_group" "aurora" {
  client_name              = commvault_rds_instance.rds.client_name
  instance_name            = commvault_rds_instance.rds.instance_name
  display_name             = "existing-aws-account"
  subclient_name           = "production-clusters"
  plan_name                = data.commvault_plan.protection.name
  number_of_backup_streams = 4
  wait_for_snapshot        = false

  content {
    display_name = "production-cluster"
    name         = "production-cluster"
    path         = "us-east-1"
    type         = "Cluster"
    update_type  = 48
    value        = "us-east-1/production-cluster"
  }
}
```

Use the [commvault_plan data source](../data-sources/commvault_plan.md) to look
up an existing compatible plan by name and pass its `name` to `plan_name`.
A literal plan name is also accepted. The examples show alternative groups;
declare a shared plan lookup only once when combining them.

## Argument Reference

### Required

- **client_name** (String) - Parent Commvault cloud account name. Changing this requires replacement.
- **instance_name** (String) - Parent Commvault RDS registration name. May follow a rename of the existing parent; moving a group to another instance is unsupported.
- **display_name** (String) - Cloud account display name. Changing this requires replacement.
- **subclient_name** (String) - Group name. Supports in-place rename.
- **plan_name** (String) - Existing compatible backup plan name. Supports in-place update.
- **content** (Block List) - One or more flat content selections. Supports in-place update.

### Optional

- **number_of_backup_streams** (Number) - Positive backup stream count. Omission retains the server setting.
- **job_start_time** (Number) - Job start time in seconds since midnight, from `0` to `86399`. Omission retains the server setting; `0` means midnight.
- **auto_update_list** (Boolean) - Whether to automatically update the snapshot list. Omission retains the server setting.
- **enable_external_snapshot_discovery** (Boolean) - Whether to discover external snapshots. Omission retains the server setting.
- **enable_snapshot_backup_copy** (Boolean) - Whether to enable snapshot backup copy. Omission retains the server setting.
- **wait_for_snapshot** (Boolean) - Whether to wait for snapshot completion. Omission retains the server setting; explicit `false` is supported.

### Nested `content` Arguments

- **display_name** (String, Required) - Display name of the selected region, database, or cluster.
- **name** (String, Required) - Selection name.
- **path** (String, Required) - Parent path of the selection, such as the AWS region for a database or cluster.
- **type** (String, Required) - `REGION`, `Instance`, or `Cluster`.
- **update_type** (Number, Required) - Numeric type matching `type`: `4` for `REGION`, `50` for `Instance`, or `48` for `Cluster`.
- **value** (String, Required) - Selection value for updates: the region for a region selection, or `region/database` or `region/cluster` for an individual selection.
- **all_or_any_children** (Boolean, Optional) - Whether the selection includes all or any children. Default: `true`.
- **negation** (Boolean, Optional) - Whether the selection is negated. Default: `false`.

Use selection names and paths matching the content discovered by Commvault.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- **id** (String) - Commvault subclient ID.
- **subclient_id** (Number) - Commvault subclient ID.
- **instance_id** (Number) - Parent Commvault RDS instance ID.
- **client_id** (Number) - Parent cloud account client ID.

## Usage Notes

Only flat REGION/Instance/Cluster selections are supported. Nested selections
and default all-database content cannot be managed by this resource.
Do not adopt the registration's default subclient; deletion of the default
subclient is blocked.

Destroy removes the managed group, not its parent registration or AWS
databases. There is no separate group-deconfigure operation. Applying group
settings does not submit an immediate backup or restore job.

## Import

Managed RDS and Aurora groups can be imported using their numeric Commvault
subclient ID:

```shell
terraform import commvault_rds_instance_group.selected 456
```

Import only a group with supported flat content into a matching configuration
and an untracked Terraform address. Do not manage the same group in multiple
states.
