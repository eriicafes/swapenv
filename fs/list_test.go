package fs

import (
	"testing"
)

func TestList(t *testing.T) {
	target := "../example"

	expected := []string{"contents", "exists", "hello", "work/work"}

	// list all files
	files, err := List(target)

	if err != nil {
		t.Error(err)
	}

	// check if length matches
	if len(files) != len(expected) {
		t.Error("unexpected number of files returned")
	}

	// check if each item matches
	for i, file := range files {
		if file != expected[i] {
			t.Error("files do not match expected result")
		}
	}
}
