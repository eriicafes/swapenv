package fs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// List all env files in directory and sub directories
func List(afs afero.Fs, dir string) ([]string, error) {
	files := []string{}

	err := afero.Walk(afs, dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// operate only on files with .env prefix
		if !info.IsDir() && strings.HasPrefix(info.Name(), ".env.") {
			// get relative path to dir
			path, err = filepath.Rel(dir, path)

			// ignore error in getting relative path
			if err != nil {
				return nil
			}

			files = append(files, PathToFormattedName(path))
		}

		return nil
	})

	return files, err
}
