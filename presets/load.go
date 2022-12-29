package presets

import (
	"fmt"
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

// UncheckedLoad loads preset to .env file.
//
// An error is returned if preset does not exist.
// The current preset is updated on successful load.
//
// NOTE: contents of .env file will not be committed by this function. Explicitly call Commit before using this function to prevent data loss.
func UncheckedLoad(cfg config.Config, afs afero.Fs, preset string) error {
	if !Exists(cfg, afs, preset) {
		return fmt.Errorf("env preset '%s' does not exist", preset)
	}

	// get preset file handle
	presetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(preset))
	presetFile, err := fs.Open(afs, presetPath, fs.FlagRead)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// get .env file handle
	envFile, err := fs.Open(afs, ".env", fs.FlagWrite)
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
