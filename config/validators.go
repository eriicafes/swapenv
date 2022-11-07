package config

import "errors"

// Ensure config file is present.
func EnsureHasInitialized() error {
	if LoadedFromFile {
		return nil
	}

	return errors.New("swapenv has not been initialized on this project, run `swapenv init`")
}

// Ensure config file is not present.
func EnsureHasNotInitialized() error {
	if LoadedFromFile {
		return errors.New("swapenv has already been initialized on this project")
	}

	return nil
}
