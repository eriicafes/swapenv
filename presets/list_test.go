package presets

import (
	"os"
	"sort"
	"testing"

	"github.com/eriicafes/swapenv/config"
	"github.com/eriicafes/swapenv/fs"
	"github.com/spf13/afero"
)

func TestList(t *testing.T) {
	cfg := config.NewMemConfig()
	afs := afero.NewMemMapFs()

	files := []string{"temp", "next", "last"}

	// populate fs
	for _, file := range files {
		afero.WriteFile(afs, fs.PathFromFormattedName(cfg.Dir()+"/"+file), []byte("testdata"), os.ModePerm)
	}

	// list presets
	got := List(cfg, afs)

	// check if length matches
	if len(files) != len(got) {
		t.Fatalf("list error, expected %d files got %d", len(files), len(got))
	}
	// check if contents matches in exact order
	sort.Strings(files)
	sort.Strings(got)
	for i, item := range got {
		if item != files[i] {
			t.Fatalf("files do not match expected result, expected %s got %s", files[i], item)
		}
	}
}
