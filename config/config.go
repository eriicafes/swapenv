package config

import (
	"errors"
)

var cfg = NewViperConfig(Getwd())

func Get() Config {
	return cfg
}

// Ensure config has been initialized.
func EnsureHasInit(c Config) error {
	if c.HasInit() {
		return nil
	}
	return errors.New("swapenv has not been initialized on this project, run `swapenv init`")
}

// Ensure config has not been initialized.
func EnsureHasNotInit(c Config) error {
	if c.HasInit() {
		return errors.New("swapenv has already been initialized on this project")
	}
	return nil
}
