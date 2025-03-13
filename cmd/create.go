package cmd

import (
	"fmt"
	"io/fs"

	"github.com/eriicafes/swapenv/args"
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/spf13/cobra"
)

type CreateArgs struct {
	Name string
}

var createArgs = args.New(func(cmd *cobra.Command, rawArgs []string) (args CreateArgs, err error) {
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}
	args.Name = rawArgs[0]
	return
})

var createFlags struct {
	Base string
	Use  bool
}

var createCmd = &cobra.Command{
	Use:     "create env",
	Short:   "Create new env preset",
	Example: "swapenv create prod -u -b=staging",
	Args:    createArgs.Validate,
	Run: createArgs.Run(func(cmd *cobra.Command, args CreateArgs) {
		fsys := wfs.OS()
		wdir, envsdir := config.Dir()
		current, _ := config.GetEnv()

		if !fs.ValidPath(args.Name) {
			exitOnError(fmt.Errorf("invalid: '%s' is not a valid branch name", args.Name))
		}

		var src env.Env
		if createFlags.Base == "" {
			// set base so it can be printed
			createFlags.Base = current
			// use .env contents because the current env file is lazily updated
			src = env.DotEnv(wdir)
		} else {
			src = env.Env{Dir: envsdir, Name: createFlags.Base}
		}
		err := env.Copy(fsys, src, env.Env{Dir: envsdir, Name: args.Name})
		exitOnError(err)

		if createFlags.Use {
			detachedMsg, err := UseEnv(args.Name)
			exitOnError(err)
			fmt.Printf("Using%s env '%s' from '%s'\n:", detachedMsg, args.Name, createFlags.Base)
		} else {
			fmt.Printf("Created env '%s' from '%s'\n:", args.Name, createFlags.Base)
		}
	}),
}

func init() {
	createCmd.Flags().BoolVarP(&createFlags.Use, "use", "u", false, "use created env")
	createCmd.Flags().StringVarP(&createFlags.Base, "base", "b", "", "base env")

	rootCmd.AddCommand(createCmd)
}
