package env

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/eriicafes/wfs"
)

func TestInstallHook(t *testing.T) {
	tests := []struct {
		name     string
		fsys     wfs.FS
		expected string
	}{
		{
			name:     "None",
			fsys:     wfs.Map(fstest.MapFS{}),
			expected: fmt.Sprintf("#!/bin/sh%s", hookContent),
		},
		{
			name: "Empty",
			fsys: wfs.Map(fstest.MapFS{
				".git/hooks/post-checkout": &fstest.MapFile{},
			}),
			expected: "#!/bin/sh" + hookContent,
		},
		{
			name: "Existing",
			fsys: wfs.Map(fstest.MapFS{
				".git/hooks/post-checkout": &fstest.MapFile{
					Data: []byte("#!/bin/sh\necho hi"),
				},
			}),
			expected: "#!/bin/sh\necho hi" + hookContent,
		},
		{
			name: "PreInitialized",
			fsys: wfs.Map(fstest.MapFS{
				".git/hooks/post-checkout": &fstest.MapFile{
					Data: []byte("#!/bin/sh\necho hi" + hookContent),
				},
			}),
			expected: "#!/bin/sh\necho hi" + hookContent,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := InstallHook(tc.fsys, ".git")
			if err != nil {
				t.Error(err)
			}
			b, _ := fs.ReadFile(tc.fsys, ".git/hooks/post-checkout")
			if tc.expected != string(b) {
				t.Errorf("InstallHook(%s) failed: expected %q got %q", tc.name, tc.expected, string(b))
			}
		})
	}
}
