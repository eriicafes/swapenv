package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

// Swap env to an existing env preset
// eg. `swapenv use staging`

// args
type UseArgs struct {
	Preset string
}

var useArgs = args.NewArgs(func(cmd *cobra.Command, rawArgs []string) (args UseArgs, err error) {
	// accept only preset argument
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	args.Preset = rawArgs[0]

	return
})

var useCmd = &cobra.Command{
	Use:   "use preset",
	Short: "Swap env to an existing env preset",
	Args:  useArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {
		// ensure preset exists

		// write .env contents into current preset

		// load target env preset into .env

		fmt.Println("switching to preset", useArgs.Fields.Preset)

		// update env preset
		if err := config.SetEnvPreset(useArgs.Fields.Preset); err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
