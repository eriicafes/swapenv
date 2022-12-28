package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/spf13/afero"
)

// Exists checks if preset exists.
func Exists(cfg config.Config, afs afero.Fs, preset string) bool {
	files := List(cfg, afs)

	for _, file := range files {
		if file == preset {
			return true
		}
	}

	return false
}
