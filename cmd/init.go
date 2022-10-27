package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Init project with base env file
// eg. `swapenv init --dir envs`

// flags
type InitFlags struct {
	Dir string
}

var initFlags = InitFlags{}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init project with base env file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {

		fmt.Println("creating base env file in", initFlags.Dir)
	},
}

func init() {
	initCmd.Flags().StringVar(&initFlags.Dir, "dir", "", "directory to store env files")

	rootCmd.AddCommand(initCmd)
}
