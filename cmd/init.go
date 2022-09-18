package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create base .env file
// eg. `swapenv init`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create base .env file",
	Long: `
	Create base .env file
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("creating .env")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
