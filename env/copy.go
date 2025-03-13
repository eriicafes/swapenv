package env

import (
	"io"
	"io/fs"
	"os"

	"github.com/eriicafes/wfs"
)

// Copy writes contents of src to dest.
//
// If dest does not exist it will be created.
func Copy(fsys wfs.FS, src, dest Env) error {
	srcFile, err := src.Open(fsys)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := dest.OpenFile(fsys, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
