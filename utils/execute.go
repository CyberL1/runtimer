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
	// Get cache
	cached, err := cache.Get()
	if err != nil {
		fmt.Print(err)
		return
	}

	// Build runtime metadata
	var run constants.RuntimesType
	if strings.HasPrefix(config.Runtime, "https://") {
		run.Url = config.Runtime
	} else {
		run, err = constants.GetDefinedRuntime(config.Runtime)
		if err != nil {
			fmt.Print(err)
			return
		}
	}

	// Get os configuration
	o := run.Os[runtime.GOOS]
	if o.Name == "" {
		o.Name = runtime.GOOS
	}
	if o.Ext == "" {
		o.Ext = run.Ext
	}
	if o.Bin == "" {
		o.Bin = run.Bin
	}

	// Get arch configuration
	a := run.Arch[runtime.GOARCH]
	if a == "" {
		a = runtime.GOARCH
	}

	replacer := strings.NewReplacer("$v", run.Version,
	"$o", o.Name,
	"$a", a,
	"$e", o.Ext)

	finalUrl := replacer.Replace(run.Url)

	// Get runtime
	archive, err := grab.Get(constants.CacheDir, finalUrl)
	if err != nil {
		fmt.Println("Error while getting runtime:", err)
		return
	}

	// Unarchive runtime
	runtimeDir := filepath.Join(constants.CacheDir, run.Name)
	archiver.Unarchive(archive.Filename, constants.CacheDir + string(os.PathSeparator) + run.Name)

	// Delete archive
	os.Remove(archive.Filename)

	// Execute runtime
	cmd := exec.Command(fixPath(replacer.Replace(o.Bin), runtimeDir), args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		if !cached[run.Name] {
			os.RemoveAll(runtimeDir)
		}
	}()
		cmd.Run()
	// Delete runtime if not cached
	if !cached[run.Name] {
		os.RemoveAll(runtimeDir)
	}
}