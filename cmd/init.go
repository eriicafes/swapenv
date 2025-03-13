package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize swapenv",
	Example: "swapenv init",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		fsys := wfs.OS()
		wdir, envsdir := config.Dir()
		branch, err := config.GitBranch()
		exitOnError(err)

		// create config path
		err = fsys.MkdirAll(envsdir, 0755)
		exitOnError(err)

		// create initial env named after current branch using contents of .env file
		err = env.Copy(fsys, env.DotEnv(wdir), env.Env{Dir: envsdir, Name: branch})
		exitOnError(err)

		// install post checkout hook
		err = env.InstallHook(fsys, filepath.Dir(envsdir))
		exitOnError(err)

		// set branch as current env
		err = config.SetEnv(branch)
		exitOnError(err)

		fmt.Println("Initialized swapenv")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
