---
page_title: "commvault_oracle_instance Data Source"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_instance data source to retrieve information about Oracle instances.
---

# commvault_oracle_instance (Data Source)

Use the commvault_oracle_instance data source to retrieve information about Oracle database instances configured in Commvault. This can be used to look up instance IDs and properties for use in other resources.

## Example Usage

### Lookup by Client and Instance Name

```hcl
data "commvault_oracle_instance" "prod" {
  client_name   = "oracle-server-01"
  instance_name = "PRODDB"
}

output "instance_id" {
  value = data.commvault_oracle_instance.prod.id
}
```

### Use in Other Resources

```hcl
data "commvault_oracle_instance" "prod" {
  client_name   = "oracle-server-01"
  instance_name = "PRODDB"
}

resource "commvault_oracle_subclient" "backup" {
  client_name   = data.commvault_oracle_instance.prod.client_name
  instance_name = data.commvault_oracle_instance.prod.instance_name
  subclient_name = "custom_backup"
  enable_backup  = true
}
```

## Argument Reference

### Required

- **client_name** (String) - Name of the client where the Oracle instance is configured.
- **instance_name** (String) - Name of the Oracle instance (SID).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The unique identifier of the Oracle instance.
- **client_id** (Number) - The ID of the client.
- **instance_id** (Number) - The ID of the instance.
- **oracle_home** (String) - Path to the Oracle home directory.
- **oracle_user** (String) - Oracle OS user name.
- **block_size** (Number) - Block size configured for the instance.
- **use_catalog_connect** (Boolean) - Whether RMAN catalog is enabled.
- **catalog_connect** (String) - RMAN catalog connect string.
- **tns_admin_path** (String) - Path to the TNS admin directory.
- **archive_log_dest** (String) - Archive log destination path.
