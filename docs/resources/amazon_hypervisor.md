---
page_title: "Commvault : commvault_amazon_hypervisor Resource"
subcategory: "Deprecated"
description: |-
 Use the commvault_amazon_hypervisor resource type to create or delete an Amazon hypervisor in the CommCell environment.
---

# commvault_amazon_hypervisor (Resource)


Use the commvault_amazon_hypervisor resource type to create or delete an Amazon hypervisor in the CommCell environment. This is deprecated in the latest version. Starting SP32, use commvault_hypervisor_aws resource.



## Syntax

```
resource "commvault_amazon_hypervisor" "<local name>"
{
	client_name = "<hypervisor name>"
	use_iam_role = <Boolean values: true or false>
	access_key = <access key>
	secret_key = "<secret key>"
	access_nodes = "<access nodes>"
}
```
## Example Usage

```
resource "commvault_amazon_hypervisor" "Amazonhyp1"
{
	client_name = "Amazon01"
	use_iam_role = false
	access_key = "##########"
	secret_key = "#############################"
	access_nodes = "AWS_proxy"
}
```
### Required

- **access_nodes** (String) Specifies The clients that have the VSA package installed and that act as proxy clients for Amazon hypervisors.
- **client_name** (String) Specifies The name of the Amazon hypervisor.
- **use_iam_role** (Boolean) Specifies whether you want to use IAM role.

### Optional

- **access_key** (String) Specifies The access key ID for your Amazon account.
- **company_id** (Number) Specifies the company id to which the Amazon Hypervisor should be associated with.
- **id** (String) The ID of this resource.
- **regions** (String) Specifies the regions used for the Hypervisor
- **secret_key** (String) Specifies The secret key ID for your Amazon account.


