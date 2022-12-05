package config

type Config interface {
	// Swpaenv config dir.
	Base() string
	// Indicate if config has been initialized.
	HasInit() bool

	// Get current preset.
	GetPreset() string
	// Set current preset.
	SetPreset(preset string) error
}
