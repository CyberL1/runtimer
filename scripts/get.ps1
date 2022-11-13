#!/usr/bin/env pwsh

$ErrorActionPreference = 'Stop'

$Build = $false

$RuntimerPath = "${Home}\.runtimer\bin"
$RuntimerZip = "$RuntimerPath\runtimer.zip"
$RuntimerExe = "$RuntimerPath\runtimer.exe"

$Target = "windows-amd64"

$DownloadUrl = "https://github.com/CyberL1/runtimer/releases/latest/download/runtimer-${Target}.zip"

if (!(Test-Path $RuntimerPath)) {
  New-Item $RuntimerPath -ItemType Directory | Out-Null
}

if ($Build) {
  go build -o $RuntimerExe
} else {
  curl.exe -Lo $RuntimerZip $DownloadUrl
  Expand-Archive -LiteralPath $RuntimerZip -DestinationPath $RuntimerPath
  Remove-Item $RuntimerZip
}

$User = [System.EnvironmentVariableTarget]::User
$Path = [System.Environment]::GetEnvironmentVariable('Path', $User)

if (!(";${Path};".ToLower() -like "*;${RuntimerPath};*".ToLower())) {
  [System.Environment]::SetEnvironmentVariable('Path', "${Path};${RuntimerPath}", $User)
  $Env:Path += ";${RuntimerPath}"
}

Write-Output "Runtimer was installed to $RuntimerExe"
Write-Output "Run 'runtimer --help' to get started"