package presets

import (
	"fmt"
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// UncheckedLoad loads preset to .env file.
//
// An error is returned if preset does not exist.
// The current preset is updated on successful load.
//
// NOTE: contents of .env file will not be committed by this function. Explicitly call Commit before using this function to prevent data loss.
func UncheckedLoad(cfg config.Config, preset string) error {
	if !Exists(cfg, preset) {
		return fmt.Errorf("env preset '%s' does not exist", preset)
	}

	// get preset file handle
	presetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(preset))
	presetFile, err := fs.OpenFileRead(presetPath)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// get .env file handle
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

	// update preset
	return UncheckedSet(cfg, preset)
}
