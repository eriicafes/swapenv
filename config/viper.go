package config

import (
	"errors"
	"path"

	"github.com/spf13/viper"
)

type env struct {
	Preset string
}

type viperConfig struct {
	env    env
	base   string
	loaded bool
}

func NewViperConfig(wd string) *viperConfig {
	base := path.Join(wd, "env")

	config := &viperConfig{
		base: base,
	}

	config.setup()

	return config
}

func (vc *viperConfig) setup() {
	// configure viper
	viper.SetConfigName(".swapenvcache")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(vc.base)

	// unmarshal into struct before exiting function
	defer viper.UnmarshalKey("env", &vc.env)

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		// return early on read error
		return
	}

	// set loaded to true on successful read from file
	vc.loaded = true
}

func (vc *viperConfig) write() error {
	if err := viper.WriteConfig(); err != nil {
		var fnfe viper.ConfigFileNotFoundError

		// retry with safe write if file not found
		if errors.As(err, &fnfe) {
			return viper.SafeWriteConfig()
		}
		return err
	}
	return nil
}

func (vc *viperConfig) Base() string {
	return vc.base
}

func (vc *viperConfig) HasInit() bool {
	return vc.loaded
}

func (vc *viperConfig) GetPreset() string {
	return vc.env.Preset
}

func (vc *viperConfig) SetPreset(preset string) error {
	viper.Set("env.preset", preset)
	vc.env.Preset = preset

	return vc.write()
}
