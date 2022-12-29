package presets

import (
	"os"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestExists(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	files := []string{"temp", "next"}

	// populate fs
	for _, file := range files {
		afero.WriteFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+file), []byte("testdata"), os.ModePerm)
	}

	// check preset
	if !Exists(cfg, afs, "next") {
		t.Fatal("preset should exist")
	}
	if Exists(cfg, afs, "last") {
		t.Fatal("preset should not exist")
	}
}
