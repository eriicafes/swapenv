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

type UseArgs struct {
	Name string
}

var useArgs = args.New(func(cmd *cobra.Command, rawArgs []string) (args UseArgs, err error) {
	if err = cobra.ExactArgs(1)(cmd, rawArgs); err != nil {
		return args, err
	}
	args.Name = rawArgs[0]
	return
})

var useFlags struct {
	Base bool
}

var useCmd = &cobra.Command{
	Use:     "use env",
	Short:   "Switch env",
	Example: "swapenv use staging -b",
	Args:    useArgs.Validate,
	Run: useArgs.Run(func(cmd *cobra.Command, args UseArgs) {
		fsys := wfs.OS()
		current, prev := config.GetEnv()
		wdir, envsdir := config.Dir()

		// alias "-" to the previous env
		if args.Name == "-" {
			args.Name = prev
		}

		if args.Name == current {
			return
		}

		var created bool
		if useFlags.Base {
			// create env if it does not exist
			newEnv := env.Env{Dir: envsdir, Name: args.Name}
			if _, err := fs.Stat(fsys, newEnv.Path()); err != nil {
				err := env.Copy(fsys, env.DotEnv(wdir), newEnv)
				exitOnError(err)
				created = true
			}
		}

		detachedMsg, err := UseEnv(args.Name)
		exitOnError(err)

		if created {
			fmt.Printf("Using%s env '%s' from '%s'\n", detachedMsg, args.Name, current)
		} else {
			fmt.Printf("Using%s env '%s'\n", detachedMsg, args.Name)
		}
	}),
}

func UseEnv(name string) (string, error) {
	fsys := wfs.OS()
	wdir, envsdir := config.Dir()
	current, _ := config.GetEnv()
	branch, err := config.GitBranch()
	exitOnError(err)

	// save current env
	err = env.Copy(fsys, env.DotEnv(wdir), env.Env{Dir: envsdir, Name: current})
	if err != nil {
		return "", err
	}
	// load new env
	err = env.Copy(fsys, env.Env{Dir: envsdir, Name: name}, env.DotEnv(wdir))
	if err != nil {
		return "", err
	}
	// update env
	err = config.SetEnv(name)
	var detachedMsg string
	if name != branch {
		detachedMsg = " DETACHED"
	}
	return detachedMsg, err
}

func init() {
	useCmd.Flags().BoolVarP(&useFlags.Base, "base", "b", false, "create env if it does not exist using current env as base")

	rootCmd.AddCommand(useCmd)
}
