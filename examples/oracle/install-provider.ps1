# install-provider.ps1
# Run this script once to install the Commvault Terraform provider on your machine.
# Usage: .\install-provider.ps1
# Usage with custom binary path: .\install-provider.ps1 -BinaryPath "C:\path\to\terraform-provider-commvault.exe"

param(
    [string]$BinaryPath = "$PSScriptRoot\..\..\terraform-provider-commvault.exe"
)

$ErrorActionPreference = "Stop"

# Resolve the binary path
$BinaryPath = Resolve-Path $BinaryPath -ErrorAction SilentlyContinue
if (-not $BinaryPath) {
    Write-Error "Provider binary not found. Please specify the path with -BinaryPath parameter."
    exit 1
}

Write-Host "Installing Commvault Terraform Provider..." -ForegroundColor Cyan
Write-Host "Source binary: $BinaryPath"

# Define the target directory
$pluginDir = "$env:APPDATA\terraform.d\plugins\registry.terraform.io\commvault\commvault\1.0.0\windows_amd64"
$targetFile = "$pluginDir\terraform-provider-commvault_v1.0.0_x5.exe"

# Create the directory
Write-Host "`nCreating plugin directory..." -ForegroundColor Yellow
New-Item -ItemType Directory -Force -Path $pluginDir | Out-Null
Write-Host "  $pluginDir" -ForegroundColor Gray

# Copy the binary
Write-Host "`nCopying provider binary..." -ForegroundColor Yellow
Copy-Item -Path $BinaryPath -Destination $targetFile -Force
Write-Host "  -> $targetFile" -ForegroundColor Gray

# Verify
if (Test-Path $targetFile) {
    $size = (Get-Item $targetFile).Length
    Write-Host "`n✅ Provider installed successfully!" -ForegroundColor Green
    Write-Host "   File size: $([math]::Round($size/1MB, 2)) MB"
} else {
    Write-Error "Installation failed - file not found at target location."
    exit 1
}

Write-Host "`n--- Next Steps ---" -ForegroundColor Cyan
Write-Host "1. Copy terraform.tfvars.example to terraform.tfvars and fill in your values"
Write-Host "2. Run: terraform init"
Write-Host "3. Run: terraform plan"
Write-Host "4. Run: terraform apply"
Write-Host ""
