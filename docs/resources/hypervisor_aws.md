# commvault_hypervisor_aws (Resource)

Use the commvault_hypervisor_aws resource type to create or delete an Amazon hypervisor in the CommCell environment.

## Example Usage

**Configure Amazon Web Services Hypervisor using IAM role**
```hcl
resource "commvault_hypervisor_aws" "aws-test-IAM-role" {
  name = "Terraform-Test-AWS-Hypervisor"
  accessnodes {
    id   = data.commvault_client.accessnode1.id
    type = 3
  }
  accessnodes {
    id   = data.commvault_client.accessnode2.id
    type = 3
  }
  accessnodes {
    id   = data.commvault_clientgroup.accessnode3.id
    type = 28
  }
  useiamrole = "true"
}
```

**Configure Amazon Web Services Hypervisor using Access and secret key**
```hcl
resource "commvault_hypervisor_aws" "aws-test-accessandsecret" {
  name = "Terraform-Test-AWS-Access&Secret"
  accessnodes {
    id   = data.commvault_client.accessnode1.id
    type = 3
  }
  accessnodes {
    id   = data.commvault_clientgroup.accessnode3.id
    type = 28
  }
  credentials {
    id = data.commvault_credential.cred.id
  }
```
**Configure Amazon Web Services Hypervisor with custom settings using Access and secret key**
```hcl
resource "commvault_hypervisor_aws" "aws-test-accessandsecret" {
  name = "Terraform-Test-AWS-Access&Secret"
  accessnodes {
    id   = data.commvault_client.accessnode1.id
    type = 3
  }
  accessnodes {
    id   = data.commvault_clientgroup.accessnode3.id
    type = 28
  }
  credentials {
    id = data.commvault_credential.cred.id
  }
  fbrunixmediaagent {
    id = data.commvault_client.frel.id
  }
  activitycontrol {
    enablebackup  = "false"
    enablerestore = "false"
  }
  displayname           = "Terraform-Test-AWS-Access&Secret-Modified"
  region                = "us-east-1"
  enableawsadminaccount = "false"
  settings {
    guestcredentials {
      password = "###############"
      name     = "test-guest-creds"
    }
    timezone {
      name = "Greenwich Standard Time"
      id   = 39
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the hypervisor group being created
- `credentials` (Block List) (see [below for nested schema](#nestedblock--credentials))
- `useiamrole` (String) if Iam Role is used

### Optional

- `accessnodes` (Block Set) (see [below for nested schema](#nestedblock--accessnodes))
- `activitycontrol` (Block List) (see [below for nested schema](#nestedblock--activitycontrol))
- `displayname` (String) The name of the hypervisor that has to be changed
- `enableawsadminaccount` (String)
- `fbrunixmediaagent` (Block List) (see [below for nested schema](#nestedblock--fbrunixmediaagent))
- `hypervisortype` (String) [Amazon]
- `region` (String) AWS region if Iam role is used
- `rolearn` (String) Role ARN for STS assume role with IAM policy
- `settings` (Block List) (see [below for nested schema](#nestedblock--settings))
- `skipcredentialvalidation` (String) if credential validation has to be skipped.
- `useserviceaccount` (String) Clientname to be used as Admin Account

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--accessnodes"></a>
### Nested Schema for `accessnodes`

Optional:

- `type` (Number) Type of access node , Ex: 3 - access Node , 28 - Access Node Groups

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--activitycontrol"></a>
### Nested Schema for `activitycontrol`

Optional:

- `backupactivitycontroloptions` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--backupactivitycontroloptions))
- `enablebackup` (String) true if Backup is enabled
- `enablerestore` (String) true if Restore is enabled
- `restoreactivitycontroloptions` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--restoreactivitycontroloptions))

<a id="nestedblock--activitycontrol--backupactivitycontroloptions"></a>
### Nested Schema for `activitycontrol.backupactivitycontroloptions`

Optional:

- `activitytype` (String) denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]
- `delaytime` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--backupactivitycontroloptions--delaytime))
- `enableactivitytype` (String) True if the activity type is enabled
- `enableafteradelay` (String) True if the activity will be enabled after a delay time interval

<a id="nestedblock--activitycontrol--backupactivitycontroloptions--delaytime"></a>
### Nested Schema for `activitycontrol.backupactivitycontroloptions.delaytime`

Optional:

- `time` (Number) delay time in unix timestamp
- `timezone` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--backupactivitycontroloptions--delaytime--timezone))
- `value` (String) actual delay time value in string format according to the timezone

<a id="nestedblock--activitycontrol--backupactivitycontroloptions--delaytime--timezone"></a>
### Nested Schema for `activitycontrol.backupactivitycontroloptions.delaytime.timezone`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.




<a id="nestedblock--activitycontrol--restoreactivitycontroloptions"></a>
### Nested Schema for `activitycontrol.restoreactivitycontroloptions`

Optional:

- `activitytype` (String) denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]
- `delaytime` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--restoreactivitycontroloptions--delaytime))
- `enableactivitytype` (String) True if the activity type is enabled
- `enableafteradelay` (String) True if the activity will be enabled after a delay time interval

<a id="nestedblock--activitycontrol--restoreactivitycontroloptions--delaytime"></a>
### Nested Schema for `activitycontrol.restoreactivitycontroloptions.delaytime`

Optional:

- `time` (Number) delay time in unix timestamp
- `timezone` (Block List) (see [below for nested schema](#nestedblock--activitycontrol--restoreactivitycontroloptions--delaytime--timezone))
- `value` (String) actual delay time value in string format according to the timezone

<a id="nestedblock--activitycontrol--restoreactivitycontroloptions--delaytime--timezone"></a>
### Nested Schema for `activitycontrol.restoreactivitycontroloptions.delaytime.timezone`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.





<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--fbrunixmediaagent"></a>
### Nested Schema for `fbrunixmediaagent`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.




<a id="nestedblock--settings"></a>
### Nested Schema for `settings`

Optional:
- `guestcredentials` (Block List) (see [below for nested schema](#nestedblock--settings--guestcredentials))
- `regioninfo` (Block List) (see [below for nested schema](#nestedblock--settings--regioninfo))
- `timezone` (Block List) (see [below for nested schema](#nestedblock--settings--timezone))



<a id="nestedblock--settings--guestcredentials"></a>
### Nested Schema for `settings.guestcredentials`

Optional:

- `name` (String) username to access the network path
- `password` (String, Sensitive) password to access the network path





<a id="nestedblock--settings--regioninfo"></a>
### Nested Schema for `settings.regioninfo`

Optional:
- `id` (Number) Region Id
- `name` (String) Region Name


<a id="nestedblock--settings--timezone"></a>
### Nested Schema for `settings.timezone`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.

