package config

import (
	"errors"

	"github.com/spf13/viper"
)

func SetEnvPreset(preset string) error {
	viper.Set("env.preset", preset)

	return writeConfig()
}

func writeConfig() error {
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
