package fs

import (
	"path/filepath"
	"strings"
)

// Returns the formatted name string of env file.
func FormatPath(path string) (filename string) {
	prefix, name := filepath.Split(path)

	// remove .env. from file name
	name = strings.TrimPrefix(name, ".env.")

	// join prefix and file name
	filename = filepath.Join(prefix, name)
	return
}

// Returns the file path string of env file.
func NormalizePath(filename string) (path string) {
	prefix, name := filepath.Split(filename)

	// add .env. to file name
	name = ".env." + name

	// join prefix and file name
	path = filepath.Join(prefix, name)
	return
}
