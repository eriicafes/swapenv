package config

import (
	"errors"
)

var cfg = NewViperConfig(Getwd())

// Get returns a pointer to the config singleton.
func Get() Config {
	return cfg
}

// EnsureHasInit checks if config has been initialized.
func EnsureHasInit(c Config) error {
	if c.HasInit() {
		return nil
	}
	return errors.New("swapenv has not been initialized on this project, run `swapenv init`")
}

// EnsureHasNotInit checks if config has not been initialized.
func EnsureHasNotInit(c Config) error {
	if c.HasInit() {
		return errors.New("swapenv has already been initialized on this project")
	}
	return nil
}
