package env

import (
	"io/fs"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/eriicafes/wfs"
)

func TestCopy(t *testing.T) {
	fsys := wfs.Map(fstest.MapFS{
		".env": &fstest.MapFile{
			Data: []byte("ENV=local"),
		},
		".git/envs/.env.staging": &fstest.MapFile{
			Data: []byte("ENV=staging"),
		},
		".git/envs/.env.dev": &fstest.MapFile{
			Data: []byte("ENV=dev"),
		},
	})

	tests := []struct {
		name         string
		src, dest    Env
		expectedDest string
	}{
		{
			name:         "MissingSrc",
			src:          Env{Name: ".env.local", raw: true},
			dest:         Env{Dir: ".git/envs", Name: "local"},
			expectedDest: "",
		},
		{
			name:         "ExistingDest",
			src:          Env{Name: ".env", raw: true},
			dest:         Env{Dir: ".git/envs", Name: "staging"},
			expectedDest: "ENV=local",
		},
		{
			name:         "EscapedDest",
			src:          Env{Name: ".env", raw: true},
			dest:         Env{Dir: ".git/envs", Name: "staging/sub"},
			expectedDest: "ENV=local",
		},
		{
			name:         "To .env",
			src:          Env{Dir: ".git/envs", Name: "dev"},
			dest:         Env{Name: ".env", raw: true},
			expectedDest: "ENV=dev",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(fsys, tc.src, tc.dest)
			if tc.expectedDest == "" {
				if err == nil {
					t.Fatalf("Commit failed, expected error")
				} else {
					return
				}
			}

			if err != nil {
				t.Fatalf("Commit failed: %v", err)
			}
			b, err := fs.ReadFile(fsys, filepath.Join(tc.dest.Path()))
			if err != nil || string(b) != tc.expectedDest {
				t.Errorf("Commit failed, expected %q got %q, err: %v", tc.expectedDest, b, err)
			}
		})
	}
}
