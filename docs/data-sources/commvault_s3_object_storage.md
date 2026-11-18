---
page_title: "commvault_s3_object_storage Data Source"
subcategory: "Amazon S3"
description: |-
  Use the commvault_s3_object_storage data source to retrieve an Amazon S3 source registration.
---

# commvault_s3_object_storage (Data Source)

Reads an S3 source registration by numeric Cloud Apps instance ID.

## Example Usage

```hcl
data "commvault_s3_object_storage" "source" {
  instance_id = 123
}

output "s3_client_id" {
  value = data.commvault_s3_object_storage.source.client_id
}
```

## Argument Reference

### Required

- **instance_id** (Number) - Positive Cloud Apps instance ID, not the client ID.

## Attribute Reference

- **id** (String) - Cloud Apps instance ID.
- **client_id** (Number) - Enclosing Commvault client ID.
- **client_name** (String) - Registration client name.
- **instance_name** (String) - Registration instance name.
- **plan_name** (String) - Assigned backup plan name.
- **host_url** (String) - S3 endpoint.
- **credential_type** (String) - Authentication type.
- **credential_name** (String) - Associated credential name, when available.
- **member_servers** (List of String) - Access-node client names.
- **use_iam_role** (Boolean) - Whether IAM-role authentication is used.
- **number_of_backup_streams** (Number) - Configured backup stream count.
- **region_name** (String) - Hosted infrastructure region, when available.
- **validate_region_info** (Boolean) - Region-validation setting.
- **is_customer_managed_access_node** (Boolean) - Customer-managed access-node setting.
- **object_storage_indexing_schema_v2** (Boolean) - Indexing-schema-v2 setting, when returned.
- **object_storage_multinode** (Number) - Multinode setting, when returned.
- **is_cloud_encryption_key_set** (Boolean) - Encryption-key status flag.
- **comm_cell_id** (Number) - CommCell ID, when available.
- **use_resource_pool_info** (Boolean) - Resource-pool setting, when available.

The fields correspond to the [resource](../resources/s3_object_storage.md).
Deconfiguration-only `company_name` and `multi_commcell_id` are not exposed.

## Usage Notes

A missing or non-S3 instance fails lookup rather than returning an empty result.
