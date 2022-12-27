package cmd

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swapenv",
	Short: "Manage env files locally",
	Long:  `Manage env files locally. Create and Swap between existing env presets`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()

		// ensure init has not been run before running `init` command
		if cmd.Name() == "init" {
			if err := config.EnsureHasNotInit(cfg); err != nil {
				cobra.CheckErr(err)
			}
			return
		}

		// ensure init has been run before running all other commands
		if err := config.EnsureHasInit(cfg); err != nil {
			cobra.CheckErr(err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
