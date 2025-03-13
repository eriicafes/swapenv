package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var listFlags struct {
	Interactive bool
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all envs",
	Example: "swapenv ls -i",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		fsys := wfs.OS()
		_, envsdir := config.Dir()
		current, _ := config.GetEnv()

		envs := env.List(fsys, envsdir)
		// list all envs in non-interactive mode
		if !listFlags.Interactive {
			for _, e := range envs {
				if e == current {
					fmt.Println("\033[36m" + "* " + e + "\033[0m")
				} else {
					fmt.Println("  " + e)
				}
			}
			return
		}

		// select env with prompt
		prompt := promptui.Select{
			Label: "Select env to use",
			Items: envs,
			Size:  10,
		}
		_, val, err := prompt.Run()
		exitOnError(err)
		if val == current {
			return
		}

		detachedMsg, err := UseEnv(val)
		exitOnError(err)

		fmt.Printf("Using%s env '%s'\n", detachedMsg, val)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listFlags.Interactive, "interactive", "i", false, "list envs in interactive mode")

	rootCmd.AddCommand(listCmd)
}
