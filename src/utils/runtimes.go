package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtimer/constants"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

func GetRuntimes() ([]GithubFile, error) {
	resp, _ := http.Get(constants.GithubRuntimesUrl)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var runtimes []GithubFile
	err = json.Unmarshal(body, &runtimes)
	if err != nil {
		return nil, err
	}

	return runtimes, nil
}

func GetRuntime(name string) (*Runtime, error) {
	var dir string

	switch runtime.GOOS {
	case "linux", "darwin":
		dir = "linux"
	case "windows":
		dir = "windows"
	}

	resp, _ := http.Get(fmt.Sprintf(constants.GithubRuntimesUrl+"/%v/%v", name, dir))

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("runtime %s not found", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	runtimeDir := filepath.Join(constants.RuntimesDir, name)
	var files []GithubFile
	err = json.Unmarshal(body, &files)
	if err != nil {
		return nil, err
	}

	runtime := &Runtime{
		Name:      name,
		Directory: runtimeDir,
		Files:     files,
	}
	return runtime, nil
}

func (r Runtime) Download() {
	for _, f := range r.Files {
		grab.Get(filepath.Join(r.Directory, f.Name), f.DownloadUrl)
	}
}

func (r Runtime) Execute(args []string) {
	var command string
	var ext string
	var script string
	var interrupted bool

	switch runtime.GOOS {
	case "linux", "darwin":
		command = "sh"
		ext = "sh"
	case "windows":
		command = "powershell"
		ext = "ps1"
	}

	if !IsCached(r.Name) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			interrupted = true
		}()

		if !interrupted {
			r.Download()
			for _, f := range r.Files {
				if interrupted {
					break
				}

				cmd := exec.Command(command, filepath.Join(r.Directory, f.Name))
				cmd.Dir = r.Directory
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout

				if strings.Split(f.Name, ".")[0] == "run" {
					if args != nil {
						cmd.Args = append(cmd.Args, args...)
					} else {
						continue
					}
				}
				cmd.Run()
			}
		}
		os.RemoveAll(r.Directory)
	} else {
		if args != nil {
			script = "run"
		} else {
			script = "build"
		}

		cmd := exec.Command(command, filepath.Join(r.Directory, fmt.Sprintf("%v.%v", script, ext))+" "+strings.Join(args, " "))
		cmd.Dir = r.Directory
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
