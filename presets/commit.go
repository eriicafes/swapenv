package presets

import (
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

// Commit commits .env file to current preset.
//
// An error is returned if .env file does not exist. However, the preset file will be created if it does not exist.
func Commit(cfg config.Config, afs afero.Fs) error {
	// get .env file handle
	envFile, err := fs.Open(afs, ".env", fs.FlagRead)
	if err != nil {
		return err
	}
	defer envFile.Close()

	// get preset file handle
	presetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(cfg.GetPreset()))
	presetFile, err := fs.Open(afs, presetPath, fs.FlagWrite)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// copy env into preset
	_, err = io.Copy(presetFile, envFile)

	return err
}
