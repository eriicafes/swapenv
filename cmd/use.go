package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// Swap env to an existing env preset
// eg. `swapenv use staging`

// args
type UseArgs struct {
	Preset string
}

var useArgs = args.New(func(cmd *cobra.Command, rawArgs []string) (args UseArgs, err error) {
	// accept only preset argument
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	args.Preset = rawArgs[0]

	return
})

var useCmd = &cobra.Command{
	Use:     "use preset",
	Short:   "Swap env to an existing env preset",
	Example: "swapenv use staging",
	Args:    useArgs.Validate,
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := config.Get()
		afs := afero.NewOsFs()

		// swap to selected preset
		err := presets.Swap(cfg, afs, useArgs.Fields.Preset)
		cobra.CheckErr(err)

		fmt.Println("using env preset:", useArgs.Fields.Preset)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
