---
page_title: "commvault_s3_object_storage_subclient Data Source"
subcategory: "Amazon S3"
description: |-
  Use the commvault_s3_object_storage_subclient data source to retrieve an Amazon S3 content group.
---

# commvault_s3_object_storage_subclient (Data Source)

Reads an S3 content group by numeric subclient ID.

## Example Usage

```hcl
data "commvault_s3_object_storage_subclient" "buckets" {
  subclient_id = 797
}

output "s3_subclient_name" {
  value = data.commvault_s3_object_storage_subclient.buckets.subclient_name
}
```

## Argument Reference

### Required

- **subclient_id** (Number) - Positive Commvault subclient ID.

## Attribute Reference

- **id** (String) - Commvault subclient ID.
- **client_id** (Number) - Parent client ID.
- **instance_id** (Number) - Parent S3 instance ID.
- **subclient_name** (String) - Content-group name.
- **plan_id** (Number) - Assigned backup plan ID.
- **plan_name** (String) - Assigned backup plan name.
- **number_of_backup_streams** (Number) - Configured backup stream count.
- **collect_acl** (Boolean) - Whether ACLs are backed up.
- **job_start_time** (Number) - Job start time in seconds since midnight UTC.
- **content** (Block List) - Path selections containing `path`, `exclude_path`, or `include_path`.
- **tag_filter** (Block List) - Tag filters containing `rule_item_type` and `rules` with `name` and `value`.
- **backup_content_type** (Number) - `0` for path-based content or `1` for tag-based content.
- **event_scan_enabled** (Boolean) - Event-scan setting.
- **event_scan_staging_path** (String) - Event-scan staging path, when available.

Nested fields correspond to the
[resource](../resources/s3_object_storage_subclient.md).

## Usage Notes

A missing or non-S3 content group fails lookup.
