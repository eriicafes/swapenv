package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/spf13/afero"
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

var createArgs = args.New(func(cmd *cobra.Command, rawArgs []string) (args CreateArgs, err error) {
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
		cfg := config.Get()
		afs := afero.NewOsFs()

		// create env preset
		if createFlags.Base != "" {
			// create preset from base env preset
			err := presets.CreateFrom(cfg, afs, createArgs.Fields.Preset, createFlags.Base)
			cobra.CheckErr(err)

			fmt.Println("created env preset:", createArgs.Fields.Preset, "from:", createFlags.Base)
		} else {
			// create preset from .env file, creating it if it does not exist
			err := presets.Create(cfg, afs, createArgs.Fields.Preset)
			cobra.CheckErr(err)

			fmt.Println("created env preset:", createArgs.Fields.Preset)
		}

		// proceed to use preset if use flag was provided
		if !createFlags.Use {
			return
		}

		// swap to newly created preset
		err := presets.Swap(cfg, afs, createArgs.Fields.Preset)
		cobra.CheckErr(err)

		fmt.Println("using env preset:", createArgs.Fields.Preset)
	},
}

func init() {
	createCmd.Flags().StringVarP(&createFlags.Base, "base", "b", "", "optional base env preset")
	createCmd.Flags().BoolVarP(&createFlags.Use, "use", "u", false, "use created env preset")

	rootCmd.AddCommand(createCmd)
}
