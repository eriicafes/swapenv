package cmd

import (
	"fmt"
	"os"

	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swapenv",
	Short: "Manage .env files locally",
	Long:  `Manage .env files locally. Create and swap between existing .env files`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// ensure project has not been initialized before running `init` command
		if cmd.Name() == "init" {
			if config.Loaded() {
				exitOnError(fmt.Errorf("swapenv has already been initialized"))
			}
			return
		} else {
			// ensure project has been initialized before running all other commands
			if !config.Loaded() {
				exitOnError(fmt.Errorf("swapenv has not been initialized, run `swapenv init`"))
			}
		}
	},
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
