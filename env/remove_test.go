package env

import (
	"io/fs"
	"net/url"
	"slices"
	"testing"
	"testing/fstest"

	"github.com/eriicafes/wfs"
)

func TestRemove(t *testing.T) {
	fsys := wfs.Map(fstest.MapFS{
		".git/envs/.env.staging": &fstest.MapFile{
			Data: []byte("ENV=staging"),
		},
		".git/envs/.env." + url.QueryEscape("staging/sub"): &fstest.MapFile{
			Data: []byte("ENV=staging/sub"),
		},
		".git/envs/.env.dev": &fstest.MapFile{
			Data: []byte("ENV=dev"),
		},
		".git/envs/.env.prod": &fstest.MapFile{
			Data: []byte("ENV=prod"),
		},
	})

	// remove single
	got := Remove(fsys, Env{Dir: ".git/envs", Name: "staging"})
	expected := []string{"staging"}
	if !slices.Equal(got, expected) {
		t.Errorf("Remove failed: expected %v got %v", expected, got)
	}
	_, err := fs.Stat(fsys, ".git/envs/.env.staging")
	if err == nil {
		t.Errorf("Remove failed: expected to remove single file")
	}

	// remove escaped
	got = Remove(fsys, Env{Dir: ".git/envs", Name: "staging/sub"})
	expected = []string{"staging/sub"}
	if !slices.Equal(got, expected) {
		t.Errorf("Remove failed: expected %v got %v", expected, got)
	}
	_, err = fs.Stat(fsys, ".git/envs/.env."+url.QueryEscape("staging/sub"))
	if err == nil {
		t.Errorf("Remove failed: expected to remove escaped file")
	}

	// remove multiple
	got = Remove(fsys, Env{Dir: ".git/envs", Name: "dev"}, Env{Dir: ".git/envs", Name: "prod"})
	expected = []string{"dev", "prod"}
	if !slices.Equal(got, expected) {
		t.Errorf("Remove failed: expected %v got %v", expected, got)
	}
	_, err = fs.Stat(fsys, ".git/envs/.env.dev")
	if err != nil {
		_, err = fs.Stat(fsys, ".git/envs/.env.prod")
	}
	if err == nil {
		t.Errorf("Remove failed: expected to remove multiple files: %v", err)
	}
}
