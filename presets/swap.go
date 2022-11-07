package presets

// Commit .env file to current preset and load preset to .env file.
func Swap(preset string) error {
	// commit current preset
	if err := Commit(); err != nil {
		return nil
	}

	// load preset
	return LoadUnchecked(preset)
}
