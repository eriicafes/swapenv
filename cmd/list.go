package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// List all env presets
// eg. `swapenv list`

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all env presets",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {

		fmt.Println("listing envs")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
