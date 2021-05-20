---
page_title: " Commvault : commvault_user Resource "
subcategory: "Security"
description: |-
  Use the commvault_user resource type to create or delete an User in the CommCell environment.
---

# commvault_user (Resource)

Use the commvault_user resource type to create or delete an User in the CommCell environment.

## Syntax

```
resource "commvault_user" "<local name>" 
{
	user_name = "<user name>"
	password = "<password in plain English text>"
	description = "<sample description>"
	full_name = "<full name>"
	email = "<email address>"
}
```

## Example Usage

```
resource "commvault_user" "user1" 
{
	user_name = "jdoe"
	password = "commvault"
	description = "This is a test user"
	full_name = "John Doe"
	email = "jdoe@cmvt.com"
}
```


### Required

- **user_name** (String) Specifies the user name for the account.
- **password** (String) Specifies the password for the account.
- **email** (String) Specifies The email address of the user.

### Optional

- **description** (String) Specifies The description of the user account.
- **full_name** (String) Specifies The first, middle, and last names of the user.
- **company_id** (Number) Specifies the company id for which the created user will be associated with.
- **id** (String) The ID of this resource.

