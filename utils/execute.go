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

func ExecuteRuntime(run *constants.RuntimesType, standalone bool, multiple bool, args []string) {
	// Get cache
	cached := cache.Get()

	// Check for custom version
	if standalone {
		if strings.Contains(args[0], "@") {
			nameVer := strings.Split(args[0], "@")
			run, _ = constants.GetDefinedRuntime(nameVer[0])
			run.Version = nameVer[1]
			args = args[1:]
		}
	}

	// Get runtime metadata
	var err error
	if run.Url == "" {
		run, err = constants.GetDefinedRuntime(run.Runtime)
	}
	if err != nil {
		fmt.Print(err)
		return
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

	if !standalone {
		// Get custom
		_, err = GetLocalConfig()
		if err == nil && multiple && args[0] == "-r" {
			custom, _ := GetCustomRuntimeByName(args[1])
			args = args[2:]
			if custom.Version != "" {
				run.Version = custom.Version
			}
		}
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