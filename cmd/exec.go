package cmd

import (
	"fmt"
	"runtimer/constants"
	"runtimer/utils"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: "exec",
	Short: "Runs a standalone runtime",
	Run: execCli,
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func execCli(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Print("Provide a runtime")
		return
	}
	run, _ := constants.GetDefinedRuntime(args[0])
	utils.ExecuteRuntime(run, true, false, args)
}