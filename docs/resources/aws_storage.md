---
page_title: " Commvault : commvault_aws_storage Resource"
subcategory: "Storage"
description: |-
  Use the commvault_aws_storage resource type to create or delete a AWS Cloud Storage in the Commcell environment.
---

# commvault_aws_storage (Resource)

  Use the commvault_aws_storage resource type to create or delete a AWS Cloud Storage in the Commcell environment.

##Syntax
```
resource "commvault_aws_storage" "<local name>"{
	storage_name = "<storage name>"
	mediaagent = "<Media agent name>"
	service_host = "<AWS Service host>"
	access_key_id = "<AWS Access Key>"
	secret_access_key = "<AWS Secret Access Key>"  //Base64 encoded password
	bucket = "<AWS Bucket name>"
	ddb_location = "<DDB Location>"
	company_id = <company ID>
}
```

## Example Usage
```
resource "commvault_aws_storage" "CAWS"{
	storage_name = "DemoAwsCloudStorage"
	mediaagent = "MediaAgent"
	service_host = "s3.us-east-1.amazonaws.com"
	access_key_id = "AccessKey_ID"
	secret_access_key = "Secret_Access_Key"  //Base64 encoded password
	bucket = "Bucket12"
	ddb_location = "c:\\DemoDiskStorage-Dedup-DDB"
	company_id = 22
}
```

### Required

- **storage_name** (String) Specifies the Name of the AWS Storage.
- **mediaagent** (String) Specifies the Media agent used for the AWS Storage.
- **service_host** (String) Specifies the service host name for the AWS storage.
- **bucket** (String) Specifies the bucket name used for AWS Storage.

### Optional

- **credentials_name** (String) Sepcifies the saved creation name for creating AWS Storage.
- **access_key_id** (String) Specifies the access key id for the AWS Storage.
- **secret_access_key** (String) Specifies the secret access key for the AWS Storage.
- **ddb_location** (String) Specifies the Deduplication path for the AWS Storage.
- **company_id** (Number) Specifies the company id to which the created AWS storage should be associated with.
- **id** (String) The ID of this resource.



