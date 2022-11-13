package cmd

import (
	"fmt"
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
	out, err := execCmd.CombinedOutput()
	if err != nil {
		fmt.Println(execCmd.Stderr)
		return
	}
	fmt.Print(string(out))
}