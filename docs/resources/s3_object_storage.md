---
page_title: "Commvault: commvault_s3_object_storage Resource"
subcategory: "Amazon S3"
description: |-
  Use the commvault_s3_object_storage resource to manage Amazon S3 source registrations for protection.
---

# commvault_s3_object_storage (Resource)

Registers an Amazon S3 **source workload** for protection by Commvault. This is
not `commvault_storage_cloud_s3`, which configures a backup storage destination.
This resource is not included in release 1.2.18.

## Example Usage

```hcl
resource "commvault_s3_object_storage" "source" {
  client_name     = "s3-production"
  instance_name   = "s3-production"
  plan_name       = "S3 protection"
  credential_name = "existing-aws-credential"
  member_servers  = ["existing-backup-gateway"]
}
```

## Argument Reference

### Required

- **client_name** (String) - Registration client name. Changing this requires replacement.
- **instance_name** (String) - Registration instance name. Supports in-place rename.
- **plan_name** (String) - Name of an existing compatible backup plan.

### Optional

- **credential_type** (String) - `AMAZON_S3` (default), `AMAZON_IAM_ROLE`, or `AMAZON_STS_IAM_ROLE`.
- **credential_name** (String) - Existing Commvault credential. Required for key-based and STS authentication; omit for `AMAZON_IAM_ROLE`. Inline AWS secrets are not accepted.
- **member_servers** (List of String) - Ordered access-node client names. Required unless `region_name` is supplied for hosted infrastructure; always required for `AMAZON_IAM_ROLE`.
- **host_url** (String) - S3 endpoint. Default: `s3.amazonaws.com`.
- **number_of_backup_streams** (Number) - Positive backup stream count. Default: `4`.
- **region_name** (String) - Hosted infrastructure region. Set `host_url` to its regional S3 endpoint.
- **validate_region_info** (Boolean) - Whether to validate region information. Default: `false`.
- **is_customer_managed_access_node** (Boolean) - Whether access nodes are customer-managed. Omission retains the server setting.
- **object_storage_indexing_schema_v2** (Boolean) - Indexing-schema-v2 setting. Omission retains the server setting. Some deployments do not return this value during refresh.
- **object_storage_multinode** (Number) - Nonnegative multinode setting supported by your deployment. Omission retains the server setting.
- **is_cloud_encryption_key_set** (Boolean) - Encryption-key status flag. Default: `false`. Does not provision an encryption key.
- **comm_cell_id** (Number) - CommCell ID. Default: `2`. Changing this requires replacement.
- **use_resource_pool_info** (Boolean) - Whether to use resource-pool information. Default: `false`. Changing this requires replacement.
- **company_name** (String) - Company used for deconfiguration. Default: `Commcell`. Set correctly for tenant-owned registrations.
- **multi_commcell_id** (Number) - Deconfiguration metadata. Default: `0`. Changing this or `company_name` only updates Terraform state.

## Attribute Reference

- **id** (String) - Cloud Apps instance ID, not the client ID.
- **instance_id** (Number) - Cloud Apps instance ID.
- **client_id** (Number) - Enclosing Commvault client ID.
- **use_iam_role** (Boolean) - Whether the instance uses an IAM role; `true` for `AMAZON_IAM_ROLE`.

## Usage Notes

Destroy deconfigures protection;
it does not delete AWS buckets, IAM roles, or credentials, and does not promise
to remove the enclosing Commvault client or retained backup data.
Destroy dependent subclients first. No backup/restore jobs are submitted.
For an explicit disable operation without destroying the owning resource, use
[commvault_cloud_instance_deconfigure](cloud_instance_deconfigure.md).

Access-node, indexing, and multinode settings require compatible infrastructure
and permissions in your Commvault deployment.

If creation succeeds but the API cannot resolve its instance ID, the provider
returns an explicit error. Locate the created instance and import it rather
than blindly reapplying and creating a duplicate client.

## Import

```shell
terraform import commvault_s3_object_storage.source 123
```

The import ID is a numeric **instance ID**. Include correct company/deconfiguration
metadata in configuration because the GET endpoint does not return it.
