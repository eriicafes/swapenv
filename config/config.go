package config

import (
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

var _wdir, _envsdir string
var _loaded bool

// Loaded returns true if config has been loaded.
func Loaded() bool { return _loaded }

// Dir returns the git project directory "." and envs directory ".git/envs".
func Dir() (wdir, envsdir string) { return _wdir, _envsdir }

// Get returns the current and previous env.
func GetEnv() (current, prev string) {
	return viper.GetString("branch.current"), viper.GetString("branch.prev")
}

// Set changes the current env.
func SetEnv(env string) error {
	current := viper.GetString("branch.current")
	if current == "" {
		current = env
	}
	viper.Set("branch.prev", current)
	viper.Set("branch.current", env)
	return viper.WriteConfig()
}

func init() {
	if gdir, err := GitDir(); err == nil {
		_wdir, _envsdir = filepath.Dir(gdir), path.Join(gdir, "envs")
		viper.SetConfigFile(path.Join(_envsdir, ".swapenv.yaml"))
		_loaded = viper.ReadInConfig() == nil
	}
}
