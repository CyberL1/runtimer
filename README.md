# Runtimer
A CLI for running different runtimes without installing them

# Instalation

Windows:
```
irm https://raw.githubusercontent.com/CyberL1/runtimer/main/scripts/get.ps1 | iex
```

Linux:
```
curl -fsSL https://raw.githubusercontent.com/CyberL1/runtimer/main/scripts/get.sh | sh
```

# Development

1. Start a local dev server:
```
runtimer dev
```

2. Change runtimes URL to localhost:
```
go run -ldflags "-X runtimer/constants.GithubRuntimesUrl=http://localhost:4786" main.go
```
