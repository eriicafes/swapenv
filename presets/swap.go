package presets

import (
	"errors"
	"os"

	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/afero"
)

// Swap commits .env file to current preset and loads provided preset.
//
// Error due to missing .env file is ignored during swap.
func Swap(cfg config.Config, afs afero.Fs, preset string) error {
	// commit current preset
	if err := Commit(cfg, afs); err != nil {
		// ignore error if error is .env file does not exist
		// if .env file does not exists we will go ahead and create it with the contents of the target preset
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	// load preset
	return UncheckedLoad(cfg, afs, preset)
}
