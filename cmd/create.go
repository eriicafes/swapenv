package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/utils"
	"github.com/spf13/cobra"
)

// Create new env preset
// eg. `swapenv create name --from preset`

// flags
var (
	from string
)

// args
type CreateArgs struct {
	Name string
}

// parse args
var createArgs = utils.ParseArgs(func(cmd *cobra.Command, rawArgs []string) (args CreateArgs, err error) {
	// accept only name argument
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	// set fields
	args.Name = rawArgs[0]

	return
})

var createCmd = &cobra.Command{
	Use:   "create name",
	Short: "Create new env preset",
	Long: `
	Create new env preset. Can optionally pass a named argument '--from' to specify base preset
	`,
	Args: createArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {

		fmt.Println("create env", createArgs.Fields.Name, "from", from)
	},
}

func init() {
	// declare flags
	createCmd.Flags().StringVarP(&from, "from", "f", "", "base env preset")

	rootCmd.AddCommand(createCmd)
}
