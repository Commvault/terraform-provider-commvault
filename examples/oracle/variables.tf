variable "web_service_url" {
  description = "Commvault CommServe API URL (e.g. https://myserver.example.com/webconsole/api)"
  type        = string
}

variable "user_name" {
  description = "Commvault admin username"
  type        = string
}

variable "password" {
  description = "Commvault admin password"
  type        = string
  sensitive   = true
}

variable "client_name" {
  description = "Name of the client with Oracle agent installed (as it appears in Commvault)"
  type        = string
}

variable "instance_name" {
  description = "Oracle database SID (instance name)"
  type        = string
}

variable "oracle_home" {
  description = "Full path to Oracle home directory on the client machine"
  type        = string
}

variable "oracle_user" {
  description = "OS username that owns the Oracle installation on the client"
  type        = string
}

variable "plan_id" {
  description = "Commvault Plan ID to associate with the instance (0 = no plan)"
  type        = number
  default     = 0
}

variable "subclient_name" {
  description = "Name for the Oracle subclient to create"
  type        = string
  default     = "terraform_subclient"
}
