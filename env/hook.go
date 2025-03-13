package env

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/eriicafes/wfs"
)

const hookContent = "\n[[ $3 = 1 ]] && swapenv use $2 -b"

func InstallHook(fsys wfs.FS, gitdir string) error {
	hookPath := filepath.Join(gitdir, "hooks/post-checkout")

	file, err := fsys.OpenFile(hookPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// append shebang if file is empty
	b, _ := io.ReadAll(file)
	if len(b) == 0 {
		_, err = file.Write([]byte("#!/bin/sh"))
		if err != nil {
			return err
		}
	}

	// already preinitialized
	if bytes.Contains(b, []byte(hookContent)) {
		return nil
	}

	// append hook contents
	_, err = file.Write([]byte(hookContent))
	return err
}
