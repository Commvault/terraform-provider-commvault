---
page_title: "commvault_oracle_rman_logs Data Source"
subcategory: "Oracle"
description: |-
  Use the commvault_oracle_rman_logs data source to retrieve RMAN logs for Oracle backup/restore jobs.
---

# commvault_oracle_rman_logs (Data Source)

Use the commvault_oracle_rman_logs data source to retrieve RMAN (Recovery Manager) logs for Oracle backup or restore jobs. This is useful for troubleshooting and auditing purposes.

## Example Usage

### Get RMAN Logs for a Job

```hcl
data "commvault_oracle_rman_logs" "backup_log" {
  job_id = "29408"
}

output "rman_output" {
  value = data.commvault_oracle_rman_logs.backup_log.log_content
}
```

### Use with Job ID from External Source

```hcl
# Retrieve RMAN logs for a specific job ID
# Job ID can be obtained from Commvault Command Center or other monitoring tools

data "commvault_oracle_rman_logs" "backup_log" {
  job_id = "12345"  # Replace with actual job ID
}

output "backup_rman_log" {
  value     = data.commvault_oracle_rman_logs.backup_log.log_content
  sensitive = true
}
```

## Argument Reference

### Required

- **job_id** (String) - The ID of the Oracle backup or restore job from Commvault.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) - The ID of this data source.
- **log_content** (String) - The full RMAN log output for the job.
- **error_string** (String) - Error message if log retrieval failed.

## Notes

- RMAN logs are available after the job has started executing.
- Large jobs may have extensive log content.
- Logs contain detailed RMAN commands and output for debugging.
