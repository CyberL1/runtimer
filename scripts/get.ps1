#!/usr/bin/env pwsh

$ErrorActionPreference = 'Stop'

$RuntimerPath = "${Home}\.runtimer\bin"
$RuntimerZip = "$RuntimerPath\runtimer.zip"
$RuntimerExe = "$RuntimerPath\runtimer.exe"
$RuntimerOldExe = "$env:Temp\runtimerold.exe"

$Target = "windows-amd64"

$DownloadUrl = "https://github.com/CyberL1/runtimer/releases/latest/download/runtimer-${Target}.zip"

if (!(Test-Path $RuntimerPath)) {
  New-Item $RuntimerPath -ItemType Directory | Out-Null
}

curl.exe -Lo $RuntimerZip $DownloadUrl

if (Test-Path $RuntimerExe) {
  Move-Item -Path $RuntimerExe -Destination $RuntimerOldExe -Force
}

Expand-Archive -LiteralPath $RuntimerZip -DestinationPath $RuntimerPath -Force
Remove-Item $RuntimerZip

$User = [System.EnvironmentVariableTarget]::User
$Path = [System.Environment]::GetEnvironmentVariable('Path', $User)

if (!(";${Path};".ToLower() -like "*;${RuntimerPath};*".ToLower())) {
  [System.Environment]::SetEnvironmentVariable('Path', "${Path};${RuntimerPath}", $User)
  $Env:Path += ";${RuntimerPath}"
}

Write-Output "Runtimer was installed to $RuntimerExe"
Write-Output "Run 'runtimer --help' to get started"