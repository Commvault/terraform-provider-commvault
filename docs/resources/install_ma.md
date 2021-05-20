---
page_title: " Commvault : commvault_install_ma Resource"
subcategory: "Install"
description: |-
    Use the commvault_install_ma resource type to Install or Uninstall a Media Agent in the Commcell environment.
---

# commvault_install_ma (Resource)

Use the commvault_install_ma resource type to Install or Uninstall a Media Agent in the Commcell environment.

## Syntax

```
resource "commvault_install_ma" "<local name>"{
	mediaagent_name = “<MA name>”
	hostname = “<machine host name>”
	user_name = “<user name of machine>”
	password = “<password>”  //Base64 encoded password
}
```

## Example Usage

```
resource "commvault_install_ma" "installma"{
	mediaagent_name = “MA name”
	hostname = “machine host name”
	user_name = “user name of machine”
	password = “password”  //Base64 encoded password
}

```

### Required

- **mediaagent_name** (String) Specifies the Media Agent name used for installation.
- **hostname** (String) Specifies the Media Agent Hostname user for the installation
- **user_name** (String) Specifies the User name of the host computer for the installation.
- **password** (String) Specifies the password for the host computer for the installation.

### Optional

- **company_id** (Number) Specifies the company id to which the installed MA should be associated with.
- **id** (String) The ID of this resource.


