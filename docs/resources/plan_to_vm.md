---
page_title: " Commvault : commvault_plan_to_vm Resource"
subcategory: "Plan to VM Association"
description: |-
    Use the commvault_plan_to_vm resource type to associate plan to a VM in the Commcell environment.
---

# commvault_plan_to_vm (Resource)

Use the commvault_plan_to_vm resource type to associate plan to a VM in the Commcell environment.

## Syntax

```
resource "commvault_plan_to_vm" "<local name>"{
	plan = "<Plan Name>"
	vm_name = "<VM Name>"
}
```

## Example Usage

```
resource "commvault_plan_to_vm" "assoc1"{
	plan = "Plan1"
	vm_name = "SP20_CCM_bkpvm1"
}
```

### Required

- **plan** (String) Specifies the plan name to associate.
- **vm_name** (String) Specifies the vm name to associate.

### Optional

- **id** (String) The ID of this resource.
- **new_plan** (String) Specifies the new plan name for association


