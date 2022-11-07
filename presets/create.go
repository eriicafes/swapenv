package presets

import "fmt"

// Create preset using contents of .env file.
//
// The .env file is created if it does not exist.
func Create(preset string) error {
	// return error if preset already exists
	if Exists(preset) {
		return fmt.Errorf("env preset '%v' already exists", preset)
	}

	// TODO: create preset and copy contents of .env file to preset, create the .env file if it does not exist

	return nil
}

// Create preset using contents base preset.
func CreateFrom(preset string, base string) error {
	// return error if preset already exists
	if Exists(preset) {
		return fmt.Errorf("env preset '%v' already exists", preset)
	}

	// return error if base preset does not exist
	if !Exists(base) {
		return fmt.Errorf("base env preset '%v' does not exist", preset)
	}

	// TODO: create preset and copy contents of base preset to preset, return error if the base preset does not exist
	return nil
}
