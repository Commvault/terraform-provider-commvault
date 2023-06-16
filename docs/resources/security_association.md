---
page_title: " Commvault : commvault_security_association Resource "
subcategory: "Deprecated"
description: |-
    Use the commvault_security_association resource type to Security Associations in the Commcell environment.  
---

# commvault_security_association (Resource)

Use the commvault_security_association resource type to Security Associations in the Commcell environment. This is deprecated in the latest version. Starting SP32, use commvault_security_association_v2 resource.


## Syntax

```
resource "commvault_security_association" "<local name>"{
	client_list = ["<Client Name>"]
	user_group_name = "<User group name>"
	permissions_list = ["<Permission name>"]
}
```

## Example Usage

```
resource "commvault_security_association" "asso"{
	client_list = ["ClientName"]
	user_group_name = "CompanyAlias\\Tenant Admin"
	permissions_list = ["Agent Management"]
}
```

### Required

- **client_list** (Set of String) Specifies the list of clients for association.
- **user_group_name** (String) Specifies the user group name used for association.
- **permissions_list** (Set of String) Specifies the permission names list used for the association.

### Optional

- **id** (String) The ID of this resource.
