package config

import (
	"os"
	"os/exec"
	"strings"
)

// DefaultDir is swapenv default config dir.
const DefaultDir = "env"

// ConfigName is swapenv config file name.
const ConfigName = ".swapenvcache.yaml"

// ConfigRootName is swapenv root config file name.
const ConfigRootName = ".swapenv.yaml"

// Getwd returns swapenv working dir.
func Getwd() string {
	if gitwd, ok := GitRoot(); ok {
		return gitwd
	}

	wd, _ := os.Getwd()

	return wd
}

// GitRoot returns the git working dir and a boolean to show if command was successful.
func GitRoot() (string, bool) {
	root, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		return "", false
	}

	return strings.TrimSpace(string(root)), true
}
