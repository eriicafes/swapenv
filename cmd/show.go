package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/spf13/cobra"
)

type ShowArgs struct {
	Name string
}

var showArgs = args.New(func(cmd *cobra.Command, rawArgs []string) (args ShowArgs, err error) {
	if err = cobra.MaximumNArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}
	if len(rawArgs) == 1 {
		args.Name = rawArgs[0]
	}
	return
})

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show the current env",
	Long:    "Show the current env. If env is not provided it will print the name of the current env",
	Example: "swapenv show staging",
	Args:    showArgs.Validate,
	Run: showArgs.Run(func(cmd *cobra.Command, args ShowArgs) {
		fsys := wfs.OS()
		current, _ := config.GetEnv()
		_, envsdir := config.Dir()

		if args.Name != "" {
			// show env contents
			namedEnv := env.Env{Dir: envsdir, Name: args.Name}
			file, err := namedEnv.Open(fsys)
			exitOnError(err)
			defer file.Close()
			io.Copy(os.Stdout, file)
		} else {
			fmt.Printf("On env '%s'\n", current)
		}
	}),
}

func init() {
	rootCmd.AddCommand(showCmd)
}
