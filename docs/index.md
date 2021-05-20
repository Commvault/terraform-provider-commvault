---
page_title: "Provider: commvault"
subcategory: ""
description: |-
  The Commvault Terraform Provider interacts with the commvault REST API's for CRED Operations.
---

# Commvault Terraform Provider

With the Commvault Terraform module, you can use Terraform to manage endpoints (called resources). Terraform is a configuration language for safely and efficiently managing infrastructure.

The Commvault Terraform module provides a set of named resource types, and specifies which arguments are allowed for each resource type. Using the resource types, you can create a configuration file, and apply changes to the Commvault REST APIs. For example, you can use the commvault_user resource type to add and delete users in your CommCell environment. You use the GO programming language to execute the APIs.

## Syntax
```
provider "commvault" {
	web_service_url = "URL of the commserver webservice"
	user_name = "username that is used to call APIs" 
	password = "password in base 64 encoded format"
}
```
## Exampale Usage

```
provider "commvault" {
	web_service_url = "http://CommCellBkp.domain.com:81/SearchSvc/CVWebService.svc/"
	user_name = "admin" 
	password = "QnVebFRgciEoMg=="
```

### Required

- **password** (String) Specifies the Password for the user name to authentication to Web Server.
- **user_name** (String) Specifies the User name used for authentication to Web Server
- **web_service_url** (String) Specifies the Web Server URL of the commserver for performing Terraform Operations.
