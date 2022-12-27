package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/presets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Init project with base env preset
// eg. `swapenv init -p test`

// flags
type InitFlags struct {
	Yes    bool
	Preset string
	Dir    string
}

var initFlags InitFlags

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Init project with base env preset",
	Example: "swapenv init -p test",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := config.Get()

		// change config dir if provided
		if initFlags.Dir != "" {
			err := cfg.ChDir(initFlags.Dir)
			cobra.CheckErr(err)

			fmt.Println("created swapenv config in:", config.Getwd())
		}

		// confirm init with prompt
		err := promptContinueInit(&initFlags, cfg.Dir())
		cobra.CheckErr(err)

		// get base preset with prompt
		preset, err := promptPresetName(&initFlags)
		cobra.CheckErr(err)

		// create base env preset
		// uses contents of .env file, creating it if it does not exist
		err = presets.Create(cfg, preset)
		// copy over preset env file to .env if preset already exists and .env does not exist or is empty
		var perr *presets.PresetAlreadyExists
		if errors.As(err, &perr) {
			err = copyExistingPreset(cfg, preset)
		}
		cobra.CheckErr(err)

		// TODO: add env dir and .env to .gitignore if project is a git repository

		fmt.Println("created base env preset:", ".env."+preset, "in:", cfg.Dir())

		// update env preset
		// calling this function will also create the swapenvcache config file
		err = presets.UncheckedSet(cfg, preset)
		cobra.CheckErr(err)
	},
}

func copyExistingPreset(cfg config.Config, preset string) error {
	// if .env does not exist or is empty, load preset into .env
	if stat, staterr := os.Stat(".env"); errors.Is(staterr, os.ErrNotExist) || stat.Size() == 0 {
		fmt.Printf("env preset '%s' already exists, copying contents to .env\n", preset)
		return presets.UncheckedLoad(cfg, preset)
	} else {
		fmt.Printf("env preset '%s' already exists, skipping...\n", preset)
	}
	return nil
}

func promptContinueInit(flags *InitFlags, dir string) error {
	if flags.Yes {
		return nil
	}

	message := fmt.Sprintf("Initialize .env files in the directory '%s'", dir)
	if flags.Preset != "" {
		message += fmt.Sprintf(" (using base preset '%s')", flags.Preset)
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%s. Do you want to continue", message),
		IsConfirm: true,
		Default:   "y",
	}

	_, err := prompt.Run()

	return err
}

func promptPresetName(flags *InitFlags) (string, error) {
	if flags.Preset != "" {
		return flags.Preset, nil
	}

	prompt := promptui.Prompt{
		Label:   "Name this env preset",
		Default: "local",
	}

	return prompt.Run()
}

func init() {
	initCmd.Flags().BoolVarP(&initFlags.Yes, "yes", "y", false, "skip initialize prompt")
	initCmd.Flags().StringVarP(&initFlags.Preset, "preset", "p", "", "base env preset")
	initCmd.Flags().StringVarP(&initFlags.Dir, "dir", "", "", fmt.Sprintf("swapenv config directory (default '%s')", config.DefaultWd))

	rootCmd.AddCommand(initCmd)
}
