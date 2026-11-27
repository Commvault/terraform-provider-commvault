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

### Cross-Account and Cross-Region Snapshot Protection

The following group shares snapshots with an existing destination cloud account
and overrides region mappings on an existing snapshot replica of its plan:

```hcl
resource "commvault_rds_instance_group" "replicated" {
  client_name                  = commvault_rds_instance.rds.client_name
  instance_name                = commvault_rds_instance.rds.instance_name
  display_name                 = "existing-aws-account"
  subclient_name               = "replicated-databases"
  plan_name                    = data.commvault_plan.protection.name
  is_cross_account_enabled     = true
  is_cross_account_copy_enabled = false

  destination_account {
    client_id = 900
  }

  snapshot_replica {
    copy_id = 2860

    region_mapping {
      source_region      = "Asia Pacific (Mumbai)"
      destination_region = "US East (N. Virginia)"
    }
  }

  content {
    display_name = "Region"
    name         = "ap-south-1"
    path         = ""
    type         = "REGION"
    update_type  = 4
    value        = "ap-south-1"
  }
}
```

Replace `900` with the destination's Commvault cloud account client ID and `2860`
with the assigned plan's secondary snapshot replica copy ID. These are not AWS
account or region IDs. The same settings apply to Aurora snapshot groups.

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
- **is_cross_account_enabled** (Boolean) - Share snapshots with the destination AWS account (share-only mode). Mutually exclusive with full-copy mode. Omission retains the server setting.
- **is_cross_account_copy_enabled** (Boolean) - Make a full snapshot copy in the destination AWS account. Mutually exclusive with share-only mode. Set both flags explicitly when switching modes; set both to `false` to disable cross-account operations. Omission retains the server setting.
- **destination_account** (Block List, Max: 1) - Existing destination account reference. Required when either cross-account mode is enabled. Omission retains the server reference.
- **snapshot_replica** (Block List, Max: 1) - Complete subclient-specific region override set on an existing secondary snapshot replica of this group's assigned plan. Omission skips replica management unless removing a previously managed block.

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

### Nested `destination_account` Arguments

- **client_id** (Number, Required) - Positive Commvault destination cloud account client ID, not an AWS account number or credential ID.

### Nested `snapshot_replica` Arguments

- **copy_id** (Number, Required) - Positive backup destination copy ID of a secondary snapshot replica on the assigned plan. Disk, primary snapshot, and mirror copies are rejected.
- **region_mapping** (Block Set, Required) - One or more source/destination region pairs. Each source may have only one destination.

### Nested `snapshot_replica.region_mapping` Arguments

- **source_region** (String, Required) - Source cloud region display name, such as `Asia Pacific (Mumbai)`.
- **destination_region** (String, Required) - Destination cloud region display name, such as `US East (N. Virginia)`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

- **id** (String) - Commvault subclient ID.
- **subclient_id** (Number) - Commvault subclient ID.
- **instance_id** (Number) - Parent Commvault RDS instance ID.
- **client_id** (Number) - Parent cloud account client ID.
- **destination_account.client_name** (String) - Destination cloud account name, when returned by Commvault.

## Usage Notes

Only flat REGION/Instance/Cluster selections are supported. Nested selections
and default all-database content cannot be managed by this resource.
Do not adopt the registration's default subclient; deletion of the default
subclient is blocked.

Destroy removes the managed group, not its parent registration or AWS
databases. There is no separate group-deconfigure operation. Applying group
settings does not submit an immediate backup or restore job.

### Cross-Account Settings

Cross-account and cross-region settings can be used independently. Set
`is_cross_account_enabled = false` and `is_cross_account_copy_enabled = true`
for full-copy mode. Set both flags explicitly when switching modes and set both
to `false` to disable cross-account operations.

The destination account must already be registered with the appropriate AWS
permissions and snapshot/KMS sharing access. These group settings do not create
AWS accounts, IAM policies, KMS grants, or a plan replica copy.

### Snapshot Replica Mappings

Create the replica using [commvault_plan_server](plan_server.md), or select an
existing replica. Copy IDs are available as `backupDestinations[].planBackupDestination.id`
from `GET /V4/ServerPlan/{planId}`. Configure `copy_id` explicitly; replicas are
not automatically discovered.

Use the `regionName` returned by `GET /Cloud/Regions?vendor=4`. AWS region codes
such as `ap-south-1` remain appropriate for `content` selections, but not for
`region_mapping` values.

Updates delete removed sources and set desired sources for this subclient and
copy only. Copy-level defaults and other subclients' mappings are not modified.
Remove the entire replica block and apply to clear these overrides and resume
plan-level inheritance. Changing `copy_id` clears overrides on the old copy
before setting the new copy. Remove the block and apply before changing
`plan_name`, then configure the new plan's replica in a subsequent apply.

Reads return effective mappings: the backend falls back to copy-level defaults
when no subclient override exists and does not expose mapping provenance.
Identical inherited and explicit mappings therefore cannot be distinguished as
drift. On import, configure the replica block explicitly to manage its mappings.

Replica verification requires read access to
`GET /V5/ServerPlan/{planId}/BackupDestination/{copyId}`. Mapping management uses
`GET /Replication/{copyId}/RegionMapping?subclientId={id}` and the corresponding
POST endpoint. Creation applies settings and mappings after obtaining the
subclient ID; failures retain that ID for repair or cleanup. Writes are
confirmed against server readback and are not atomic across endpoints.

Applying these settings does not run a snapshot replication job or verify
cross-account snapshot sharing, IAM/KMS permissions, or recoverability.

## Import

Managed RDS and Aurora groups can be imported using their numeric Commvault
subclient ID:

```shell
terraform import commvault_rds_instance_group.selected 456
```

Import only a group with supported flat content into a matching configuration
and an untracked Terraform address. Do not manage the same group in multiple
states.
