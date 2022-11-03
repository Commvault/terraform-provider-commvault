---
page_title: "commvault_usergroup Resource"
subcategory: "Security"
description: |-
  
---

# commvault_usergroup (Resource)




## Schema

### Required

- `name` (String) To create an active directory usergroup, the domain name should be mentioned along with the usergroup name (domainName\usergroupName) and localUserGroup value must be given.

### Optional

- `allowmultiplecompanymembers` (Boolean) This property can be used to allow addition of users/groups from child companies. Only applicable for commcell and reseller company group.
- `associatedexternalgroups` (Block Set) (see [below for nested schema](#nestedblock--associatedexternalgroups))
- `azureguid` (String) Azure Object ID used to link this user group to Azure AD group and manage group membership of the user during SAML login
- `description` (String)
- `enabled` (Boolean) allows the enabling/disabling of the user group.
- `enabletwofactorauthentication` (String) Allows two-factor authentication to be enabled for the specific types of usergroups. it can be turned on or off based on user preferences. There will be usergroups that will not have this option.
- `enforcefsquota` (Boolean) Used to determine if a backup data limit will be set for the user group being created
- `laptopadmins` (Boolean) When set to true, users in this group cannot activate or be set as server owner
- `planoperationtype` (String) determines if an existing user has to be added to the user group or removed from the user group
- `quotalimitingb` (Number) if enforceFSQuota is set to true, the quota limit can be set in GBs
- `restrictconsoletypes` (Block List) (see [below for nested schema](#nestedblock--restrictconsoletypes))
- `users` (Block Set) (see [below for nested schema](#nestedblock--users))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--associatedexternalgroups"></a>
### Nested Schema for `associatedexternalgroups`

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--restrictconsoletypes"></a>
### Nested Schema for `restrictconsoletypes`

Optional:

- `consoletype` (Set of String)


<a id="nestedblock--users"></a>
### Nested Schema for `users`

Read-Only:

- `id` (Number) The ID of this resource.


