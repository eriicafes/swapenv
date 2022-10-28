package cmd

import (
	"fmt"
	"strings"

	"github.com/eriicafes/swapenv/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// List all env presets
// eg. `swapenv list`

type ListFlags struct {
	Interactive bool
}

var listFlags ListFlags

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all env presets",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		envs := []string{"dev", "test", "staging", "prod"}

		// simply list all env presets in non-interactive mode
		if !listFlags.Interactive {
			fmt.Println(strings.Join(envs, "\n"))
			return
		}

		// select env with prompt
		preset, err := promptSelectEnv(envs)
		if err != nil {
			fmt.Println(err)
			return
		}

		// write .env contents into current preset

		// load target env preset into .env

		fmt.Println("selected env preset:", preset)

		// update env preset
		if err = config.SetEnvPreset(preset); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func promptSelectEnv(envs []string) (string, error) {
	prompt := promptui.Select{
		Label: "Choose env preset to load",
		Items: envs,
		Size:  10,
	}

	_, env, err := prompt.Run()

	return env, err
}

func init() {
	listCmd.Flags().BoolVarP(&listFlags.Interactive, "interactive", "i", false, "list presets in interactive mode")

	rootCmd.AddCommand(listCmd)
}
