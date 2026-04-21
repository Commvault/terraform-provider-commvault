---
page_title: "Commvault: commvault_oracle_instance Resource"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_instance resource to create, update, or delete Oracle database instances in the Commcell environment.
---

# commvault_oracle_instance (Resource)

Use the commvault_oracle_instance resource to manage Oracle database instances for backup and recovery operations. This resource allows you to configure Oracle instance properties including Oracle home, user credentials, and RMAN catalog settings.

## Syntax

```hcl
resource "commvault_oracle_instance" "<local_name>" {
  client_name         = "<Client Name>"
  instance_name       = "<Oracle SID>"
  oracle_home         = "<Oracle Home Path>"
  oracle_user         = "<Oracle OS User>"
  block_size          = <Block Size>
  use_catalog_connect = <true|false>
  catalog_connect     = "<Catalog Connect String>"
}
```

## Example Usage

### Basic Oracle Instance

```hcl
resource "commvault_oracle_instance" "prod_db" {
  client_name   = "oracle-server-01"
  instance_name = "PRODDB"
  oracle_home   = "/u01/app/oracle/product/19c/dbhome_1"
  oracle_user   = "oracle"
  block_size    = 65536
}
```

### Oracle Instance with RMAN Catalog

```hcl
resource "commvault_oracle_instance" "prod_db_with_catalog" {
  client_name         = "oracle-server-01"
  instance_name       = "PRODDB"
  oracle_home         = "/u01/app/oracle/product/19c/dbhome_1"
  oracle_user         = "oracle"
  block_size          = 131072
  use_catalog_connect = true
  catalog_connect     = "rman/rman@RMANCAT"
  tns_admin_path      = "/u01/app/oracle/product/19c/dbhome_1/network/admin"
}
```

## Argument Reference

### Required

- **client_name** (String) - Name of the client where the Oracle instance is configured. This cannot be changed after creation.
- **instance_name** (String) - Name of the Oracle instance (SID). This cannot be changed after creation.
- **oracle_home** (String) - Path to the Oracle home directory (e.g., `/u01/app/oracle/product/19c/dbhome_1`).

### Optional

- **oracle_user** (String) - Oracle OS user name (e.g., `oracle`).
- **sql_connect** (String, Sensitive) - SQL connect string for the Oracle instance.
- **tns_admin_path** (String) - Path to the TNS admin directory containing `tnsnames.ora`.
- **block_size** (Number) - Block size for RMAN backup operations. Default: `65536`.
- **use_catalog_connect** (Boolean) - Whether to use an RMAN recovery catalog. Default: `false`.
- **catalog_connect** (String) - RMAN catalog connect string (required if `use_catalog_connect` is `true`).
- **archive_log_dest** (String) - Archive log destination path.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The unique identifier of the Oracle instance in Commvault.

## Import

Oracle instances can be imported using the instance ID:

```shell
terraform import commvault_oracle_instance.prod_db <instance_id>
```
