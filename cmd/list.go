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
	Long: `
	List all env presets by default this program will look for env presets in the 'envs' directory
	on the same directory it was called
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("listing envs")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
