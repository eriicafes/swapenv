package cmd

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Init project with base env file
// eg. `swapenv init -p test`

// flags
type InitFlags struct {
	Preset string
}

var initFlags InitFlags

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init project with base env preset",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		// check if init has been run previously
		if config.LoadedFromFile {
			fmt.Println("swapenv has already been initialized on this project")
			return
		}

		// confirm init and get base env preset with prompt
		preset, err := promptInit(initFlags.Preset)
		if err != nil {
			fmt.Println(err)
			return
		}

		// create base env preset

		// load newly created env preset into .env

		// create swapenvcache config store file

		// add env dir and .env to .gitignore if project is a git repository

		fmt.Println("created base env preset:", ".env."+preset, "in: envs")

		// update env preset
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
