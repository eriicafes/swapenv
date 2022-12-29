package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/eriicafes/swapenv/presets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/afero"
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
		afs := afero.NewOsFs()

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
		err = presets.Create(cfg, afs, preset)
		// copy over preset env file to .env if preset already exists and .env does not exist or is empty
		var perr *presets.PresetAlreadyExists
		if errors.As(err, &perr) {
			err = copyExistingPreset(cfg, afs, preset)
		}
		cobra.CheckErr(err)

		// add .env and config dir to .gitignore if project is a git repository
		if _, ok := config.GitRoot(); ok {
			ignorePaths := []string{
				".env",
				"/" + filepath.Base(cfg.Dir()),
				config.ConfigRootName,
				config.ConfigName,
			}
			// add ignore paths to gitignore
			err := addToGitIgnore(cfg, afs, ignorePaths)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("created base env preset:", ".env."+preset, "in:", cfg.Dir())

		// update env preset
		// calling this function will also create the config file
		err = presets.UncheckedSet(cfg, preset)
		cobra.CheckErr(err)
	},
}

func copyExistingPreset(cfg config.Config, afs afero.Fs, preset string) error {
	// if .env does not exist or is empty, load preset into .env
	if stat, staterr := afs.Stat(".env"); errors.Is(staterr, os.ErrNotExist) || stat.Size() == 0 {
		fmt.Printf("env preset '%s' already exists, copying contents to .env\n", preset)
		return presets.UncheckedLoad(cfg, afs, preset)
	} else {
		fmt.Printf("env preset '%s' already exists, skipping...\n", preset)
	}
	return nil
}

func addToGitIgnore(cfg config.Config, afs afero.Fs, paths []string) error {
	// open .gitignore file
	file, err := fs.Open(afs, ".gitignore", fs.FlagReadWriteAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// track if .gitignore has contents
	var hasContents bool

	// filter ignore paths
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(paths) > 0 { // stop once paths is empty
		for i, path := range paths {
			line := scanner.Text()
			// set hasContents to true once a non empty line is read from scanner
			if !hasContents && len(line) > 0 {
				hasContents = true
			}
			if path == line {
				paths = append(paths[:i], paths[i+1:]...) // filter out path
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return err
	}

	// return early if ignore paths is now empty
	if len(paths) == 0 {
		return nil
	}

	// write ignore paths to .gitignore
	// append new line if .gitignore has contents
	var content string
	if hasContents {
		content += "\n"
	}
	for _, path := range paths {
		content += path + "\n"
	}
	_, err = io.WriteString(file, content)
	return err
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
	initCmd.Flags().StringVarP(&initFlags.Dir, "dir", "", "", fmt.Sprintf("swapenv config directory (default '%s')", config.DefaultDir))

	rootCmd.AddCommand(initCmd)
}
