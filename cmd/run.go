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

	runtime := &config.Runtimes[0]
	multiple := false
	if len(config.Runtimes) > 1 {
		multiple = true
		runtime, err := utils.GetPrimaryRuntime(config)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 0 {
			utils.ExecuteRuntime(runtime, false, multiple, args)
			return
		}

		if args[0] == "-r" {
			runtime, err = utils.GetCustomRuntimeByName(args[1])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	utils.ExecuteRuntime(runtime, false, multiple, args)
}