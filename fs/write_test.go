package fs

import "testing"


func TestCreateFile(t *testing.T) {
	target := "../example/.env.create"
	CreateFile(target)
	got := Exists(target)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
	RemoveFile(target)
}

func TestCopyFile(t *testing.T){
	target := "../example/.env.copy"
	got := CopyFile("../example/.env.exists", target)
	var want error = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	RemoveFile(target)
}

func TestWriteFileContents(t *testing.T) {
	target := "../example/.env.write"
	WriteFileContents(target, "hello=world")
	got := string(ReadFileContents(target))
	want := "hello=world"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	RemoveFile(target)
}

func TestRemoveFile(t *testing.T) {
	target := "../example/.env.remove"
	CreateFile(target)
	got := RemoveFile(target)
	var want error = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}


func TestRemoveDirectory(t *testing.T) {
	target := "../exist"
	MakeDir(target)
	got := RemoveDir(target)
	var want error = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMakeDir(t *testing.T) {
	target := "../example/make"
	got := MakeDir(target)
	var want error = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	RemoveDir(target)
}