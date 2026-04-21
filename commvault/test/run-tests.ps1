# PowerShell script to run Commvault Oracle Terraform Provider tests
# Usage: .\run-tests.ps1 [test-pattern]
# Example: .\run-tests.ps1 TestAccResourceOracleInstance_basic

param(
    [string]$TestPattern = "",
    [int]$Timeout = 30,
    [switch]$Build = $false,
    [switch]$Verbose = $false
)

# Color output functions
function Write-Success { param($Message) Write-Host $Message -ForegroundColor Green }
function Write-Info { param($Message) Write-Host $Message -ForegroundColor Cyan }
function Write-Warning { param($Message) Write-Host $Message -ForegroundColor Yellow }
function Write-Error { param($Message) Write-Host $Message -ForegroundColor Red }

# Check if environment variable is set
function Test-EnvVar {
    param($Name)
    $value = [Environment]::GetEnvironmentVariable($Name)
    if (-not $value) {
        Write-Error "Required environment variable $Name is not set!"
        return $false
    }
    Write-Info "Using $Name = $value"
    return $true
}

Write-Info "=== Commvault Oracle Terraform Provider Test Runner ==="
Write-Info ""

# Enable acceptance tests
$env:TF_ACC = "1"

# Verify all required environment variables are set
Write-Info "Checking required environment variables..."
$allSet = $true
$allSet = (Test-EnvVar "CV_TEST_WEB_SERVICE_URL") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_USER_NAME") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_PASSWORD") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_CLIENT_NAME") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_INSTANCE_NAME") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_ORACLE_HOME") -and $allSet
$allSet = (Test-EnvVar "CV_TEST_ORACLE_USER") -and $allSet

if (-not $allSet) {
    Write-Error ""
    Write-Error "Please set all required environment variables before running tests."
    Write-Error "See README.md or .env.example for configuration details."
    exit 1
}
Write-Info ""

# Ensure Go is in PATH
$env:Path = "C:\Program Files\Go\bin;$env:Path"

# Change to project directory
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$projectDir = Split-Path -Parent $scriptDir
$projectDir = Split-Path -Parent $projectDir
Set-Location $projectDir

# Build if requested
if ($Build) {
    Write-Info "Building provider..."
    go build -o terraform-provider-commvault.exe
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Build failed!"
        exit 1
    }
    Write-Success "Build successful!"
    Write-Info ""
}

# Construct test command
$testCmd = "go test"
if ($Verbose) {
    $testCmd += " -v"
}
if ($TestPattern) {
    $testCmd += " -run $TestPattern"
}
$testCmd += " ./commvault/test/... -timeout ${Timeout}m"

Write-Info "Running tests..."
Write-Info "Command: $testCmd"
Write-Info ""

# Run tests
Invoke-Expression $testCmd
$testResult = $LASTEXITCODE

Write-Info ""
if ($testResult -eq 0) {
    Write-Success "=== Tests Passed ==="
} else {
    Write-Error "=== Tests Failed ==="
}

exit $testResult
