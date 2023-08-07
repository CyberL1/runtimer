#!/usr/bin/pwsh

Invoke-WebRequest "https://go.dev/dl/go1.20.7.windows-amd64.zip" -o go.zip
Expand-Archive go.zip
Remove-Item go.zip
