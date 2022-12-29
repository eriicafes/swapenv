package presets

import (
	"os"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestCreateFromEnv(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	newPreset := "new"
	expected := []byte("ROLE=swap")

	// populate fs
	afero.WriteFile(afs, ".env", expected, os.ModePerm)

	// create preset
	err := Create(cfg, afs, newPreset)
	if err != nil {
		t.Fatal(err)
	}

	// read preset
	got, err := afero.ReadFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+newPreset))
	if err != nil {
		t.Error(err)
	}
	err = compareRead("create", expected, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateFromPreset(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	preset := "temp"
	newPreset := "new"
	expected := []byte("ROLE=swap")

	// populate fs
	afero.WriteFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+preset), expected, os.ModePerm)

	// create preset
	err := CreateFrom(cfg, afs, newPreset, preset)
	if err != nil {
		t.Fatal(err)
	}

	// read preset
	got, err := afero.ReadFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+newPreset))
	if err != nil {
		t.Error(err)
	}
	err = compareRead("create", expected, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateWithoutBase(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	// create preset without base preset
	err := CreateFrom(cfg, afs, "temp", "base")
	if err == nil {
		t.Fatal("expected error while creating preset without base preset")
	}
	t.Log(err)
}
