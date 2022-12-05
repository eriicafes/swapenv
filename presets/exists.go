package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// Check if preset exists.
func Exists(preset string) bool {
	cfg := config.Get()

	files, _ := fs.List(cfg.Base())

	for _, file := range files {
		if file == preset {
			return true
		}
	}

	return false
}
