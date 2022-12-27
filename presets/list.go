package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// List returns all presets.
func List(cfg config.Config) []string {
	files, _ := fs.List(cfg.Dir())

	return files
}
