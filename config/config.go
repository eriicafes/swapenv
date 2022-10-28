package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

// Config directory
const Base = "env"

// Config filename
const Filename = ".swapenvcache.yaml"

var Env env
var LoadedFromFile bool

func getConfigPath() string {
	wd, _ := os.Getwd()

	return path.Join(wd, Base)
}

func init() {
	// configure viper
	viper.SetConfigName(".swapenvcache")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getConfigPath())

	// set default values
	viper.SetDefault("env.preset", "local")

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		// reset config on error
		viper.WriteConfig()
	} else {
		// set LoadedFromFile to true on successful read from file
		LoadedFromFile = true
	}

	// unmarshal into struct
	viper.UnmarshalKey("env", &Env)
}
