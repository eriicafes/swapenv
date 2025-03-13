package env

import (
	"net/url"
	"slices"
	"testing"
	"testing/fstest"

	"github.com/eriicafes/wfs"
)

func TestList(t *testing.T) {
	fsys := wfs.Map(fstest.MapFS{
		".git/envs/.env.staging": &fstest.MapFile{
			Data: []byte("ENV=staging"),
		},
		".git/envs/.env." + url.QueryEscape("staging/sub"): &fstest.MapFile{
			Data: []byte("ENV=staging/sub"),
		},
	})

	got := List(fsys, ".git/envs")
	expected := []string{"staging", "staging/sub"}
	// check regular name
	if !slices.Equal(got, expected) {
		t.Fatalf("List failed: expected %v got %v", expected, got)
	}
}
