if (-not ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Host "[System] Requesting Administrative privileges..." -ForegroundColor Yellow
    
    # Relaunch this exact script as Administrator, keeping the same working folder
    Start-Process powershell.exe -ArgumentList "-NoProfile -ExecutionPolicy Bypass -WorkingDirectory `"$PSScriptRoot`" -File `"$PSCommandPath`"" -Verb RunAs
    Exit
}

$PORT = 8080
$RuleName = "Remote Server Port $PORT"

function Cleanup {
    Write-Host "`n[Firewall] Closing port $PORT..." -ForegroundColor Cyan
    Remove-NetFirewallRule -DisplayName $RuleName -ErrorAction SilentlyContinue
    Write-Host "[Go-Remote Server] Shutting down." -ForegroundColor Red
    Exit
}

# Open the port in Windows Defender Firewall
Write-Host "[Firewall] Opening port $PORT..." -ForegroundColor Green
New-NetFirewallRule -DisplayName $RuleName -Direction Inbound -Action Allow -Protocol TCP -LocalPort $PORT -ErrorAction SilentlyContinue | Out-Null

try {
    Write-Host "[Go-Remote Server] Starting on port $PORT..." -ForegroundColor Green
    
    $env:PORT = $PORT
    .\remote-server-windows.exe
} 
finally {
    Cleanup
}