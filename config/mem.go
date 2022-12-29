package config

type env struct {
	preset string
}

type memConfig struct {
	env    env
	dir    string
	loaded bool
}

// NewMemConfig returns an in memmory config ready for use.
//
// Dir is set to "env", change dir by calling config.ChDir.
//
// Config is uninitialized, initialize by calling config.Init.
//
// Flush is noop.
//
// Preset is not set, set by calling config.SetPreset.
func NewMemConfig() *memConfig {
	return &memConfig{
		dir:    "env",
		loaded: false,
	}
}

func (mc *memConfig) Dir() string {
	return mc.dir
}

func (mc *memConfig) ChDir(dir string) error {
	mc.dir = dir
	return nil
}

func (mc *memConfig) Init() {
	mc.loaded = true
}

func (mc *memConfig) HasInit() bool {
	return mc.loaded
}

func (mc *memConfig) Flush() error {
	return nil
}

func (mc *memConfig) GetPreset() string {
	return mc.env.preset
}

func (mc *memConfig) SetPreset(preset string) {
	mc.env.preset = preset
}
