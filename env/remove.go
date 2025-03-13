package env

import (
	"github.com/eriicafes/wfs"
)

// Remove deletes and returns a slice of removed env names.
func Remove(fsys wfs.FS, envs ...Env) (removed []string) {
	for _, env := range envs {
		if err := fsys.Remove(env.Path()); err == nil {
			removed = append(removed, env.Name)
		}
	}
	return
}
