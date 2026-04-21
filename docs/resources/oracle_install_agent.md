# commvault_oracle_install_agent

Installs the Commvault Oracle iDataAgent on a database server. This is a one-time operation that deploys the Oracle backup agent to enable backup and recovery operations.

## Example Usage

### Basic Installation on Linux Server

```hcl
resource "commvault_oracle_install_agent" "example" {
  client_name         = "oracle-db-server"
  host_name           = "192.168.1.100"
  commserve_host_name = "commserve.example.com"
  user_name           = "root"
  password            = var.ssh_password
  install_os_type     = 2  # Unix/Linux
  unix_group          = "oinstall"
}
```

### Installation with Plan Association

```hcl
resource "commvault_oracle_install_agent" "with_plan" {
  client_name         = "oracle-prod-server"
  host_name           = "oracle-prod.example.com"
  display_name        = "Oracle Production Server"
  commserve_host_name = "commserve.example.com"
  user_name           = "root"
  password            = var.ssh_password
  install_os_type     = 2
  plan_id             = data.commvault_plan.oracle_plan.id
  
  unix_group        = "oinstall"
  unix_group_access = 7
  unix_other_access = 5
  
  override_unix_group_permissions = true
  add_to_firewall_exclusion       = true
}
```

### Windows Installation

```hcl
resource "commvault_oracle_install_agent" "windows" {
  client_name         = "oracle-win-server"
  host_name           = "192.168.1.101"
  commserve_host_name = "commserve.example.com"
  user_name           = "Administrator"
  password            = var.windows_password
  install_os_type     = 1  # Windows
  
  disable_os_firewall        = false
  add_to_firewall_exclusion  = true
  override_client_info       = true
}
```

## Argument Reference

The following arguments are supported:

### Required

* `client_name` - (Required, ForceNew) Name of the client where Oracle agent will be installed. This becomes the client name in Commvault.
* `host_name` - (Required, ForceNew) Hostname or IP address of the server where Oracle agent will be installed.
* `commserve_host_name` - (Required, ForceNew) CommServe hostname that the agent will connect to.
* `user_name` - (Required, ForceNew) Username for client authentication. For Unix/Linux, this is the SSH user (typically root). For Windows, this is a local administrator account.
* `password` - (Required, Sensitive, ForceNew) Password for client authentication.

### Optional

* `display_name` - (Optional, ForceNew) Display name for the client in Commvault. Defaults to `client_name` if not specified.
* `install_os_type` - (Optional, ForceNew) OS type of the target server. Valid values:
  - `1` - Windows (default)
  - `2` - Unix/Linux
* `plan_id` - (Optional, ForceNew) Plan ID to associate with the installed agent for automatic protection.
* `unix_group` - (Optional, ForceNew) Unix group for Oracle installation. Only applicable for Unix/Linux. Default: `"oinstall"`.
* `unix_group_access` - (Optional, ForceNew) Unix group access permissions. Only applicable for Unix/Linux. Default: `7`.
* `unix_other_access` - (Optional, ForceNew) Unix other access permissions. Only applicable for Unix/Linux. Default: `5`.
* `override_unix_group_permissions` - (Optional, ForceNew) Whether to override Unix group and permissions. Default: `true`.
* `allow_multiple_instances` - (Optional, ForceNew) Allow multiple Commvault instances on the same client. Default: `false`.
* `disable_os_firewall` - (Optional, ForceNew) Disable OS firewall during installation. Default: `false`.
* `add_to_firewall_exclusion` - (Optional, ForceNew) Add Commvault to firewall exclusion list. Default: `true`.
* `force_reboot` - (Optional, ForceNew) Force reboot after installation. Default: `false`.
* `stop_oracle_services` - (Optional, ForceNew) Stop Oracle services during installation. Default: `false`.
* `override_client_info` - (Optional, ForceNew) Override client information if client already exists. Default: `true`.
* `enable_firewall_config` - (Optional, ForceNew) Enable firewall configuration. Default: `false`.
* `firewall_port` - (Optional, ForceNew) Firewall port number. Default: `0`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The task ID of the installation job.
* `task_id` - The task ID of the installation job.
* `job_id` - The job ID of the installation job (if available).

## Import

This resource does not support import as it represents a one-time installation operation.

## Notes

1. **One-time Operation**: This resource represents a one-time installation. The agent cannot be "deleted" through this resource - it must be uninstalled separately through Commvault.

2. **Prerequisites**: 
   - The target server must be reachable from the CommServe
   - For Unix/Linux: SSH access with the provided credentials
   - For Windows: WMI/Remote PowerShell access with the provided credentials
   - Required ports must be open between the CommServe and the target server

3. **Post-Installation**: After the agent is installed, use `commvault_oracle_instance` to configure Oracle instances for backup.
