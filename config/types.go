package config

type Config interface {
	// Dir returns swapenv config dir.
	Dir() string
	// ChDir updates swapenv config dir.
	ChDir(dir string) error

	// HasInit indicates if config has been initialized.
	HasInit() bool

	// Flush commits changes made to config.
	//
	// This makes multiple calls to update config more efficient by avoiding
	// the performance cost of writing to file after each Set call
	// by batching Set calls together and calling Flush when ready.
	Flush() error

	// Get current preset.
	GetPreset() string
	// Set current preset.
	SetPreset(preset string)
}
