---
page_title: "commvault_user_v2 Resource"
subcategory: "Security"
description: |-
  
---

# commvault_user_v2 (Resource)




## Schema

### Required

- `email` (String) Used to provide an email-id to the new user. This email-id is used for logging in the user. Please note that email ids are compulsory for company and local users and optional for external users.

### Optional

- `company` (Block List) (see [below for nested schema](#nestedblock--company))
- `enabled` (Boolean) enable or disable the user.
- `fullname` (String) Used to provide a name to the new user.
- `inviteuser` (Boolean) User will receive an email to install backup software package on their device if this is set to true.
- `name` (String) Used to provide the new user with a username. This username can be used for logging in the user instead of email-id when duplicate email-ids are present. For external user, it is necessary to provide the domain name along with the username (domainName\username). To create a company user, the company id or name needs to be provided in the company entity.
- `password` (String, Sensitive) Used to provide a password to the user being created. This will be accepted when the useSystemGeneratePassword tag is false. The password has to be provided in Base64 format.
- `plan` (Block List) (see [below for nested schema](#nestedblock--plan))
- `userprincipalname` (String) Change User Principal Name(UPN) for existing user. This User Principal Name can be used for logging-in.
- `usesystemgeneratepassword` (Boolean) Choose to provide a system generated password to the user instead of providing your own password. An email will be sent to the user to reset the password. If it is set to true, password tag need not be provided. If it is set to false, password needs to be provided in the password tag in Base64 format.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--company"></a>
### Nested Schema for `company`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--plan"></a>
### Nested Schema for `plan`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


