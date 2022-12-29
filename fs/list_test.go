package fs

import (
	"os"
	"sort"
	"testing"

	"github.com/spf13/afero"
)

func TestList(t *testing.T) {
	testfs := afero.NewMemMapFs()

	target := "dest"
	expected := []string{"first", "second", "thrid", "next/first"}

	// populate fs
	for _, file := range expected {
		err := afero.WriteFile(testfs, PathFromFormattedName(target+"/"+file), []byte("testdata"), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	// list all files
	got, err := List(testfs, target)
	if err != nil {
		t.Fatal(err)
	}

	// check if length matches
	if len(expected) != len(got) {
		t.Fatalf("unexpected number of files returned, expected %d got %d", len(expected), len(got))
	}

	// check if all items match in exact order
	sort.Strings(expected)
	sort.Strings(got)
	for i, item := range got {
		if item != expected[i] {
			t.Fatalf("files do not match expected result, expected %s got %s", item, expected[i])
		}
	}
}

func TestListIgnoreNonEnvFiles(t *testing.T) {
	testfs := afero.NewMemMapFs()

	target := "dest"
	files := []struct {
		path   string
		prefix bool
	}{
		{path: "first", prefix: true},
		{path: "second", prefix: false},
		{path: "third", prefix: false},
		{path: "fourth", prefix: true},
		{path: "next/first", prefix: true},
	}

	// populate fs
	for _, file := range files {
		var err error
		// add prefix conditionally
		if file.prefix {
			err = afero.WriteFile(testfs, PathFromFormattedName(target+"/"+file.path), []byte("testdata"), os.ModePerm)
		} else {
			err = afero.WriteFile(testfs, target+"/"+file.path, []byte("testdata"), os.ModePerm)
		}
		if err != nil {
			t.Error(err)
		}
	}

	var expected []string
	// expect only files with prefix
	for _, file := range files {
		if file.prefix {
			expected = append(expected, file.path)
		}
	}

	// list all files
	got, err := List(testfs, target)
	if err != nil {
		t.Fatal(err)
	}

	// check if length matches
	if len(expected) != len(got) {
		t.Fatalf("unexpected number of files returned, expected %d got %d", len(expected), len(got))
	}

	// check if all items match in exact order
	sort.Strings(expected)
	sort.Strings(got)
	for i, item := range got {
		if item != expected[i] {
			t.Fatalf("files do not match expected result, expected %s got %s", expected[i], item)
		}
	}
}
