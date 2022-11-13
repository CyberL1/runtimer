package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/mholt/archiver"
)

func ExecuteRuntime(config Runtime, args []string) {
	// Get os configuration
	o := config.Os[runtime.GOOS]
	if o.Name == "" {
		o.Name = runtime.GOOS
	}
	if o.Ext == "" {
		o.Ext = config.Ext
	}
	if o.Bin == "" {
		o.Bin = config.Bin
	}

	// Get arch configuration
	a := config.Arch[runtime.GOARCH]
	if a == "" {
		a = runtime.GOARCH
	}

	// Check for cache directory
	homeDir, _ := os.UserHomeDir()
	TempDir := filepath.Join(homeDir, ".runtimer", "tmp")

	_, err := os.Stat(TempDir)

	if err != nil {
		os.Mkdir(TempDir, 0755)
	}

	// Build runtime url
	var runtimeUrl string
	if strings.HasPrefix(config.Runtime, "https://") {
		runtimeUrl = config.Runtime
	} else {
		runtimeUrl = RuntimeUrls[config.Runtime]
	}

	replacer := strings.NewReplacer("$v", config.Version,
	"$o", o.Name,
	"$a", a,
	"$e", o.Ext)

	finalUrl := replacer.Replace(runtimeUrl)

	// Get runtime
	archive, err := grab.Get(TempDir, finalUrl)
	if err != nil {
		fmt.Println("Error while getting runtime:", err)
		return
	}

	// Unarchive runtime
	archiver.Unarchive(archive.Filename, TempDir)
	unarchived := strings.TrimSuffix(filepath.Base(archive.Filename), "."+o.Ext)

	// Delete archive
	os.Remove(archive.Filename)

	// Execute runtime
	cmd := exec.Command(TempDir + string(os.PathSeparator) + filepath.FromSlash(strings.ReplaceAll(o.Bin, "$d", unarchived)), args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(cmd.Stderr)
		os.RemoveAll(TempDir)
		return
	}
	fmt.Print(string(out))

	// Delete runtime
	os.RemoveAll(TempDir)
}