package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtimer/constants"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

type GithubFiles []GithubFile

func GetRuntimes() (GithubFiles, error) {
	resp, _ := http.Get(constants.GithubRuntimesUrl)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var runtimes GithubFiles
	err = json.Unmarshal(body, &runtimes)
	if err != nil {
		return nil, err
	}

	return runtimes, nil
}

func GetRuntime(name string) (GithubFiles, error) {
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

	var files GithubFiles
	err = json.Unmarshal(body, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (files GithubFiles) Execute(args []string) {
	runtimeName := strings.Split(files[0].Path, "/")[1]
	runtimeDir := filepath.Join(constants.RuntimesDir, runtimeName)
	var command string
	var ext string

	switch runtime.GOOS {
	case "linux", "darwin":
		command = "sh"
		ext = "sh"
	case "windows":
		command = "powershell"
		ext = "ps1"
	}

	_, err := os.Stat(runtimeDir)
	if IsCached(runtimeName) && err == nil {
		cmd := exec.Command(command, filepath.Join(runtimeDir, fmt.Sprintf("run.%v", ext)+" "+strings.Join(args, " ")))
		cmd.Dir = runtimeDir
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		cmd.Run()

		return
	}

	for _, f := range files {
		grab.Get(filepath.Join(runtimeDir, f.Name), f.DownloadUrl)
	}

	for _, f := range files {
		cmd := exec.Command(command, filepath.Join(runtimeDir, f.Name))
		cmd.Dir = runtimeDir
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if strings.Split(f.Name, ".")[0] == "run" {
			cmd.Args = append(cmd.Args, args...)
		}

		cmd.Run()
	}

	if !IsCached(runtimeName) {
		os.RemoveAll(runtimeDir)
	}
}
