---
page_title: "Commvault: commvault_s3_object_storage_subclient Resource"
subcategory: "Amazon S3"
description: |-
  Use the commvault_s3_object_storage_subclient resource to manage Amazon S3 content groups.
---

# commvault_s3_object_storage_subclient (Resource)

Use the commvault_s3_object_storage_subclient resource to manage an Amazon S3
content group under an existing source registration. An existing compatible
backup plan is required.

## Example Usage

### Path-Based Content

```hcl
resource "commvault_s3_object_storage_subclient" "buckets" {
  client_id      = commvault_s3_object_storage.source.client_id
  instance_id    = commvault_s3_object_storage.source.instance_id
  subclient_name = "production-buckets"
  plan_id        = 42
  collect_acl    = false
  job_start_time = 75600

  content {
    path = "/production-bucket"
  }
  content {
    exclude_path = "/production-bucket/temporary"
  }
  content {
    include_path = "/production-bucket/temporary/retain.txt"
  }
}
```

### Tag-Based Content

```hcl
resource "commvault_s3_object_storage_subclient" "tagged" {
  client_id      = commvault_s3_object_storage.source.client_id
  instance_id    = commvault_s3_object_storage.source.instance_id
  subclient_name = "production-tagged-buckets"
  plan_id        = 42

  tag_filter {
    rules {
      name  = "Environment"
      value = "Production"
    }
  }
}
```

## Argument Reference

### Required

- **client_id** (Number) - Positive Commvault client ID. Changing this requires replacement.
- **instance_id** (Number) - Positive S3 instance ID. Changing this requires replacement.
- **subclient_name** (String) - Content-group name. Supports in-place rename.
- **plan_id** (Number) - ID of an existing compatible backup plan.

### Optional

- **number_of_backup_streams** (Number) - Positive backup stream count. Default: `4`.
- **collect_acl** (Boolean) - Whether to back up ACLs. Default: `false`.
- **event_scan_enabled** (Boolean) - Whether to enable event scan. Omission retains the server setting.
- **event_scan_staging_path** (String) - Nonempty staging path required when enabling event scan. An explicit empty string can be supplied when disabling it.
- **job_start_time** (Number) - Seconds since midnight UTC, from `0` to `86399`. Omission retains the server value; `0` means midnight.
- **content** (Block List) - Ordered path selections. Each block must set exactly one nonempty `path`, `exclude_path`, or `include_path`. At least one `path` is required for path-based content.
- **tag_filter** (Block List) - Alternative to path-based content. Cannot be combined with `content`.

### Nested `content` Arguments

- **path** (String) - Bucket or path to protect.
- **exclude_path** (String) - Path to exclude from protection.
- **include_path** (String) - Exception path to include within an exclusion.

### Nested `tag_filter` Arguments

- **rule_item_type** (Number, Optional) - Rule item type. Only `2` is supported; default: `2`.
- **rules** (Block List, Required) - One or more rules, each containing required `name` and `value` strings.

## Attribute Reference

- **id** (String) - Commvault subclient ID.
- **subclient_id** (Number) - Commvault subclient ID.
- **plan_name** (String) - Plan name, when returned by Commvault.
- **backup_content_type** (Number) - `0` for path-based content or `1` for tag-based content.

## Usage Notes

Content updates replace the managed path, exclusion, and exception lists.
After a partially failed update, refresh and re-plan before retrying.
Event scan requires compatible Commvault multinode support, AWS staging
resources, and IAM permissions.

Destroy removes the content
group, not the AWS bucket. Applying does not run an immediate backup or verify
backup success. The parent registration's default subclient is not adopted
or modified by this resource.
Explicit deconfiguration is available for parent instances only.

## Import

S3 content groups can be imported using the subclient ID:

```shell
terraform import commvault_s3_object_storage_subclient.buckets 797
```
