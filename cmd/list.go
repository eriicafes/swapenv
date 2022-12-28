package cmd

import (
	"fmt"
	"strings"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// List all env presets
// eg. `swapenv ls -i`

type ListFlags struct {
	Interactive bool
}

var listFlags ListFlags

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all env presets",
	Example: "swapenv ls -i",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := config.Get()
		afs := afero.NewOsFs()
		envs := presets.List(cfg, afs)

		// list all env presets in non-interactive mode
		if !listFlags.Interactive {
			listEnvs(envs, cfg.GetPreset())
			return
		}

		// select env with prompt
		preset, err := promptSelectEnv(envs)
		cobra.CheckErr(err)

		// swap to selected preset
		err = presets.Swap(cfg, afs, preset)
		cobra.CheckErr(err)

		fmt.Println("using env preset:", preset)
	},
}

func coloured(s string) string {
	return "\033[36m" + s + "\033[0m"
}

func listEnvs(envs []string, preset string) {
	fenvs := make([]string, 0, len(envs))

	for _, env := range envs {
		if env == preset {
			env = coloured("* " + env)
		} else {
			env = "  " + env
		}
		fenvs = append(fenvs, env)
	}

	fmt.Println(strings.Join(fenvs, "\n"))
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
