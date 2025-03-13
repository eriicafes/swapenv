package env

import (
	"io/fs"
	"net/url"
	"strings"
)

// List returns all envs.
func List(fsys fs.FS, envsdir string) (envs []string) {
	entries, err := fs.ReadDir(fsys, envsdir)
	if err != nil {
		return
	}
	for _, d := range entries {
		name, hasPrefix := strings.CutPrefix(d.Name(), ".env.")
		if !hasPrefix {
			continue
		}
		if s, err := url.QueryUnescape(name); err == nil {
			envs = append(envs, s)
		}
	}
	return
}
