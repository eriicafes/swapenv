package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/utils"
	"github.com/spf13/cobra"
)

// Swap env to existing preset
// eg. `swapenv to preset`

// args
type ToArgs struct {
	Preset string
}

// parse args
var toArgs = utils.ParseArgs(func(cmd *cobra.Command, rawArgs []string) (args ToArgs, err error) {
	// optionally accept a preset argument
	if err = cobra.MaximumNArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}

	// set fields
	if len(rawArgs) == 1 {
		args.Preset = rawArgs[0]
	}

	return
})

var toCmd = &cobra.Command{
	Use:   "to [preset]",
	Short: "Swap env to existing preset",
	Long: `
	Swap env to existing preset
	`,
	Args: toArgs.Validate,
	Run: func(cmd *cobra.Command, args []string) {
		if toArgs.Fields.Preset == "" {
			fmt.Println("listing and switching to preset")
			return
		}

		fmt.Println("switching to preset", toArgs.Fields.Preset)
	},
}

func init() {
	rootCmd.AddCommand(toCmd)
}
