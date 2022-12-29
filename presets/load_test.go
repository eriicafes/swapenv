package presets

import (
	"os"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestLoad(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	preset := "temp"
	nextPreset := "next"
	content := []byte("ROLE=temp")
	nextContent := []byte("ROLE=next")

	// setup config
	cfg.SetPreset(preset)

	// populate fs
	afero.WriteFile(afs, ".env", content, os.ModePerm)
	afero.WriteFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+nextPreset), nextContent, os.ModePerm)

	// load preset
	err := UncheckedLoad(cfg, afs, nextPreset)
	if err != nil {
		t.Fatal(err)
	}

	// read preset
	got, err := afero.ReadFile(afs, ".env")
	if err != nil {
		t.Error(err)
	}
	err = compareRead("load", nextContent, got)
	if err != nil {
		t.Fatal(err)
	}
}
