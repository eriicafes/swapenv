package fs

import (
	"bytes"
	"testing"
)

func TestRead(t *testing.T) {
	target := "../example/.env.hello"
	check := []byte("hello=world")

	content, err := Read(target)

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(content, check) {
		t.Errorf("Read %d bytes, expected %d bytes", len(content), len(check))
	}

}
