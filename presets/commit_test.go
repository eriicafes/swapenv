package presets

import (
	"os"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestCommitToPreset(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	preset := "temp"
	expected := []byte("ROLE=swap")

	// setup config
	cfg.SetPreset(preset)

	// populate fs
	afero.WriteFile(afs, ".env", expected, os.ModePerm)

	// commit preset
	err := Commit(cfg, afs)
	if err != nil {
		t.Fatal(err)
	}

	// read preset
	got, err := afero.ReadFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+preset))
	if err != nil {
		t.Error(err)
	}
	err = compareRead("commit", expected, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommitWithNoEnv(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	// setup config
	cfg.SetPreset("temp")

	// commit preset without .env
	err := Commit(cfg, afs)
	if err == nil {
		t.Fatal("expected error while commiting preset without .env file")
	}
	t.Log(err)
}
