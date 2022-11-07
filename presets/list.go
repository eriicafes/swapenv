package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
)

// List all presets.
func List() []string {
	files, _ := fs.List(config.Base)

	return files
}
