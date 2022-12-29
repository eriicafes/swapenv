package presets

import (
	"os"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestSwap(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	initialPreset := "temp"
	nextPreset := "next"
	content := []byte("ROLE=temp")
	nextContent := []byte("ROLE=next")

	// setup config
	cfg.SetPreset(initialPreset)

	// populate fs
	afero.WriteFile(afs, ".env", content, os.ModePerm)
	afero.WriteFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+nextPreset), nextContent, os.ModePerm)

	// swap preset
	err := Swap(cfg, afs, nextPreset)
	if err != nil {
		t.Fatal(err)
	}

	// check preset
	gotPreset := cfg.GetPreset()
	if nextPreset != gotPreset {
		t.Fatalf("set error, expected %s got %s", nextPreset, gotPreset)
	}
	// read preset
	got, err := afero.ReadFile(afs, ".env")
	if err != nil {
		t.Error(err)
	}
	err = compareRead("swap", nextContent, got)
	if err != nil {
		t.Fatal(err)
	}
	// confirm initial preset contents now contains .env previous contents
	got, err = afero.ReadFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+initialPreset))
	if err != nil {
		t.Error(err)
	}
	err = compareRead("swap", content, got)
	if err != nil {
		t.Fatal(err)
	}
}
