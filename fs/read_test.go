package fs

import (
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

func TestRead(t *testing.T) {
    target := "../example/.env.hello"
    check := []byte("hello=world")

    n, err := Read(target)

    if err != nil {
        t.Error(err)
    }

    if len(n) != len(check) {
        t.Errorf("Read %d bytes, expected %d bytes", len(n), len(check))
    }

}