# Terraform Provider Commvault - Oracle Tests

This directory contains acceptance tests for the Oracle resources in the Commvault Terraform provider.

## Prerequisites

1. **Commvault Environment**: A running Commvault environment with:
   - Web Console API accessible
   - Valid admin credentials
   - Oracle client(s) registered (for instance and subclient tests)

2. **Go**: Go 1.21 or later installed

3. **Test Environment Variables**: Set the required environment variables (see Configuration below)

## Test Files

- `helpers_test.go` - Common test helpers, provider configuration, and environment variable handling
- `instance_test.go` - Tests for `commvault_oracle_instance` resource
- `subclient_test.go` - Tests for `commvault_oracle_subclient` resource
- `install_agent_test.go` - Tests for `commvault_oracle_install_agent` resource (skipped by default)

## Configuration

The tests require environment variables for configuration. **All environment variables are required** - there are no default values.

### Required Environment Variables

```powershell
# Enable acceptance tests (required)
$env:TF_ACC = "1"

# Commvault environment (required)
$env:CV_TEST_WEB_SERVICE_URL = "https://your-commserve.example.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-admin-password"

# Oracle test configuration (required)
$env:CV_TEST_CLIENT_NAME = "your-oracle-client-name"
$env:CV_TEST_INSTANCE_NAME = "YOUR_ORACLE_INSTANCE"
$env:CV_TEST_ORACLE_HOME = "/path/to/oracle/home"
$env:CV_TEST_ORACLE_USER = "oracle-os-user"
```

**Important**: Replace all placeholder values with your actual Commvault environment details. Tests will fail with a clear error message if any required environment variable is not set.

## Running Tests

### Run All Oracle Tests

```powershell
# Set environment variables
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://your-commserve.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-password"
$env:CV_TEST_CLIENT_NAME = "your-oracle-client"
$env:CV_TEST_INSTANCE_NAME = "YOUR_INSTANCE"
$env:CV_TEST_ORACLE_HOME = "/path/to/oracle/home"

# Run tests
cd c:\Users\maheshp\terraform-provider-commvault
go test -v ./commvault/test/... -timeout 30m
```

### Run Specific Test

```powershell
# Run only the basic instance test
go test -v -run TestAccResourceOracleInstance_basic ./commvault/test/... -timeout 30m

# Run all instance tests
go test -v -run TestAccResourceOracleInstance ./commvault/test/... -timeout 30m

# Run subclient tests
go test -v -run TestAccResourceOracleSubclient ./commvault/test/... -timeout 30m
```

### Run with Different Configuration

```powershell
# Test against a different environment
$env:CV_TEST_WEB_SERVICE_URL = "https://prod-cs.example.com/webconsole/api"
$env:CV_TEST_CLIENT_NAME = "oracle-prod-01"
$env:CV_TEST_INSTANCE_NAME = "PROD_DB"
go test -v ./commvault/test/... -timeout 30m
```

## Test Scenarios

### Oracle Instance Tests

1. **TestAccResourceOracleInstance_basic**: Creates a basic Oracle instance with minimal configuration
2. **TestAccResourceOracleInstance_withSqlConnect**: Tests SQL connection configuration
3. **TestAccResourceOracleInstance_withCatalog**: Tests RMAN catalog configuration

### Oracle Subclient Tests

1. **TestAccResourceOracleSubclient_basic**: Creates a basic Oracle subclient

### Install Agent Tests

The install agent test is skipped by default as it requires a target server for agent installation. To enable:

```go
// In install_agent_test.go, comment out or remove this line:
// t.Skip("Skipping install agent test - requires target server setup")
```

## Troubleshooting

### Test Fails with "Instance Already Exists"

If the test instance already exists on your client, you can:

1. **Use a different instance name**:
   ```powershell
   $env:CV_TEST_INSTANCE_NAME = "TESTDB_TF"
   ```

2. **Delete the existing instance** through Commvault before running tests

3. **Update the test** to handle existing instances (for update tests)

### Authentication Errors

Ensure your credentials are correct:
```powershell
$env:CV_TEST_USER_NAME = "your-username"
$env:CV_TEST_PASSWORD = "your-password"
```

The provider automatically base64 encodes passwords for the API.

### Connection Errors

Check that:
- The Web Console API is accessible from your test machine
- SSL certificate issues are handled (the tests use `ignore_cert = true`)
- Firewall rules allow access to the CommServe

### Timeout Issues

Some operations may take longer on slower environments. Increase the timeout:
```powershell
go test -v ./commvault/test/... -timeout 60m
```

## Best Practices

1. **Use Dedicated Test Environment**: Don't run tests against production
2. **Clean Test Data**: Clean up test resources after running tests
3. **Use Unique Names**: Use unique instance/subclient names to avoid conflicts
4. **Secure Credentials**: Use environment variables, never commit credentials to code
5. **Test Isolation**: Each test should be independent and not rely on other tests

## Example: Complete Test Run

```powershell
# PowerShell script to run all Oracle tests

# Set environment
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://your-commserve.example.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-password"
$env:CV_TEST_CLIENT_NAME = "hugo"
$env:CV_TEST_INSTANCE_NAME = "TESTDB_TF_$(Get-Date -Format 'yyyyMMddHHmmss')"
$env:CV_TEST_ORACLE_HOME = "/u01/app/oracle/product/19c/dbhome_1"
$env:CV_TEST_ORACLE_USER = "oracle"

# Build provider
Write-Host "Building provider..." -ForegroundColor Green
go build -o terraform-provider-commvault.exe

# Run tests
Write-Host "Running Oracle tests..." -ForegroundColor Green
go test -v ./commvault/test/... -timeout 30m

Write-Host "Tests complete!" -ForegroundColor Green
```

## CI/CD Integration

For CI/CD pipelines, set environment variables in your pipeline configuration:

```yaml
# Example GitHub Actions
env:
  TF_ACC: "1"
  CV_TEST_WEB_SERVICE_URL: ${{ secrets.CV_WEB_SERVICE_URL }}
  CV_TEST_USER_NAME: ${{ secrets.CV_USER_NAME }}
  CV_TEST_PASSWORD: ${{ secrets.CV_PASSWORD }}
  CV_TEST_CLIENT_NAME: "ci-test-client"
  CV_TEST_INSTANCE_NAME: "CI_TEST_DB"
```

## Contributing

When adding new tests:

1. Follow the existing test structure
2. Use the helper functions from `helpers_test.go`
3. Use environment variables for all configuration
4. Add documentation for any new environment variables
5. Ensure tests can run independently
6. Add cleanup steps where necessary
