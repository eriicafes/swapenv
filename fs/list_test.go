package fs

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	target := "../example"

	// List all files
	files, err := List(target)

	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Files: %v", files)
}