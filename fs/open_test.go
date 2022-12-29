package fs

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenRead(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	content := []byte("testdata")

	// populate fs
	afero.WriteFile(testfs, filename, content, os.ModePerm)

	// read file
	file, err := Open(testfs, filename, FlagRead)
	if err != nil {
		t.Error(err)
	}
	got, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare length and contents
	err = compareReadLengthAndContents(content, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenReadCreate(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	content := []byte("")

	// read non-existent file
	file, err := Open(testfs, filename, FlagReadCreate)
	if err != nil {
		t.Error(err)
	}
	got, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare length and contents
	err = compareReadLengthAndContents(content, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenWrite(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	content := []byte("testdata")

	// write to file
	file, err := Open(testfs, filename, FlagWrite)
	if err != nil {
		t.Error(err)
	}
	n, err := file.Write(content)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare written length
	err = compareWrittenLength(len(content), n)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenWriteOverride(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	initialContent := []byte("dummydata")
	content := []byte("testdata")

	// populate fs
	afero.WriteFile(testfs, filename, initialContent, os.ModePerm)

	// write to file
	file, err := Open(testfs, filename, FlagWrite)
	if err != nil {
		t.Error(err)
	}
	n, err := file.Write(content)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare written length
	err = compareWrittenLength(len(content), n)
	if err != nil {
		t.Fatal(err)
	}
	// compare written contents
	got, err := afero.ReadFile(testfs, filename)
	if err != nil {
		t.Fatal(err)
	}
	err = compareWrittenContents(content, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenWriteAppend(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	initialContent := []byte("test")
	content := []byte("data")
	finalContent := append(initialContent, content...)

	// populate fs
	afero.WriteFile(testfs, filename, initialContent, os.ModePerm)

	// write to file
	file, err := Open(testfs, filename, FlagWriteAppend)
	if err != nil {
		t.Error(err)
	}
	n, err := file.Write(content)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare written length
	err = compareWrittenLength(len(content), n)
	if err != nil {
		t.Fatal(err)
	}
	// compare written contents
	got, err := afero.ReadFile(testfs, filename)
	if err != nil {
		t.Fatal(err)
	}
	err = compareWrittenContents(finalContent, got)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenReadWriteAppend(t *testing.T) {
	testfs := afero.NewMemMapFs()

	filename := "example/.env.demo"
	initialContent := []byte("test")
	content := []byte("data")
	finalContent := append(initialContent, content...)

	// populate fs
	afero.WriteFile(testfs, filename, initialContent, os.ModePerm)

	// open file
	file, err := Open(testfs, filename, FlagReadWriteAppend)
	if err != nil {
		t.Error(err)
	}
	// manually seek to start due to a possible bug in afero.NewMemMapFs
	file.Seek(0, io.SeekStart)

	// read file
	got, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	// compare read length and contents
	err = compareReadLengthAndContents(initialContent, got)
	if err != nil {
		t.Fatal(err)
	}

	// write to file
	n, err := file.Write(content)
	if err != nil {
		t.Error(err)
	}
	// close file
	file.Close()

	// compare written length
	err = compareWrittenLength(len(content), n)
	if err != nil {
		t.Fatal(err)
	}
	// compare written contents
	got, err = afero.ReadFile(testfs, filename)
	if err != nil {
		t.Fatal(err)
	}
	err = compareWrittenContents(finalContent, got)
	if err != nil {
		t.Fatal(err)
	}
}

func compareReadLengthAndContents(expected []byte, got []byte) error {
	if len(expected) != len(got) {
		return fmt.Errorf("read failed, expected to read %d bytes read %d bytes", len(expected), len(got))
	}
	expectedString := string(expected)
	gotString := string(got)
	if expectedString != gotString {
		return fmt.Errorf("read failed, expected %s got %s", expectedString, gotString)
	}
	return nil
}

func compareWrittenLength(expected int, got int) error {
	if expected != got {
		return fmt.Errorf("write failed, expected to write %d bytes wrote %d", expected, got)
	}
	return nil
}

func compareWrittenContents(expected []byte, got []byte) error {
	expectedString := string(expected)
	gotString := string(got)
	if expectedString != gotString {
		return fmt.Errorf("write failed, expected %s got %s", expectedString, gotString)
	}
	return nil
}
