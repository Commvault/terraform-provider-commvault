---
page_title: "Commvault: commvault_cloud_instance_deconfigure Resource"
subcategory: "Cloud Apps"
description: |-
  Use the commvault_cloud_instance_deconfigure resource to explicitly disable RDS, Aurora, or S3 protection.
---

# commvault_cloud_instance_deconfigure (Resource)

Use the commvault_cloud_instance_deconfigure resource to explicitly disable
protection for an RDS, Aurora snapshot, or S3 registration. This resource records
a **one-time operation**, rather than managing the registration itself.
Obtain approval and review retention before applying.

## Example Usage

```hcl
resource "commvault_cloud_instance_deconfigure" "approved" {
  instance_id         = var.approved_instance_id
  client_name         = var.expected_client_name
  instance_name       = var.expected_instance_name
  workload            = "AMAZON_RDS"
  confirm_deconfigure = true
}
```

## Argument Reference

### Required

- **instance_id** (Number) - Positive Cloud Apps instance ID, not the client ID.
- **client_name** (String) - Exact current Commvault client name.
- **instance_name** (String) - Exact current Commvault instance name.
- **workload** (String) - `AMAZON_RDS` for RDS or Aurora snapshot protection, or `AMAZON_S3`.
- **confirm_deconfigure** (Boolean) - Must be `true` to confirm that protection should be disabled.

Changing any argument requires replacement of the operation record.

## Attribute Reference

- **id** (String) - The instance ID associated with the operation.
- **completed** (Boolean) - `true` only after the instance is confirmed disabled.

## Usage Notes

- Refresh does not re-run the operation or monitor ongoing registration state.
- Destroy only forgets the receipt; it does **not** reactivate or delete the
  instance. There is no import or automatic reconfiguration.
- The instance ID, client name, instance name, and workload must match the
  existing registration. An already-disabled matching instance is not
  deconfigured again.
- If an operation fails or times out, check the actual registration status
  before retrying or removing the operation record from state.
- If another state owns the registration, coordinate changes and refresh it.
  A receipt does not replace that ownership or bypass its lifecycle policy.
- Normal RDS resource destroy still deletes the registration; normal S3
  resource destroy still deconfigures it. Delete dependent groups first when
  removing a registration.

This does not delete AWS databases/buckets or guarantee erasure of retained
backup data or Commvault metadata.
