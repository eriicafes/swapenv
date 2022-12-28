package fs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

// Open opens file in path with provided flags.
func Open(afs afero.Fs, path string, flag flag) (afero.File, error) {
	// attempt to open file
	file, err := afs.OpenFile(path, flag.flag, os.ModePerm)
	if err == nil {
		return file, nil
	}

	// retry if error is because parent directories does not exist
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	// attempt to create parent directories
	dirname := filepath.Dir(path)
	if merr := afs.MkdirAll(dirname, os.ModePerm); merr == nil {
		// retry opening file
		return Open(afs, path, flag)
	}

	return nil, err
}
