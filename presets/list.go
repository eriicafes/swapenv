package presets

import (
	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

// List returns all presets.
func List(cfg config.Config, afs afero.Fs) []string {
	files, _ := fs.List(afs, cfg.Dir())

	return files
}
