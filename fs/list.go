package fs

import (
	"os"
	"path/filepath"
	"strings"
)

// List all env files in directory and sub directories
func List(dir string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// operate only on files
		if !info.IsDir() {
			// get relative path to dir
			path, err = filepath.Rel(dir, path)

			// ignore error in getting relative path
			if err != nil {
				return nil
			}

			// check if file has .env. prefix
			if strings.HasPrefix(filepath.Base(path), ".env.") {
				files = append(files, PathToFormattedName(path))
			}
		}

		return nil
	})

	return files, err
}
