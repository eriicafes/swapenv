package presets

import (
	"github.com/eriicafes/swapenv/config"
)

// Exists checks if preset exists.
func Exists(cfg config.Config, preset string) bool {
	files := List(cfg)

	for _, file := range files {
		if file == preset {
			return true
		}
	}

	return false
}
