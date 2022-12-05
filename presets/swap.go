package presets

import (
	"errors"
	"os"
)

// Commit .env file to current preset and load preset to .env file.
func Swap(preset string) error {
	// commit current preset
	if err := Commit(); err != nil {
		// ignore error if error is .env file does not exist
		// if .env file does not exists we will go ahead and create it with the contents of the target preset
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	// load preset
	return LoadUnchecked(preset)
}
