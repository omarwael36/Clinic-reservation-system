# Jenkins Installation and Setup Script
# Run this as Administrator

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Jenkins Installation and Setup" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

# Check if running as administrator
$isAdmin = ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $isAdmin) {
    Write-Host "ERROR: This script must be run as Administrator!" -ForegroundColor Red
    Write-Host "Right-click PowerShell and select 'Run as Administrator'" -ForegroundColor Yellow
    exit 1
}

# Step 1: Download Jenkins
Write-Host "[1/7] Downloading Jenkins..." -ForegroundColor Green
$jenkinsUrl = "https://get.jenkins.io/windows-stable/2.479.3/jenkins.msi"
$installerPath = "$env:TEMP\jenkins.msi"

try {
    Invoke-WebRequest -Uri $jenkinsUrl -OutFile $installerPath
    Write-Host "  ✓ Jenkins downloaded successfully" -ForegroundColor Green
} catch {
    Write-Host "  ✗ Failed to download Jenkins: $_" -ForegroundColor Red
    exit 1
}

# Step 2: Install Jenkins
Write-Host "`n[2/7] Installing Jenkins..." -ForegroundColor Green
try {
    Start-Process msiexec.exe -Wait -ArgumentList "/I `"$installerPath`" /quiet /qn /norestart"
    Write-Host "  ✓ Jenkins installed successfully" -ForegroundColor Green
} catch {
    Write-Host "  ✗ Failed to install Jenkins: $_" -ForegroundColor Red
    exit 1
}

# Step 3: Wait for Jenkins service to be created
Write-Host "`n[3/7] Waiting for Jenkins service..." -ForegroundColor Green
Start-Sleep -Seconds 10

# Step 4: Start Jenkins service
Write-Host "`n[4/7] Starting Jenkins service..." -ForegroundColor Green
try {
    $service = Get-Service -Name "Jenkins" -ErrorAction SilentlyContinue
    if ($service) {
        if ($service.Status -ne 'Running') {
            Start-Service -Name "Jenkins"
            Start-Sleep -Seconds 20
        }
        Write-Host "  ✓ Jenkins service is running" -ForegroundColor Green
    } else {
        Write-Host "  ✗ Jenkins service not found" -ForegroundColor Red
        exit 1
    }
} catch {
    Write-Host "  ✗ Failed to start Jenkins: $_" -ForegroundColor Red
    exit 1
}

# Step 5: Get initial admin password
Write-Host "`n[5/7] Retrieving initial admin password..." -ForegroundColor Green
$passwordFile = "C:\ProgramData\Jenkins\.jenkins\secrets\initialAdminPassword"
$jenkinsHome = "C:\ProgramData\Jenkins\.jenkins"

# Wait for password file to be created
$maxWait = 60
$waited = 0
while (-not (Test-Path $passwordFile) -and $waited -lt $maxWait) {
    Start-Sleep -Seconds 5
    $waited += 5
    Write-Host "  Waiting for Jenkins to initialize... ($waited/$maxWait seconds)" -ForegroundColor Yellow
}

if (Test-Path $passwordFile) {
    $initialPassword = Get-Content $passwordFile
    Write-Host "  ✓ Initial Admin Password: $initialPassword" -ForegroundColor Green
    Write-Host "  ✓ Jenkins Home: $jenkinsHome" -ForegroundColor Green
    
    # Save to a file for easy access
    $initialPassword | Out-File -FilePath "jenkins-initial-password.txt"
    Write-Host "  ✓ Password saved to: jenkins-initial-password.txt" -ForegroundColor Green
} else {
    Write-Host "  ✗ Could not find initial password file" -ForegroundColor Red
    Write-Host "  Check manually at: $passwordFile" -ForegroundColor Yellow
}

# Step 6: Open Jenkins in browser
Write-Host "`n[6/7] Opening Jenkins in browser..." -ForegroundColor Green
Start-Sleep -Seconds 5
Start-Process "http://localhost:8080"
Write-Host "  ✓ Jenkins should open at http://localhost:8080" -ForegroundColor Green

# Step 7: Display next steps
Write-Host "`n[7/7] Installation Complete!" -ForegroundColor Green
Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "NEXT STEPS:" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

Write-Host "1. Jenkins is now running at: http://localhost:8080" -ForegroundColor White
Write-Host "2. Use the initial admin password above to unlock Jenkins" -ForegroundColor White
Write-Host "3. Select 'Install suggested plugins'" -ForegroundColor White
Write-Host "4. Create your admin user" -ForegroundColor White
Write-Host "5. After setup, run: .\configure-jenkins.ps1" -ForegroundColor Yellow
Write-Host "`n========================================`n" -ForegroundColor Cyan

Write-Host "Press any key to exit..." -ForegroundColor Gray
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")

