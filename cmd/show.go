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
		cfg := config.Get()

		fmt.Println(cfg.GetPreset())
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
