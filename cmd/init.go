package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Init project with base env preset
// eg. `swapenv init -p test`

// flags
type InitFlags struct {
	Preset string
}

var initFlags InitFlags

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Init project with base env preset",
	Example: "swapenv init -p test",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		// check if init has been run previously
		if config.LoadedFromFile {
			fmt.Println("swapenv has already been initialized on this project")
			return
		}

		// confirm init and get base preset with prompt
		preset, err := promptInit(initFlags.Preset)
		if err != nil {
			fmt.Println(err)
			return
		}

		// create base env preset
		// uses contents of .env file, creating it if it does not exist
		if err = presets.Create(preset); err != nil {
			fmt.Println(err)
			return
		}

		// TODO: add env dir and .env to .gitignore if project is a git repository

		fmt.Println("created base env preset:", ".env."+preset, "in:", config.Base)

		// update env preset
		// we call this function directly because we do not want to repopulate the .env file (presets.LoadUnchecked does this) which already serves as the source of truth
		// calling this function will also create the swapenvcache config file
		if err = config.SetEnvPreset(preset); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func promptInit(preset string) (string, error) {
	message := fmt.Sprintf("Initialize .env files in the directory './%s'", config.Base)
	if preset != "" {
		message += fmt.Sprintf(" (using base preset '%s')", preset)
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%s. Do you want to continue", message),
		IsConfirm: true,
		Default:   "y",
	}

	if _, err := prompt.Run(); err != nil {
		return "", err
	}

	if preset != "" {
		return preset, nil
	}

	prompt = promptui.Prompt{
		Label:   "Name this env preset",
		Default: "local",
	}

	return prompt.Run()
}

func init() {
	initCmd.Flags().StringVarP(&initFlags.Preset, "preset", "p", "", "base env preset")

	rootCmd.AddCommand(initCmd)
}
