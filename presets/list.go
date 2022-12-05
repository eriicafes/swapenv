package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// List all presets.
func List() []string {
	cfg := config.Get()

	files, _ := fs.List(cfg.Base())

	return files
}
