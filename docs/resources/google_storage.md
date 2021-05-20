---
page_title: " Commvault : commvault_google_storage Resource"
subcategory: "Storage"
description: |-
  Use the commvault_google_storage resource type to create or delete a Google Cloud Storage in the Commcell environment.
---

# commvault_google_storage (Resource)

Use the commvault_google_storage resource type to create or delete a Google Cloud Storage in the Commcell environment.


## Syntax

```
resource "commvault_google_storage" "<local name>"{
	storage_name = “<Storage Name>”
	mediaagent = “<MediaAgent Name>”
	service_host = “<Googleservice host>”
	secret_access_key = “<Google access key>”
	access_key_id = “<Google access key password>”  //Base64 encoded password
	credentials_name = “<Google credentials Name>”
	container = “<Google container name>”
	ddb_location = “<DDB Location>”
	company_id = <Company ID>
}

```

## Example Usage

```
resource "commvault_google_storage" "CAWS1"{
	storage_name = “StorageName”
	mediaagent = “MediaAgent1”
	service_host = “service host”
	secret_access_key = “access key”
	access_key_id = “access key password”  //Base64 encoded password
	credentials_name = “azure_credentials”
	container = “container name”
	ddb_location = “c:\\Location”
	company_id = 22
}

```

### Required

- **storage_name** (String) Specifies the Name of the Google Storage.
- **mediaagent** (String) Specifies the Media agent used for the Google Storage.
- **service_host** (String) Specifies the service host name for the Google storage.
- **bucket** (String) Specifies the bucket name user for the Google Storage.

### Optional

- **credentials_name** (String) Specifies the saved creation name for creating Google Storage.
- **access_key_id** (String) Specifies the access key id for the Google Storage.
- **secret_access_key** (String) Specifies the secret access key for Google Storage.
- **ddb_location** (String) Specifies the Deduplication path for the Google Storage
- **company_id** (Number) Specifies the company id to which the created Google storage should be associated with.
- **id** (String) The ID of this resource.
