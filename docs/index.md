---
page_title: "Provider: commvault"
subcategory: ""
description: |-
  The Commvault Terraform Provider interacts with the commvault REST API's for CRED Operations.
---

# Commvault Terraform Provider

With the Commvault Terraform provider, you can use Terraform to manage endpoints (called resources). Terraform is a configuration language for safely and efficiently managing infrastructure.

The Commvault Terraform module provides a set of named resource types, and specifies which arguments are allowed for each resource type. Using the resource types, you can create a configuration file, and apply changes to the Commvault REST APIs. For example, you can use the commvault_user resource type to add and delete users in your CommCell environment.

## Syntax
```
provider "commvault" {
	web_service_url = "URL of the commserver webservice/webconsole api endpoint"
	user_name = "username that is used to call APIs" 
	password = "password in base 64 encoded format"
    ignore_cert = "true/false to ignore certificate warnings for https endpoints"
}
```
## Example Usage

```
provider "commvault" {
	web_service_url = "https://webconsole.domain.com/webconsole/api"
	user_name = "admin" 
	password = "QnVebFRgciEoMg=="
```

### Required

- **web_service_url** (String) Specifies the Web Server URL of the commserver for performing Terraform Operations.

### Optional

- `password` (String) Specifies the Password for the user name to authentication to Web Server. Alternatively set CV_TER_PASSWORD environment variable for terraform to pick it.
- `user_name` (String) Specifies the User name used for authentication to Web Server.
- `ignore_cert` (Bool) true/false to ignore certificate warnings for https endpoints.



## Support Matrix
| SOFTWARE VERSION  | SUPPORTED RESOURCES |
| --------  | ------------------- | 
| 11.24 |  <ul><li>commvault_plan</li><li>commvault_user</li><li>commvault_vm_group</li><li>commvault_vmware_hypervisor</li><li>commvault_amazon_hypervisor</li><li>commvault_azure_hypervisor</li><li>commvault_plan_to_vm</li><li>commvault_company</li><li>commvault_disk_storage</li><li>commvault_aws_storage</li><li>commvault_azure_storage</li><li>commvault_google_storage</li><li>commvault_install_ma</li><li>commvault_security_association</li></ul> |
| 11.28      | <ul><li>commvault_user_v2</li><li>commvault_usergroup</li></ul> | 
| 11.30.28      | <ul><li>commvault_hypervisor_aws</li><li>commvault_hypervisor_azure</li><li>commvault_vmgroup_v2</li></ul> | 