---
page_title: " Commvault : commvault_azure_storage (deprecated) Resource"
subcategory: "Deprecated"
description: |-
  Use the commvault_azure_storage resource type to create or delete a Azure Cloud Storage in the Commcell environment.
---

# commvault_azure_storage (Resource)

Use the commvault_azure_storage resource type to create or delete a Azure Cloud Storage in the Commcell environment. This is deprecated in the latest version. Starting SP32, use commvault_storage_cloud_azure resource.


## Syntax

```
resource "commvault_azure_storage" "<local name>"{
	storage_name = “<Storage name>”
	mediaagent = “<Media Agent name>”
	service_host = “<Azure service host>”
	account_name = “<Azure account name>”
	access_key_id = “<Azure access key password>”  //Base64 encoded password
	credentials_name = “<azure credentials>”
	container = “<container name>”
	ddb_location = “<Location>”
	company_id = <company ID>
}
```

## Example Usage

```
resource "commvault_azure_storage" "CAWS1"{
	storage_name = “StorageName”
	mediaagent = “MediaAgent1”
	service_host = “service host”
	account_name = “account name”
	access_key_id = “access key password”  //Base64 encoded password
	credentials_name = “azure_credentials”
	container = “container name”
	ddb_location = “c:\\Location”
	company_id = 22
}

```

### Required

- **storage_name** (String) Specifies the Name of the Azure Storage.
- **service_host** (String) Specifies the service host name for the Azure storage.
- **mediaagent** (String) Specifies the Media agent used for the Azure Storage.
- **container** (String) Specifies the container name user for the Azure Storage.

### Optional

- **credentials_name** (String) Specifies the saved creation name for creating Azure Storage.
- **account_name** (String) Specifies the Account name for the Azure Storage.
- **access_key_id** (String) Specifies the access key id for the Azure Storage.
- **ddb_location** (String) Specifies the Deduplication path for the Azure Storage
- **company_id** (Number) Specifies the company id to which the created Azure storage should be associated with.
- **id** (String) The ID of this resource.
