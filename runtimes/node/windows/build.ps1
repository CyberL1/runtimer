#!/usr/bin/pwsh

Invoke-WebRequest "https://nodejs.org/dist/v20.5.0/node-v20.5.0-win-x64.zip" -o node.zip
Expand-Archive node.zip
Remove-Item node.zip
