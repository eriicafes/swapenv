package presets

import (
	"fmt"

	"github.com/eriicafes/swapenv/config"
)

// Load preset to .env file.
//
// NOTE: contents of .env file will not be committed by this function. Explicitly call Commit before using this function to prevent data loss.
func LoadUnchecked(preset string) error {
	if !Exists(preset) {
		return fmt.Errorf("env preset '%v' does not exist", preset)
	}

	// TODO: write contents of preset to .env file

	// update config
	if err := config.SetEnvPreset(preset); err != nil {
		return err
	}

	return nil
}
