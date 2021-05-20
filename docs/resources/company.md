---
page_title: " Commvault : commvault_company Resource"
subcategory: "Company"
description: |-
  Use the commvault_company resource type to create or delete a Company in the Commcell environment.
---

# commvault_company (Resource)

Use the commvault_company resource type to create or delete a Company in the Commcell environment.


## Syntax

```
resource "commvault_company" "<local name>"{
	company_name = "<Company Name>"
	email = "<Email ID>"
	contact_name = "<Contact Name>"
	company_alias = "<Company Alias>"
	plans =  [“<Plan name1>”,”<Plan name2>”]
	associated_smtp = “<SMTP Server>”
	send_email = <Boolean values: true or false>
}

```

## Example Usage

```
resource "commvault_company" "Company1"{
	company_name = "CompanyName"
	email = "DemoCompany@company.com"
	contact_name = "ContactName"
	company_alias = "CompanyAlias"
	plans =  [“Plan1”,”Plan2”]
	associated_smtp = “SMTP_Server”
	send_email = false
}

```

### Required

- **company_name** (String) Specifies the name of the Company.
- **email** (String) Specifies Email address for the tenant administrator.
- **contact_name** (String) Specifies Name of the tenant administrator.
- **company_alias** (String) Specifies the Alias name for the company.

### Optional

- **associated_smtp** (String) Specifies the SMTP address of the company.
- **send_email** (Boolean) Specifies whether email needs to be sent ot not
- **plans** (Set of String) Specifies the data protection plans to use for the company. The plans you select are the plans that the tenant administrator can choose from.
- **company_id** (Number) Specifies the company id to which the child company should be associated with.
- **id** (String) The ID of this resource.
