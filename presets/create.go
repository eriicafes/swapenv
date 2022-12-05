package presets

import (
	"fmt"
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Create preset using contents of .env file.
//
// The .env file is created if it does not exist.
func Create(preset string) error {
	cfg := config.Get()

	// return error if preset already exists
	if Exists(preset) {
		return fmt.Errorf("env preset '%v' already exists", preset)
	}

	// get env file handle
	envFile, err := fs.OpenFileReadCreate(".env")
	if err != nil {
		return err
	}
	defer envFile.Close()

	// get preset file handle
	presetPath := path.Join(cfg.Base(), fs.PathFromFormattedName(preset))
	presetFile, err := fs.OpenFileWrite(presetPath)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// copy env into preset
	_, err = io.Copy(presetFile, envFile)

	return err
}

// Create preset using contents base preset.
func CreateFrom(preset string, base string) error {
	cfg := config.Get()

	// return error if preset already exists
	if Exists(preset) {
		return fmt.Errorf("env preset '%v' already exists", preset)
	}

	// return error if base preset does not exist
	if !Exists(base) {
		return fmt.Errorf("base env preset '%v' does not exist", base)
	}

	// get src preset file handle
	srcPresetPath := path.Join(cfg.Base(), fs.PathFromFormattedName(base))
	srcPresetFile, err := fs.OpenFileRead(srcPresetPath)
	if err != nil {
		return err
	}

	// get dest preset file handle
	destPresetPath := path.Join(cfg.Base(), fs.PathFromFormattedName(preset))
	destPresetFile, err := fs.OpenFileWrite(destPresetPath)
	if err != nil {
		return err
	}

	// copy src preset into dest preset
	_, err = io.Copy(destPresetFile, srcPresetFile)

	return err
}
