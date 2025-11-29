# Jenkins Configuration Script
# Run this AFTER Jenkins is set up and running

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Jenkins Configuration Helper" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

$jenkinsUrl = "http://localhost:8081"
$dockerhubUsername = "omarwael01"
$dockerhubPassword = "k:a269K/3#VDrHy"
$githubRepo = "https://github.com/omarwael36/Clinic-reservation-system.git"

Write-Host "Configuration Details:" -ForegroundColor Yellow
Write-Host "  Jenkins URL: $jenkinsUrl"
Write-Host "  DockerHub Username: $dockerhubUsername"
Write-Host "  GitHub Repository: $githubRepo"
Write-Host ""

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "MANUAL CONFIGURATION STEPS" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

Write-Host "STEP 1: Install Required Plugins" -ForegroundColor Green
Write-Host "  1. Go to: Manage Jenkins → Manage Plugins" -ForegroundColor White
Write-Host "  2. Click 'Available' tab and search for:" -ForegroundColor White
Write-Host "     - Docker Pipeline" -ForegroundColor Cyan
Write-Host "     - Docker Plugin" -ForegroundColor Cyan
Write-Host "     - Git Plugin" -ForegroundColor Cyan
Write-Host "  3. Check all and click 'Install without restart'" -ForegroundColor White
Write-Host ""

Write-Host "STEP 2: Add DockerHub Credentials" -ForegroundColor Green
Write-Host "  1. Go to: Manage Jenkins → Manage Credentials" -ForegroundColor White
Write-Host "  2. Click: (global) domain" -ForegroundColor White
Write-Host "  3. Click: Add Credentials" -ForegroundColor White
Write-Host "  4. Fill in:" -ForegroundColor White
Write-Host "     - Kind: Username with password" -ForegroundColor Cyan
Write-Host "     - Username: $dockerhubUsername" -ForegroundColor Cyan
Write-Host "     - Password: $dockerhubPassword" -ForegroundColor Cyan
Write-Host "     - ID: dockerhub-credentials" -ForegroundColor Cyan
Write-Host "     - Description: DockerHub Credentials" -ForegroundColor Cyan
Write-Host "  5. Click: OK" -ForegroundColor White
Write-Host ""

Write-Host "STEP 3: Create Pipeline Job" -ForegroundColor Green
Write-Host "  1. Go to Jenkins Dashboard" -ForegroundColor White
Write-Host "  2. Click: New Item" -ForegroundColor White
Write-Host "  3. Enter name: Clinic-Reservation-System-Pipeline" -ForegroundColor Cyan
Write-Host "  4. Select: Pipeline" -ForegroundColor White
Write-Host "  5. Click: OK" -ForegroundColor White
Write-Host ""

Write-Host "STEP 4: Configure Pipeline" -ForegroundColor Green
Write-Host "  1. Scroll to 'Build Triggers' section" -ForegroundColor White
Write-Host "  2. Check: Poll SCM" -ForegroundColor White
Write-Host "  3. Schedule: H/5 * * * * (checks every 5 minutes)" -ForegroundColor Cyan
Write-Host ""
Write-Host "  4. Scroll to 'Pipeline' section" -ForegroundColor White
Write-Host "  5. Definition: Pipeline script from SCM" -ForegroundColor White
Write-Host "  6. SCM: Git" -ForegroundColor White
Write-Host "  7. Repository URL: $githubRepo" -ForegroundColor Cyan
Write-Host "  8. Credentials: None (public repo)" -ForegroundColor White
Write-Host "  9. Branch Specifier: */main" -ForegroundColor Cyan
Write-Host "  10. Script Path: Jenkinsfile" -ForegroundColor Cyan
Write-Host "  11. Click: Save" -ForegroundColor White
Write-Host ""

Write-Host "STEP 5: Test the Pipeline" -ForegroundColor Green
Write-Host "  1. Click: Build Now" -ForegroundColor White
Write-Host "  2. Watch the build progress" -ForegroundColor White
Write-Host "  3. Check Console Output for any errors" -ForegroundColor White
Write-Host ""

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "QUICK REFERENCE" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

Write-Host "Jenkins URL: " -NoNewline -ForegroundColor White
Write-Host $jenkinsUrl -ForegroundColor Cyan

Write-Host "DockerHub Images will be:" -ForegroundColor White
Write-Host "  - $dockerhubUsername/clinic-database:latest" -ForegroundColor Cyan
Write-Host "  - $dockerhubUsername/clinic-backend:latest" -ForegroundColor Cyan
Write-Host "  - $dockerhubUsername/clinic-frontend:latest" -ForegroundColor Cyan

Write-Host "`nGitHub Repository: " -NoNewline -ForegroundColor White
Write-Host $githubRepo -ForegroundColor Cyan

Write-Host "`n========================================`n" -ForegroundColor Cyan

Write-Host "Opening Jenkins in browser..." -ForegroundColor Green
Start-Process $jenkinsUrl

Write-Host "`nPress any key to exit..." -ForegroundColor Gray
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")

