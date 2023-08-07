package cmd

import (
	"fmt"
	"runtimer/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                "run",
	Short:              "Runs chosen runtime",
	DisableFlagParsing: true,
	Run:                run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("List of runtimes:")
		runtimes, err := utils.GetRuntimes()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, r := range runtimes {
			fmt.Println(r.Name)
		}
		return
	}

	runtime, err := utils.GetRuntime(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	runtime.Execute(args[1:])
}
