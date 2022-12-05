package presets

import (
	"fmt"
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Load preset to .env file.
//
// NOTE: contents of .env file will not be committed by this function. Explicitly call Commit before using this function to prevent data loss.
func LoadUnchecked(preset string) error {
	cfg := config.Get()

	if !Exists(preset) {
		return fmt.Errorf("env preset '%v' does not exist", preset)
	}

	// get preset file handle
	presetPath := path.Join(cfg.Base(), fs.PathFromFormattedName(preset))
	presetFile, err := fs.OpenFileRead(presetPath)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// get env file handle
	envFile, err := fs.OpenFileWrite(".env")
	if err != nil {
		return err
	}
	defer envFile.Close()

	// copy preset into env
	_, err = io.Copy(envFile, presetFile)
	if err != nil {
		return err
	}

	// update config
	if err := cfg.SetPreset(preset); err != nil {
		return err
	}

	return nil
}
