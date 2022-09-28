package fs

import (
	"testing"
)

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