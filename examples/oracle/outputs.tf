output "instance_id" {
  description = "The Commvault internal ID for the Oracle instance"
  value       = commvault_oracle_instance.example.id
}

output "subclient_id" {
  description = "The Commvault internal ID for the Oracle subclient"
  value       = commvault_oracle_subclient.example.id
}

output "instance_oracle_home" {
  description = "Oracle home path confirmed by Commvault"
  value       = data.commvault_oracle_instance.example.oracle_home
}

output "instance_block_size" {
  description = "RMAN block size in bytes"
  value       = data.commvault_oracle_instance.example.block_size
}

output "instance_use_catalog" {
  description = "Whether RMAN recovery catalog is configured"
  value       = data.commvault_oracle_instance.example.use_catalog_connect
}

output "subclient_backup_enabled" {
  description = "Whether backup is currently enabled for the subclient"
  value       = data.commvault_oracle_subclient.example.enable_backup
}
