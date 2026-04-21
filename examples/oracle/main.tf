# -----------------------------------------------------------------------
# Oracle Instance
# Registers an Oracle database instance in Commvault for backup management
# -----------------------------------------------------------------------
resource "commvault_oracle_instance" "example" {
  client_name   = var.client_name
  instance_name = var.instance_name
  oracle_home   = var.oracle_home
  oracle_user   = var.oracle_user

  # Database connection using OS authentication
  # '/' as username means "use OS-level authentication" (no password needed)
  sql_connect_user   = "/"
  sql_connect_domain = var.instance_name   # Oracle SID or TNS alias

  # RMAN settings
  block_size          = 1048576   # 1 MB - recommended default
  cross_check_timeout = 600       # 10 minutes

  # Uncomment to associate a Commvault protection plan
  # plan_id = var.plan_id
}

# -----------------------------------------------------------------------
# Oracle Subclient
# Defines what gets backed up within the Oracle instance
# Must be created after the instance
# -----------------------------------------------------------------------
resource "commvault_oracle_subclient" "example" {
  client_name    = var.client_name
  instance_name  = var.instance_name
  subclient_name = var.subclient_name
  description    = "Managed by Terraform"

  # What to include in the backup
  enable_backup                   = true
  backup_archive_log              = true    # Back up archive logs
  backup_sp_file                  = true    # Back up the server parameter file
  backup_control_file             = true    # Back up the control file
  delete_archive_log_after_backup = false   # Keep archive logs after backup

  depends_on = [commvault_oracle_instance.example]
}

# -----------------------------------------------------------------------
# Data Source: Read back the instance details from Commvault
# -----------------------------------------------------------------------
data "commvault_oracle_instance" "example" {
  client_name   = var.client_name
  instance_name = var.instance_name

  depends_on = [commvault_oracle_instance.example]
}

# -----------------------------------------------------------------------
# Data Source: Read back the subclient details from Commvault
# -----------------------------------------------------------------------
data "commvault_oracle_subclient" "example" {
  client_name    = var.client_name
  instance_name  = var.instance_name
  subclient_name = var.subclient_name

  depends_on = [commvault_oracle_subclient.example]
}
