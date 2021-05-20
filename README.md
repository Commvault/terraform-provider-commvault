<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Commvault

With the Commvault Terraform module, you can use Terraform to manage endpoints (called resources). Terraform is a configuration language for safely and efficiently managing infrastructure.

The Commvault Terraform module provides a set of named resource types, and specifies which arguments are allowed for each resource type. Using the resource types, you can create a configuration file, and apply changes to the Commvault REST APIs.

## Maintainers

This provider plugin is maintained by [Commvault](https://www.commvault.com/)

## Quick Starts :scroll:

 - [Software Requirements](#SoftwareRequirements)
 - [Building the Provider](#BuildProvider)
 - [Using the Provider](#using-the-provider)
 
## <a name ="SoftwareRequirements"></a> Software Requirements :clipboard:

-	[Terraform](https://www.terraform.io/downloads.html) 0.15.0+
-	[Go](https://golang.org/doc/install) 1.12+ (to build the provider plugin)
-   [Commvault](https://www.commvault.com/) V11 Sp22+

## <a name ="BuildProvider"></a> Building The Provider :pencil2:

- Clone the Repository to local machine. 
- Navigate to the Cloned Repository.

######For Windows:

Open Command Prompt in the Cloned Repository locatoion and use below command to build Executable.

```sh
go build -o terraform-provider-commvault.exe
```

######For Linux:

Open Terminal in the Cloned Repository locatoion and use below command to build Executable.

```sh
go build -o terraform-provider-commvault
```


## <a name ="using-the-provider"></a> Using the provider :arrow_forward:

The Commvault provider documentation can be found on [provider's website](https://documentation.commvault.com/11.23/essential/129185_commvault_terraform_module.html) or in [Terraform Regitry](https://registry.terraform.io/providers/Commvault/commvault/latest/docs/index.html)
