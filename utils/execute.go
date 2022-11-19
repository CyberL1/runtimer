package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtimer/cache"
	"runtimer/constants"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	"github.com/mholt/archiver"
)

func ExecuteRuntime(config RuntimeType, args []string) {
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

	// Get cache
	cached, err := cache.Get()
	if err != nil {
		fmt.Print(err)
		return
	}

	// Build runtime metadata
	var runtime constants.RuntimesType
	if strings.HasPrefix(config.Runtime, "https://") {
		runtime.Url = config.Runtime
	} else {
		runtime, err = constants.GetDefinedRuntime(config.Runtime)
		if err != nil {
			fmt.Print(err)
			return
		}
	}

	replacer := strings.NewReplacer("$v", runtime.Version,
	"$o", o.Name,
	"$a", a,
	"$e", o.Ext)

	finalUrl := replacer.Replace(runtime.Url)

	// Get runtime
	archive, err := grab.Get(constants.CacheDir, finalUrl)
	if err != nil {
		fmt.Println("Error while getting runtime:", err)
		return
	}

	// Unarchive runtime
	runtimeDir := filepath.Join(constants.CacheDir, runtime.Name)
	archiver.Unarchive(archive.Filename, constants.CacheDir + string(os.PathSeparator) + runtime.Name)

	// Delete archive
	os.Remove(archive.Filename)

	// Execute runtime
	cmd := exec.Command(runtimeDir + string(os.PathSeparator) + filepath.FromSlash(replacer.Replace(o.Bin)), args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		if !cached[runtime.Name] {
			os.RemoveAll(runtimeDir)
		}
	}()
		cmd.Run()
	// Delete runtime if not cached
	if !cached[runtime.Name] {
		os.RemoveAll(runtimeDir)
	}
}