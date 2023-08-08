package cmd

import (
	"fmt"
	"runtimer/utils"

	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Caches a runtime",
	Run:   cache,
}

func init() {
	cacheCmd.Flags().BoolP("remove", "r", false, "Remove from cache")
	rootCmd.AddCommand(cacheCmd)
}

func cache(cmd *cobra.Command, args []string) {
	cache := utils.GetCache()

	if len(args) == 0 {
		if len(cache) == 0 {
			fmt.Println("Nothing is cached")
			return
		}

		_, err := utils.GetRuntime(args[0])
		if err != nil {
			fmt.Printf("runtime %v not found", args[0])
			return
		}

		for _, r := range cache {
			fmt.Println(r)
		}
		return
	}

	remove := cmd.Flag("remove")
	var action string
	if remove.Changed {
		if !utils.IsCached(args[0]) {
			fmt.Printf("%v is already uncached\n", args[0])
			return
		}
		action = "uncached"
	} else {
		if utils.IsCached(args[0]) {
			fmt.Printf("%v is already cached\n", args[0])
			return
		}
		action = "cached"
	}
	utils.SetCache(args[0], remove.Changed)
	fmt.Printf("%v is now %v\n", args[0], action)
}
