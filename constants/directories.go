package constants

import (
	"os"
	"path/filepath"
)

var HomeDir, _ = os.UserHomeDir()
var RuntimerDir = filepath.Join(HomeDir, ".runtimer")
var CacheDir = filepath.Join(RuntimerDir, "cache")
var CachedFile = filepath.Join(CacheDir, "cached.json")