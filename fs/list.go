package fs

import (
	"os"
	"path/filepath"
	"strings"
)



func List(dir string) ([]string, error) {
	// List the files in a directory using filepath.Walk
	// Change dir to absolute path
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	files := []string{}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() {
			// Remove the directory from the path
			base := strings.TrimPrefix(path, dir)
			filename := filepath.Base(path)
			prefix := filepath.Dir(base)
			if strings.HasPrefix(filename, ".env.") {
				files = append(files, filepath.Join(prefix[1:],strings.TrimPrefix(filename, ".env.")))
			}
		}
		
		return nil
	})
	return files, err
}