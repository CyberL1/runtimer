package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtimer/constants"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Displays your CLI version",
	Run: version,
}
var updateCmd = &cobra.Command{
	Use: "update",
	Short: "Updates your CLI version",
	Run: update,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.AddCommand(updateCmd)
}

func version(cmd *cobra.Command, args []string) {
	latestRelease, _ := constants.GetLatestCliVersion()
	if constants.Version < latestRelease.TagName {
		fmt.Println("A new update is avaliable")
		fmt.Println("Run 'runtimer version update' to update now")
	}
	fmt.Println("Your CLI Version:", constants.Version)
	fmt.Println("Latest CLI version:", latestRelease.TagName)
}

func update(cmd *cobra.Command, args []string) {
	var command string
	var cmdArgs []string
	switch runtime.GOOS {
	case "linux", "darwin":
		command = "curl -fsSL https://raw.githubusercontent.com/CyberL1/runtimer/master/scripts/get.sh | sh"
	case "windows":
		command = "powershell"
		cmdArgs = []string{"irm https://raw.githubusercontent.com/CyberL1/runtimer/master/scripts/get.ps1 | iex"}
	}
	execCmd := exec.Command(command, cmdArgs...)
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Run()
}