# Commvault Oracle Terraform Provider — Test Team Guide

This guide explains how to install the custom Terraform provider binary and test the Oracle resources.

---

## Prerequisites

- [Terraform](https://developer.hashicorp.com/terraform/downloads) **1.0 or later** installed
- Access to a Commvault CommServe environment
- A client machine with Oracle agent already installed (for instance/subclient tests)
- Windows OS (the binary is a `.exe`)

---

## Step 1: Install the Provider Binary

Terraform looks for custom providers in a specific local directory. You need to place the `terraform-provider-commvault.exe` file in the right location.

### Directory Structure to Create

On **Windows**, create this exact folder path:

```
%APPDATA%\terraform.d\plugins\registry.terraform.io\commvault\commvault\1.0.0\windows_amd64\
```

In PowerShell:

```powershell
# Create the plugin directory
$pluginDir = "$env:APPDATA\terraform.d\plugins\registry.terraform.io\commvault\commvault\1.0.0\windows_amd64"
New-Item -ItemType Directory -Force -Path $pluginDir

# Copy the provider binary into it
Copy-Item "c:\Users\maheshp\terraform-provider-commvault\terraform-provider-commvault.exe" "$pluginDir\terraform-provider-commvault_v1.0.0_x5.exe"
```

> **Why rename it?** Terraform requires the binary filename to follow the pattern `terraform-provider-{name}_v{version}_x5.exe` on Windows.

---

## Step 2: Create a Terraform Working Directory

Create a new folder for your Terraform configuration files:

```powershell
New-Item -ItemType Directory -Force -Path "C:\terraform-oracle-test"
cd C:\terraform-oracle-test
```

---

## Step 3: Create the Terraform Configuration Files

### File 1: `versions.tf` — Tell Terraform to use the local provider

```hcl
terraform {
  required_providers {
    commvault = {
      source  = "commvault/commvault"
      version = "1.0.0"
    }
  }
  required_version = ">= 1.0.0"
}
```

### File 2: `variables.tf` — Define all inputs (no hardcoded values)

```hcl
variable "web_service_url" {
  description = "Commvault CommServe API URL"
  type        = string
}

variable "user_name" {
  description = "Commvault admin username"
  type        = string
}

variable "password" {
  description = "Commvault admin password"
  type        = string
  sensitive   = true
}

variable "client_name" {
  description = "Name of the client with Oracle agent installed"
  type        = string
}

variable "instance_name" {
  description = "Oracle database SID (instance name)"
  type        = string
}

variable "oracle_home" {
  description = "Full path to Oracle home directory on the client"
  type        = string
}

variable "oracle_user" {
  description = "OS username that owns the Oracle installation"
  type        = string
}

variable "plan_id" {
  description = "Commvault Plan ID to associate with the instance"
  type        = number
  default     = 0
}
```

### File 3: `provider.tf` — Configure the provider

```hcl
provider "commvault" {
  web_service_url = var.web_service_url
  user_name       = var.user_name
  password        = var.password
  ignore_cert     = true   # Set to false if your CommServe has a valid SSL certificate
}
```

### File 4: `main.tf` — The Oracle resources

```hcl
# -----------------------------------------------------------------------
# Oracle Instance
# Creates an Oracle instance (SID) in Commvault for the given client
# -----------------------------------------------------------------------
resource "commvault_oracle_instance" "example" {
  client_name   = var.client_name
  instance_name = var.instance_name
  oracle_home   = var.oracle_home
  oracle_user   = var.oracle_user

  # Database connection - OS authentication (default '/' means no password needed)
  sql_connect_user   = "/"
  sql_connect_domain = var.instance_name   # TNS alias or SID

  # RMAN settings
  block_size         = 1048576   # 1 MB (default)
  cross_check_timeout = 600      # 10 minutes

  # Optional: associate a Commvault plan
  # plan_id = var.plan_id
}

# -----------------------------------------------------------------------
# Oracle Subclient
# Defines WHAT gets backed up and HOW for the instance above
# -----------------------------------------------------------------------
resource "commvault_oracle_subclient" "example" {
  client_name    = var.client_name
  instance_name  = var.instance_name
  subclient_name = "terraform_subclient"
  description    = "Created by Terraform"

  # Backup options
  enable_backup               = true
  backup_archive_log          = true
  backup_sp_file              = true
  backup_control_file         = true
  delete_archive_log_after_backup = false

  # This subclient depends on the instance being created first
  depends_on = [commvault_oracle_instance.example]
}

# -----------------------------------------------------------------------
# Data Source: Read back the instance we just created
# Useful for referencing instance_id in other resources
# -----------------------------------------------------------------------
data "commvault_oracle_instance" "example" {
  client_name   = var.client_name
  instance_name = var.instance_name

  depends_on = [commvault_oracle_instance.example]
}

# -----------------------------------------------------------------------
# Data Source: Read back the subclient we just created
# -----------------------------------------------------------------------
data "commvault_oracle_subclient" "example" {
  client_name    = var.client_name
  instance_name  = var.instance_name
  subclient_name = "terraform_subclient"

  depends_on = [commvault_oracle_subclient.example]
}
```

### File 5: `outputs.tf` — Display useful values after apply

```hcl
output "oracle_instance_id" {
  description = "The Commvault instance ID assigned to the Oracle instance"
  value       = commvault_oracle_instance.example.id
}

output "oracle_subclient_id" {
  description = "The Commvault subclient ID"
  value       = commvault_oracle_subclient.example.id
}

output "instance_oracle_home" {
  description = "Oracle home path read back from Commvault"
  value       = data.commvault_oracle_instance.example.oracle_home
}

output "instance_use_catalog" {
  description = "Whether RMAN catalog is configured"
  value       = data.commvault_oracle_instance.example.use_catalog_connect
}
```

### File 6: `terraform.tfvars` — Your actual values (DO NOT commit to git)

```hcl
web_service_url = "https://your-commserve.example.com/webconsole/api"
user_name       = "admin"
password        = "your-password"
client_name     = "hugo"
instance_name   = "ORCL"
oracle_home     = "/u01/app/oracle/product/19c/dbhome_1"
oracle_user     = "oracle"
plan_id         = 2
```

> ⚠️ **Security Note**: Add `terraform.tfvars` to `.gitignore`. Never commit credentials to source control.

---

## Step 4: Run Terraform

Open PowerShell in `C:\terraform-oracle-test` and run:

```powershell
# 1. Initialize — downloads/validates the provider
terraform init

# 2. Preview the changes — see what will be created
terraform plan

# 3. Apply — create the resources in Commvault
terraform apply

# Type 'yes' when prompted
```

### Expected Output After `terraform apply`

```
commvault_oracle_instance.example: Creating...
commvault_oracle_instance.example: Creation complete after 5s [id=4214]
commvault_oracle_subclient.example: Creating...
commvault_oracle_subclient.example: Creation complete after 3s [id=8731]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

Outputs:

oracle_instance_id  = "4214"
oracle_subclient_id = "8731"
instance_oracle_home = "/u01/app/oracle/product/19c/dbhome_1"
instance_use_catalog = false
```

---

## Step 5: Verify in Commvault Console

1. Log into the Commvault **Command Center** or **CommCell Console**
2. Navigate to **Protect → Databases → Oracle**
3. Find the client (`hugo`) — the instance `ORCL` and subclient `terraform_subclient` should appear

---

## Step 6: Clean Up

To delete the resources created by Terraform:

```powershell
terraform destroy
# Type 'yes' when prompted
```

---

## Advanced Scenarios

### Scenario: Instance with Stored Credential (instead of OS auth)

If you have a Commvault credential stored for Oracle DB authentication:

```hcl
resource "commvault_oracle_instance" "with_credential" {
  client_name   = var.client_name
  instance_name = var.instance_name
  oracle_home   = var.oracle_home
  oracle_user   = var.oracle_user

  # Use a stored credential instead of inline username/password
  # Find the credential ID in Commvault: Manage → Security → Credentials
  db_connect_credential_id = 42
}
```

### Scenario: Instance with RMAN Recovery Catalog

```hcl
resource "commvault_oracle_instance" "with_catalog" {
  client_name   = var.client_name
  instance_name = var.instance_name
  oracle_home   = var.oracle_home
  oracle_user   = var.oracle_user

  use_catalog_connect    = true
  catalog_connect_user   = "rman_user"
  catalog_connect_domain = "RMANCAT"    # TNS alias for the catalog DB
}
```

### Scenario: Install Oracle Agent on a New Server

> ⚠️ This will trigger an actual installation job on the target server.

```hcl
resource "commvault_oracle_install_agent" "new_server" {
  client_name         = "oracle-db-server-03"
  host_name           = "192.168.10.50"           # IP or hostname of the target
  commserve_host_name = "your-commserve.example.com"
  user_name           = "root"                    # SSH user for Linux
  password            = var.ssh_password          # SSH password (use variable!)
  install_os_type     = 2                         # 1 = Windows, 2 = Unix/Linux
  unix_group          = "oinstall"
  plan_id             = var.plan_id
}
```

### Scenario: Read Backup Pieces (Data Source)

```hcl
# Get backup pieces for the last 7 days
data "commvault_oracle_backup_pieces" "recent" {
  instance_id = commvault_oracle_instance.example.id
  from_time   = 1700000000   # Unix timestamp (epoch seconds)
  to_time     = 1700604800   # Unix timestamp (epoch seconds)
}

output "backup_pieces" {
  value = data.commvault_oracle_backup_pieces.recent.backup_pieces
}
```

### Scenario: Read RMAN Logs for a Job

```hcl
data "commvault_oracle_rman_logs" "job_logs" {
  job_id = "44725"   # Job ID from Commvault job history
}

output "rman_log_content" {
  value = data.commvault_oracle_rman_logs.job_logs.log_content
}
```

---

## Troubleshooting

### "This provider is not available in the registry"

You are missing the provider binary or it is in the wrong folder. Re-run Step 1.

### "Error: Failed to authenticate"

- Check `web_service_url` — should end with `/webconsole/api` (NOT `/commandcenter/api`)
- Check username and password
- Check that `ignore_cert = true` if the CommServe uses a self-signed certificate

### "Error: Oracle instance ORCL not found on client hugo"

- Verify the `client_name` exactly matches the client name in Commvault (case-sensitive)
- Verify the Oracle agent is installed and running on the client
- Verify the `instance_name` matches the Oracle SID

### "Error: operation [CreateOracleInstance] failed"

- Check that `oracle_home` is the correct path on the **client** machine (not local)
- Check that `oracle_user` is the OS user that owns the Oracle installation on the client

### Terraform `init` Fails

Make sure the binary filename exactly matches `terraform-provider-commvault_v1.0.0_x5.exe` — Terraform is strict about this naming pattern.

---

## All Available Resource Schema Fields

### `commvault_oracle_instance`

| Field | Required | Type | Description |
|---|---|---|---|
| `client_name` | ✅ Yes | string | Client name in Commvault |
| `instance_name` | ✅ Yes | string | Oracle SID |
| `oracle_home` | ✅ Yes | string | Path to Oracle home on the client |
| `oracle_user` | No | string | OS user that owns Oracle |
| `oracle_wallet_authentication` | No | bool | Use Oracle wallet auth (default: false) |
| `sql_connect_user` | No | string | DB username for SQL connect (default: "/") |
| `sql_connect_domain` | No | string | Oracle SID/TNS entry for SQL connect |
| `db_connect_credential_id` | No | int | Credential ID of type 'Oracle' (preferred) |
| `tns_admin_path` | No | string | Path to TNS admin directory |
| `block_size` | No | int | RMAN block size (default: 1048576) |
| `cross_check_timeout` | No | int | Crosscheck timeout seconds (default: 600) |
| `use_catalog_connect` | No | bool | Use RMAN recovery catalog |
| `catalog_connect_user` | No | string | Catalog DB username |
| `catalog_connect_domain` | No | string | Catalog DB TNS entry |
| `catalog_connect_credential_id` | No | int | Credential ID for catalog (preferred) |
| `os_user_credential_id` | No | int | Windows impersonation credential ID |
| `archive_log_dest` | No | string | Archive log destination path |
| `plan_id` | No | int | Commvault plan ID to associate |
| `client_id` | Computed | int | Client ID (set by Commvault) |

### `commvault_oracle_subclient`

| Field | Required | Type | Description |
|---|---|---|---|
| `subclient_name` | ✅ Yes | string | Subclient name |
| `client_name` | ✅ Yes | string | Client name |
| `instance_name` | ✅ Yes | string | Oracle instance name |
| `storage_policy` | No | block | Data backup storage policy `{ name, id }` |
| `log_storage_policy` | No | block | Log backup storage policy `{ name, id }` |
| `enable_backup` | No | bool | Enable backup (default: true) |
| `description` | No | string | Description |
| `backup_archive_log` | No | bool | Backup archive logs (default: true) |
| `backup_sp_file` | No | bool | Backup SP file (default: true) |
| `backup_control_file` | No | bool | Backup control file (default: true) |
| `delete_archive_log_after_backup` | No | bool | Delete archive logs after backup (default: false) |
| `instance_id` | Computed | int | Instance ID |
| `client_id` | Computed | int | Client ID |

### `commvault_oracle_install_agent`

| Field | Required | Type | Description |
|---|---|---|---|
| `client_name` | ✅ Yes | string | New client name in Commvault |
| `host_name` | ✅ Yes | string | IP or hostname of target server |
| `commserve_host_name` | ✅ Yes | string | CommServe hostname |
| `user_name` | ✅ Yes | string | SSH/OS username on target server |
| `password` | ✅ Yes | string | SSH/OS password (sensitive) |
| `install_os_type` | No | int | 1=Windows, 2=Unix/Linux (default: 1) |
| `plan_id` | No | int | Plan to associate |
| `unix_group` | No | string | Unix group (default: "oinstall") |
| `unix_group_access` | No | int | Group access bits (default: 7) |
| `unix_other_access` | No | int | Other access bits (default: 5) |
| `display_name` | No | string | Display name (defaults to client_name) |
| `force_reboot` | No | bool | Force reboot after install (default: false) |
| `task_id` | Computed | int | Installation task ID |
| `job_id` | Computed | string | Installation job ID |

---

## Quick Reference: Environment Variables Alternative

Instead of `terraform.tfvars`, you can set variables via environment variables (useful for CI/CD):

```powershell
$env:TF_VAR_web_service_url = "https://your-commserve.example.com/webconsole/api"
$env:TF_VAR_user_name       = "admin"
$env:TF_VAR_password        = "your-password"
$env:TF_VAR_client_name     = "hugo"
$env:TF_VAR_instance_name   = "ORCL"
$env:TF_VAR_oracle_home     = "/u01/app/oracle/product/19c/dbhome_1"
$env:TF_VAR_oracle_user     = "oracle"

terraform apply
```
