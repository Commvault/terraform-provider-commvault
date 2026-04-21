# Security Update: Removed Hardcoded Values from Tests

## Changes Made

All hardcoded default values have been removed from the test configuration files to improve security and prevent accidental exposure of sensitive information.

### Files Updated

1. **`commvault/test/helpers_test.go`**
   - Changed from `getEnvOrDefault()` with hardcoded defaults to `getEnvOrFail()`
   - All environment variables are now **required**
   - Tests will fail with clear error message if any variable is missing

2. **`commvault/test/run-tests.ps1`**
   - Removed all default values
   - Added validation to check all required environment variables
   - Fails fast with helpful error message if variables are not set

3. **`commvault/test/.env.example`**
   - Replaced actual values with placeholder text
   - Users must replace with their own values

4. **`commvault/test/README.md`**
   - Updated to reflect that all variables are required
   - Removed references to "default values"

## Required Environment Variables

All of these must be set before running tests:

```powershell
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://your-commserve.example.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-password"
$env:CV_TEST_CLIENT_NAME = "your-oracle-client"
$env:CV_TEST_INSTANCE_NAME = "YOUR_INSTANCE"
$env:CV_TEST_ORACLE_HOME = "/path/to/oracle/home"
$env:CV_TEST_ORACLE_USER = "oracle-user"
```

## Validation

The test runner now validates all environment variables before running tests:

```powershell
cd c:\Users\maheshp\terraform-provider-commvault\commvault\test
.\run-tests.ps1

# Output if variables not set:
=== Commvault Oracle Terraform Provider Test Runner ===

Checking required environment variables...
Required environment variable CV_TEST_WEB_SERVICE_URL is not set!
Required environment variable CV_TEST_USER_NAME is not set!
Required environment variable CV_TEST_PASSWORD is not set!
Required environment variable CV_TEST_CLIENT_NAME is not set!
Required environment variable CV_TEST_INSTANCE_NAME is not set!
Required environment variable CV_TEST_ORACLE_HOME is not set!
Required environment variable CV_TEST_ORACLE_USER is not set!

Please set all required environment variables before running tests.
See README.md or .env.example for configuration details.
```

## Security Benefits

1. **No Credentials in Code**: All sensitive information must be provided at runtime
2. **No Accidental Commits**: Can't accidentally commit actual credentials in test files
3. **Environment-Specific**: Forces users to configure for their environment
4. **Clear Errors**: Users know immediately if configuration is missing
5. **CI/CD Friendly**: Works seamlessly with secret management in pipelines

## Usage

### Option 1: Set Variables Directly
```powershell
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://your-cs.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-password"
$env:CV_TEST_CLIENT_NAME = "your-client"
$env:CV_TEST_INSTANCE_NAME = "YOUR_DB"
$env:CV_TEST_ORACLE_HOME = "/path/to/oracle"
$env:CV_TEST_ORACLE_USER = "oracle"

cd c:\Users\maheshp\terraform-provider-commvault\commvault\test
.\run-tests.ps1 -Verbose
```

### Option 2: Use .env File
```powershell
# Copy and edit .env
Copy-Item .env.example .env
notepad .env  # Edit with your actual values

# Load variables
Get-Content .env | ForEach-Object {
    if($_ -match '^([^=]+)=(.*)$') {
        [Environment]::SetEnvironmentVariable($matches[1], $matches[2])
    }
}

# Run tests
.\run-tests.ps1 -Verbose
```

### Option 3: CI/CD Pipeline
```yaml
# GitHub Actions example
env:
  TF_ACC: "1"
  CV_TEST_WEB_SERVICE_URL: ${{ secrets.CV_WEB_SERVICE_URL }}
  CV_TEST_USER_NAME: ${{ secrets.CV_USER_NAME }}
  CV_TEST_PASSWORD: ${{ secrets.CV_PASSWORD }}
  CV_TEST_CLIENT_NAME: "ci-test-client"
  CV_TEST_INSTANCE_NAME: "CI_TEST_DB"
  CV_TEST_ORACLE_HOME: "/u01/app/oracle"
  CV_TEST_ORACLE_USER: "oracle"

steps:
  - name: Run tests
    run: |
      cd commvault/test
      ./run-tests.ps1 -Verbose
```

## Migration for Existing Users

If you were relying on the old default values, you now need to explicitly set them:

**Before** (automatic defaults):
```powershell
cd commvault/test
.\run-tests.ps1  # Used built-in defaults
```

**After** (explicit configuration required):
```powershell
# Set your environment
$env:TF_ACC = "1"
$env:CV_TEST_WEB_SERVICE_URL = "https://your-commserve.example.com/webconsole/api"
$env:CV_TEST_USER_NAME = "your-admin-username"
$env:CV_TEST_PASSWORD = "your-password"
$env:CV_TEST_CLIENT_NAME = "hugo"
$env:CV_TEST_INSTANCE_NAME = "TESTDB"
$env:CV_TEST_ORACLE_HOME = "/u01/app/oracle/product/19c/dbhome_1"
$env:CV_TEST_ORACLE_USER = "oracle"

# Run tests
cd commvault/test
.\run-tests.ps1 -Verbose
```

## Summary

✅ All hardcoded values removed from test files
✅ Environment variables are now mandatory
✅ Clear validation and error messages
✅ Improved security posture
✅ Better CI/CD integration
✅ Forces explicit configuration per environment
