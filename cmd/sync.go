package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// Manually sync current env preset with .env file
// eg. `swapenv sync`

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Manually sync current env preset with .env file",
	Long: `Manually sync current env preset with .env file.
This command is not required to be called as syncing is always done before changing env presets`,
	Example: "swapenv sync",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := config.Get()
		afs := afero.NewOsFs()

		// commit .env contents into current preset
		err := presets.Commit(cfg, afs)
		cobra.CheckErr(err)

		fmt.Println("synced:", cfg.GetPreset(), "with .env file")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
