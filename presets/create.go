package presets

import (
	"io"
	"path"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Create preset using contents of .env file.
//
// The .env file is created if it does not exist.
// An error is returned if preset already exists.
func Create(cfg config.Config, preset string) error {
	// return error if preset already exists
	if Exists(cfg, preset) {
		return &PresetAlreadyExists{preset: preset}
	}

	// get .env file handle
	envFile, err := fs.OpenFileReadCreate(".env")
	if err != nil {
		return err
	}
	defer envFile.Close()

	// get preset file handle
	presetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(preset))
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
//
// An error is returned if preset already exists or if base preset does not exist.
func CreateFrom(cfg config.Config, preset string, base string) error {
	// return error if preset already exists
	if Exists(cfg, preset) {
		return &PresetAlreadyExists{preset: preset}
	}

	// return error if base preset does not exist
	if !Exists(cfg, base) {
		return &PresetDoesNotExist{preset: base, base: true}
	}

	// get src preset file handle
	srcPresetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(base))
	srcPresetFile, err := fs.OpenFileRead(srcPresetPath)
	if err != nil {
		return err
	}

	// get dest preset file handle
	destPresetPath := path.Join(cfg.Dir(), fs.PathFromFormattedName(preset))
	destPresetFile, err := fs.OpenFileWrite(destPresetPath)
	if err != nil {
		return err
	}

	// copy src preset into dest preset
	_, err = io.Copy(destPresetFile, srcPresetFile)

	return err
}
