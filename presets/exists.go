package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Check if preset exists.
func Exists(preset string) bool {
	files, _ := fs.List(config.Base)

	for _, file := range files {
		if file == preset {
			return true
		}
	}

	return false
}
