package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

// Show current env preset
// eg. `swapenv show`

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show current env preset",
	Example: "swapenv show",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		// ensure init has been run previously
		if err := config.EnsureHasInitialized(); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println(config.Env.Preset)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
