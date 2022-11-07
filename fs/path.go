package fs

import (
	"path/filepath"
	"strings"
)

// Convert path to formatted name string.
func PathToFormattedName(path string) (filename string) {
	prefix, name := filepath.Split(path)

	// remove .env. from file name
	name = strings.TrimPrefix(name, ".env.")

	// join prefix and file name
	filename = filepath.Join(prefix, name)
	return
}

// Get path from formatted name string.
func PathFromFormattedName(filename string) (path string) {
	prefix, name := filepath.Split(filename)

	// add .env. to file name
	name = ".env." + name

	// join prefix and file name
	path = filepath.Join(prefix, name)
	return
}
