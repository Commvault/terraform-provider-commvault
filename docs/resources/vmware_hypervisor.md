---
page_title: " Commvault : commvault_vmware_hypervisor Resource"
subcategory: "Virtualization"
description: |-
  Use the commvault_vmware_hypervisor resource type to create and delete a VMware hypervisor in the CommCell environment.

---

# commvault_vmware_hypervisor (Resource)

Use the commvault_vmware_hypervisor resource type to create and delete a VMware hypervisor in the CommCell environment.

## Syntax

```
resource "commvault_vmware_hypervisor" "<local name>"
{
	display_name = "<Display name>"
	host_name = "<Hypervisor hostname"
	user_name = "<User name>"
	password = "<Password>"
	access_nodes = "<Access nodes>"
}
```

## Example Usage

```
resource "commvault_vmware_hypervisor" "hyp1"
{
	display_name = "hyp-disp-name"
	host_name = "cvlt.hyphostname.com"
	user_name = "johndoe"
	password = "commvault"
	access_nodes = "cvlt-idc-loc"
}
```

### Required

- **display_name** (String) Specifies The display name of the hypervisor.
- **host_name** (String) Specifies The host name of the hypervisor.
- **user_name** (String) Specifies The user name for the account.
- **password** (String) Specifies The password for the account.
- **access_nodes** (String) Specifies The clients that have the VSA package installed and that act as proxy clients for hypervisors.

### Optional

- **company_id** (Number) Specifies the company id to which the Vmware Hypervisor should be associated with.
- **id** (String) The ID of this resource.


