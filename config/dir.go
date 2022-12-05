package config

import "os"

// Swapenv working dir.
func Getwd() string {
	wd, _ := os.Getwd()

	return wd
}
