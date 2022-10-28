package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/cobra"
)

// Create new env preset
// eg. `swapenv create test -u -b prod`

// flags
type CreateFlags struct {
	Base string
	Use  bool
}

var createFlags CreateFlags

// args
type CreateArgs struct {
	Preset string
}

var createArgs = args.NewArgs(func(cmd *cobra.Command, rawArgs []string) (args CreateArgs, err error) {
	// accept only preset argument
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	args.Preset = rawArgs[0]

	return
})

var createCmd = &cobra.Command{
	Use:   "create preset",
	Short: "Create new env preset",
	Args:  createArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {
		// ensure base preset exists if provided

		// create env preset (using optional base preset if provided), return error if it already exists

		fmt.Println("created env", createArgs.Fields.Preset, "from", createFlags.Base)

		// proceed to use env preset if use flag was provided
		if !createFlags.Use {
			return
		}

		// write .env contents into current preset

		// load newly created env preset into .env

		fmt.Println("using", createArgs.Fields.Preset)

		// update env preset
		if err := config.SetEnvPreset(createArgs.Fields.Preset); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&createFlags.Base, "base", "b", "", "optional env preset to base env file")
	createCmd.Flags().BoolVarP(&createFlags.Use, "use", "u", false, "use created env preset")

	rootCmd.AddCommand(createCmd)
}
