package cmd

import (
	"fmt"
	"runtimer/cache"

	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use: "cache",
	Short: "Caches a runtime",
	Run: cacheCli,
}

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.Flags().BoolP("remove", "r", false, "Remove from cache")
}

func cacheCli(cmd *cobra.Command, args []string) {
	cached := cache.Get()

	if len(args) == 0 {
		if len(cached) == 0 {
			fmt.Print("Nothing is cached")
			return
		}
		for item, state := range cached {
			var stateString string
			if state {
				stateString = "cached"
			} else {
				stateString = "not cached"
			}
			fmt.Printf("%s is %s\n", item, stateString)
		}
		return
	}
	remove := cmd.Flag("remove")
	var removedString string
	if remove.Changed {
		removedString = "false"
	} else {
		removedString = "true"
	}
	cache.Set(args[0], !remove.Changed)
	fmt.Printf("%s cache set to %s", args[0], removedString)
}