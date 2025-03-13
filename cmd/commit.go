package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Commit .env file contents to the current env",
	Long:    "Commit .env file contents to the current env. Swapenv will always commit before changing env",
	Example: "swapenv sync",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		fsys := wfs.OS()
		wdir, envsdir := config.Dir()
		current, _ := config.GetEnv()

		// save current env
		err := env.Copy(fsys, env.DotEnv(wdir), env.Env{Dir: envsdir, Name: current})
		exitOnError(err)

		fmt.Printf("Committed .env to '%s'\n", current)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
