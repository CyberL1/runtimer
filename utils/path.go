package utils

import (
	"path/filepath"
	"strings"
)

func fixPath(path string, runtimeDir string) string {
	if strings.HasPrefix(path, "/") {
		return runtimeDir + filepath.Clean(path)
	}
	return path
}