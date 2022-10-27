package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swapenv",
	Short: "Manage env files locally",
	Long:  `Manage env files locally. Create and Swap between existing env presets`,
}

func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
