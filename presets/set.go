package presets

import "github.com/eriicafes/swapenv/config"

// UncheckedSet sets preset as the current preset.
//
// NOTE: no checks are made before updating preset.
func UncheckedSet(cfg config.Config, preset string) error {
	cfg.SetPreset(preset)
	return cfg.Flush()
}
