package config

import (
	"os"
	"os/exec"
	"strings"
)

var DefaultWd = "env"

// Getwd returns swapenv config dir.
func Getwd() string {
	if gitwd, ok := gitRoot(); ok {
		return gitwd
	}

	wd, _ := os.Getwd()

	return wd
}

// gitRoot returns the git working dir and a boolean to show if command was successful.
func gitRoot() (string, bool) {
	root, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		return "", false
	}

	return strings.TrimSpace(string(root)), true
}
