package presets

import (
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Commit .env file to current preset.
func Commit() error {
	// get env file handle
	envFile, err := fs.OpenFileRead(".env")
	if err != nil {
		return err
	}
	defer envFile.Close()

	// get preset file handle
	presetPath := path.Join(config.Base, fs.PathFromFormattedName(config.Env.Preset))
	presetFile, err := fs.OpenFileWrite(presetPath)
	if err != nil {
		return err
	}
	defer presetFile.Close()

	// copy env into preset
	_, err = io.Copy(presetFile, envFile)

	return err
}
