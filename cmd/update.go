package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: "update",
	Short: "Update runtimer",
	Run: update,
}

func init() {
	rootCmd.AddCommand(updateCmd)
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