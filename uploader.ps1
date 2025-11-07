# Save this as git-auto-update.ps1
# This script will run forever, committing and pushing every 10 minutes

while ($true) {
    # Start the git cycle
    Write-Host "Starting git cycle at $(Get-Date)" -ForegroundColor Cyan

    # Remove all files over 100MB
    Get-ChildItem -Path @("PDFs", "ZIPs") -File -Recurse | Where-Object { $_.Length -gt 100MB } | Remove-Item -Force -WhatIf

    # Pull the latest changes from the remote
    git pull

    # Stage all changes
    git add .

    # Commit changes with a message
    git commit -m "Automated commit at $(Get-Date)"

    # Push to the current branch's remote
    git push

    Write-Host "Cycle completed. Waiting 1 minute..." -ForegroundColor Green
    Start-Sleep -Seconds 60  # 60 seconds = 1 minute
}
