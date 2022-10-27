package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/spf13/cobra"
)

// Create new env preset
// eg. `swapenv create test -b prod`

// flags
type CreateFlags struct {
	Base string
}

var createFlags = CreateFlags{}

// args
type CreateArgs struct {
	Name string
}

var createArgs = args.NewArgs(func(cmd *cobra.Command, rawArgs []string) (args CreateArgs, err error) {
	// accept only name argument
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	args.Name = rawArgs[0]

	return
})

var createCmd = &cobra.Command{
	Use:   "create [preset]",
	Short: "Create new env preset",
	Args:  createArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {

		fmt.Println("create env", createArgs.Fields.Name, "from", createFlags.Base)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createFlags.Base, "base", "b", "", "optional env preset to base env file")

	rootCmd.AddCommand(createCmd)
}
