package cmd

import (
	"fmt"
	"runtimer/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "Runs a runtime listed in config file",
	Run: run,
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	config, err := utils.GetLocalConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	runtime := config.Runtimes[0]
	if len(config.Runtimes) > 1 {
		primary := utils.GetPrimaryRuntime(config)
		runtime = config.Runtimes[primary]

		if args[0] == "-r" {
			chosen := utils.GetRuntimeByName(args[1])
			runtime = config.Runtimes[chosen]
			args = args[2:]
		}
	}

	utils.ExecuteRuntime(runtime, args)
}