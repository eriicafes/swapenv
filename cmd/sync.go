package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

// Manually sync current env preset with .env file
// eg. `swapenv sync`

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Manually sync current env preset with .env file",
	Long: `Manually sync current env preset with .env file.
This command is not required to be called as syncing is always done before changing env presets`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		// write .env contents into current preset

		fmt.Println("synced", config.Env.Preset, "with .env file")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
