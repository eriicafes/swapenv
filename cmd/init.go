package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Init project with base env file
// eg. `swapenv init --dir envs`

// flags
var (
	dir string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init project with base env file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {

		fmt.Println("creating base env file")
	},
}

func init() {
	initCmd.Flags().StringVar(&dir, "dir", "", "directory to store env files")

	rootCmd.AddCommand(initCmd)
}
