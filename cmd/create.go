package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/spf13/cobra"
)

// Create new env preset
// eg. `swapenv create staging -u -b prod`

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
	Use:     "create preset",
	Short:   "Create new env preset",
	Example: "swapenv create staging -u -b prod",
	Args:    createArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {
		// ensure init has been run previously
		if err := config.EnsureHasInitialized(); err != nil {
			cobra.CheckErr(err)
		}

		// create env preset
		var err error
		if createFlags.Base != "" {
			// create preset from base
			err = presets.CreateFrom(createArgs.Fields.Preset, createFlags.Base)
		} else {
			// create preset from .env file, creating it if it does not exist
			err = presets.Create(createArgs.Fields.Preset)
		}
		cobra.CheckErr(err)

		fmt.Println("created env preset:", createArgs.Fields.Preset, "from:", createFlags.Base)

		// proceed to use preset if use flag was provided
		if !createFlags.Use {
			return
		}

		// swap to newly created preset
		if err = presets.Swap(createArgs.Fields.Preset); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println("using env preset:", createArgs.Fields.Preset)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createFlags.Base, "base", "b", "", "optional env preset to base env file")
	createCmd.Flags().BoolVarP(&createFlags.Use, "use", "u", false, "use created env preset")

	rootCmd.AddCommand(createCmd)
}
