package config

import (
	"os/exec"
	"strings"
)

func GitDir() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	return strings.TrimSpace(string(out)), err
}

func GitBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(out)), err
}
