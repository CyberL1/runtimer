package constants

import (
	"os"
	"path/filepath"
)

var (
	homeDir, _  = os.UserHomeDir()
	RuntimerDir = filepath.Join(homeDir, ".runtimer")
	RuntimesDir = filepath.Join(RuntimerDir, "runtimes")
	CacheFile   = filepath.Join(RuntimesDir, "cache.json")
	Version     string
)

var (
	GithubReleaseUrl  = "https://api.github.com/repos/CyberL1/runtimer/releases/latest"
	GithubRuntimesUrl = "https://api.github.com/repos/CyberL1/runtimer-test/contents/runtimes"
)
