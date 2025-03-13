package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/env"
	"github.com/eriicafes/wfs"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rmFlags struct {
	All   bool
	Force bool
}

var rmCmd = &cobra.Command{
	Use:     "rm [envs]...",
	Short:   "Remove env",
	Example: "swapenv rm staging dev prod",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fsys := wfs.OS()
		_, envsdir := config.Dir()

		if rmFlags.All && !rmFlags.Force {
			// confirm before removing all
			prompt := promptui.Prompt{
				Label: "Are you sure you want to remove all envs?",
				Validate: func(s string) error {
					if s == "y" || s == "n" {
						return nil
					}
					//lint:ignore ST1005 This error is prompted to the user
					return fmt.Errorf("Enter y/n")
				},
			}
			val, _ := prompt.Run()
			if val != "y" {
				return
			}
		}

		var envs []env.Env
		if rmFlags.All {
			for _, e := range env.List(fsys, envsdir) {
				dir, name := filepath.Split(e)
				envs = append(envs, env.Env{Dir: dir, Name: name})
			}
		} else {
			for _, arg := range args {
				envs = append(envs, env.Env{Dir: envsdir, Name: arg})
			}
		}

		removed := env.Remove(fsys, envs...)
		if len(removed) > 0 {
			fmt.Printf("Removed %d envs:\n", len(removed))
			for _, name := range removed {
				fmt.Println(name)
			}
		} else {
			fmt.Printf("Removed 0 envs\n")
		}
	},
}

func init() {
	rmCmd.Flags().BoolVarP(&rmFlags.All, "all", "i", false, "remove all envs")
	rmCmd.Flags().BoolVarP(&rmFlags.Force, "force", "f", false, "force remove")

	rootCmd.AddCommand(rmCmd)
}
