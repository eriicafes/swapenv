package fs

import (
	"errors"
	"os"
	"path/filepath"
)

// Open file in file path with read only access.
func OpenFileRead(path string) (*os.File, error) {
	return openFile(path, os.O_RDONLY, false)
}

// Open file in file path with read only access. File will be created if it does not exist.
func OpenFileReadCreate(path string) (*os.File, error) {
	return openFile(path, os.O_RDONLY|os.O_CREATE, true)
}

// Open file in file path with write only access. File will be created if it does not exist.
func OpenFileWrite(path string) (*os.File, error) {
	return openFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, true)
}

// Open file in file path with write only access. File will be appended to and created if it does not exist.
func OpenFileAppend(path string) (*os.File, error) {
	return openFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, true)
}

// Open file in file path with provided flags. File will be created if it does not exist and createOnNotExist is true.
//
// os.O_RDONLY flag opens file with read only access
// os.O_WRONLY flag opens file with write only access
// os.O_APPEND flag appends to file while writing
// os.O_CREATE flag creates file if it does not exist
func openFile(path string, flag int, createOnNotExist bool) (*os.File, error) {
	// attempt to open file
	file, err := os.OpenFile(path, flag, os.ModePerm)

	// return file handle if no errors
	if err == nil {
		return file, nil
	}

	// error while opening file

	// check if should create file on not exist
	if !createOnNotExist {
		return nil, err
	}

	// check if error is not because parent directories does not exist
	// if true this is an unexpected error
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	// attempt to create parent directories
	dirname := filepath.Dir(path)
	if merr := os.MkdirAll(dirname, os.ModePerm); merr == nil {
		// retry opening file
		return openFile(path, flag, createOnNotExist)
	}

	return nil, err
}
