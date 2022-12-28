package config

import (
	"path"

	"github.com/spf13/viper"
)

type viperConfig struct {
	wd     string
	loaded bool

	config     *viper.Viper
	rootConfig *viper.Viper
}

func NewViperConfig(wd string) *viperConfig {
	vc := &viperConfig{
		wd:         wd,
		config:     viper.New(),
		rootConfig: viper.New(),
	}

	vc.setupRoot()
	vc.setup()

	return vc
}

func (vc *viperConfig) setup() {
	// setup config
	vc.config.SetConfigFile(path.Join(vc.Dir(), ConfigName))

	// read from file
	if err := vc.config.ReadInConfig(); err != nil {
		// return early on read error
		return
	}

	// set loaded to true on successful read from file
	vc.loaded = true
}

func (vc *viperConfig) setupRoot() {
	// setup root config
	vc.rootConfig.SetConfigFile(path.Join(vc.wd, ConfigRootName))

	// set default dir
	vc.rootConfig.SetDefault("dir", DefaultDir)

	// read from file
	_ = vc.rootConfig.ReadInConfig()
}

func (vc *viperConfig) Dir() string {
	dir := vc.rootConfig.GetString("dir")
	return path.Join(vc.wd, dir)
}

func (vc *viperConfig) ChDir(dir string) error {
	vc.rootConfig.Set("dir", dir)

	// reset config
	vc.config = viper.New()
	vc.setup()

	return vc.rootConfig.WriteConfig()
}

func (vc *viperConfig) HasInit() bool {
	return vc.loaded
}

func (vc *viperConfig) Flush() error {
	return vc.config.WriteConfig()
}

func (vc *viperConfig) GetPreset() string {
	return vc.config.GetString("env.preset")
}

func (vc *viperConfig) SetPreset(preset string) {
	vc.config.Set("env.preset", preset)
}
