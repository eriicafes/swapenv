package fs

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T){
    target := "../example/.env.exists"
    got := Exists(target)
    want := true

    if got != want {
        t.Errorf("got %t, wanted %t", got, want)
    }
}

func TestExistsFalse(t *testing.T){
    target := "../example/.env.doesnotexist"
    got := Exists(target)
    want := false

    if got != want {
        t.Errorf("got %t, wanted %t", got, want)
    }
}

func TestReadFileContents(t *testing.T) {
    target := "../example/.env.hello"
    got := string(ReadFileContents(target))
    want := "hello=world"

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}

func TestGetDirectoryContents(t *testing.T) {
    target := "../example"
    got := GetDirectoryContents(target)
    // want := []string{".env.exists", ".env.hello"}
    // print got
    fmt.Println(got)
    // Check the length of the got slice
    if len(got) != 3 {
        t.Errorf("got %v, wanted %v", got, 3)
    }
}