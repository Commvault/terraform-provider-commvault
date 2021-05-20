---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: " Commvault : commvault_vm_group Resource "
subcategory: "VM Group"
description: |-
  Use the commvault_vm_group resource type to create and delete a VM group in the CommCell environment.
---

# commvault_vm_group (Resource)

  Use the commvault_vm_group resource type to create and delete a VM group in the CommCell environment.


## Syntax

```
resource "commvault_vm_group" "<local name>" 
{
	vm_group_name = "<vm group name>"
	plan_id = <plan ID>
	client_id = <client ID>
	vms = <List VMs> //Add VMs enclosed in quotes and separated by a comma
}
```

## Example Usage

```
resource "commvault_vm_group" "g1" 
{
	vm_group_name = "vg1"
	plan_id = "101"
	client_id = "266"
	vms = ["vm1","vm2"]
}
```

### Required

- **vm_group_name** (String) Specifies The name of the VM group.
- **client_id** (Number) Specifies The ID of the hypervisor client.
- **plan_id** (Number) Specifies The ID of the plan that you want to associate with the VM group.


### Optional

- **tags** (Set of String) Specifies The Tags that you want to back up in a VM group.
- **vms** (Set of String) Specifies The VMs that you want to back up in a VM group.
- **company_id** (Number) Specifies the company id to which the vm group should be associated with.
- **id** (String) The ID of this resource.


