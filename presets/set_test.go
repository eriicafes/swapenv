package presets

import (
	"testing"

	"github.com/eriicafes/swapenv/config"
)

func TestSet(t *testing.T) {
	cfg := config.NewMemConfig()

	preset := "temp"
	nextPreset := "next"

	// setup config
	cfg.SetPreset(preset)

	// set preset
	err := UncheckedSet(cfg, nextPreset)
	if err != nil {
		t.Fatal(err)
	}

	// check preset
	got := cfg.GetPreset()
	if nextPreset != got {
		t.Fatalf("set error, expected %s got %s", nextPreset, got)
	}
}
