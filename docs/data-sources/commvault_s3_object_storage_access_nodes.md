---
page_title: "commvault_s3_object_storage_access_nodes Data Source"
subcategory: "Amazon S3"
description: |-
  Use the commvault_s3_object_storage_access_nodes data source to list eligible Amazon S3 access nodes.
---

# commvault_s3_object_storage_access_nodes (Data Source)

Lists existing access nodes eligible for Amazon S3 source protection.

## Example Usage

```hcl
data "commvault_s3_object_storage_access_nodes" "available" {
  exclude_client_groups = true
  exclude_clients       = false
}

output "s3_access_node_names" {
  value = [
    for node in data.commvault_s3_object_storage_access_nodes.available.access_nodes :
    node.client_name
  ]
}
```

## Argument Reference

### Optional

- **region_id** (Number) - Positive region ID. Omit to list nodes without a region filter.
- **exclude_client_groups** (Boolean) - Whether to exclude client groups. Default: `false`.
- **exclude_clients** (Boolean) - Whether to exclude individual clients. Default: `false`.

## Attribute Reference

- **id** (String) - Identifier for the lookup and its filters.
- **access_nodes** (List of Object) - Matching access nodes. Each object contains:
  - **client_id** (Number) - Client or client-group ID.
  - **client_name** (String) - Client or client-group name.
  - **entity_type** (Number) - Commvault entity type.

An empty `access_nodes` list means no eligible nodes matched the filters.
