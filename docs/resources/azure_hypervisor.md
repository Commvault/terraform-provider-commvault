---
page_title: "Commvault : commvault_azure_hypervisor Resource"
subcategory: "Virtualization"
description: |-
 Use the commvault_azure_hypervisor resource type to create or delete an Azure hypervisor in the CommCell environment.
---

# commvault_azure_hypervisor (Resource)


Use the commvault_azure_hypervisor resource type to create or delete an Azure hypervisor in the CommCell environment.


## Syntax

```
resource "commvault_azure_hypervisor" "<local name>"
{
	hypervisor_name = "<hypervisor name>"
	subscription_id = "<subscription Id>"
	tenant_id = "<tenant Id>"
	application_id= "<application Id>"
	application_password = "<application password>"
	access_nodes = "<access nodes>"
}
```

## Example Usage

```
resource "commvault_azure_hypervisor" "AzureHyp11"
{
	hypervisor_name = "AzureHypTest"
	subscription_id = "##########################"
	tenant_id = "#############################"
	application_id= "##############################"
	application_password = "################################"
	access_nodes = "AWSproxy"
}
```

### Required

- **hypervisor_name** (String) Specifies The name of the hypervisor.
- **subscription_id** (String) Specifies The subscription ID for your Azure account.
- **tenant_id** (String) Specifies The tenant ID for your Azure account.
- **application_id** (String) Specifies The application ID of the tenant.
- **application_password** (String) Specifies The password for the application ID of the tenant.
- **access_nodes** (String) Specifies The clients that have the VSA package installed and that act as proxy clients for Azure hypervisors.

### Optional

- **company_id** (Number) Specifies the company id to which the Azure Hypervisor should be associated with.
- **id** (String) The ID of this resource.


