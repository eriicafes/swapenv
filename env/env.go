package env

import (
	"errors"
	"fmt"
	"io/fs"
	"net/url"
	"path/filepath"

	"github.com/eriicafes/wfs"
)

type Env struct {
	Dir, Name string
	raw       bool
}

func DotEnv(dir string) Env {
	return Env{Dir: dir, Name: ".env", raw: true}
}

// Path returns the filepath where env is stored after adding ".env." prefix and escaping.
func (e Env) Path() string {
	if e.raw {
		return filepath.Join(e.Dir, e.Name)
	}
	return filepath.Join(e.Dir, ".env."+url.QueryEscape(e.Name))
}

// Open opens a file returning a custom error if open fails.
func (e Env) Open(fsys fs.FS) (fs.File, error) {
	file, err := fsys.Open(e.Path())
	var perr *fs.PathError
	if errors.As(err, &perr) {
		err = fmt.Errorf("error: failed to open env '%s': %w", e.Name, perr.Err)
	}
	return file, err
}

// OpenFile opens a file for read/write returning a custom error if open fails.
func (e Env) OpenFile(fsys wfs.FS, flag int, perm fs.FileMode) (wfs.File, error) {
	file, err := fsys.OpenFile(e.Path(), flag, perm)
	var perr *fs.PathError
	if errors.As(err, &perr) {
		err = fmt.Errorf("error: failed to open env '%s': %w", e.Name, perr.Err)
	}
	return file, err
}
