package fs

import (
	"testing"
)

// Write writes data to a file named by filename
func TestWrite(t *testing.T) {
	target := "../example/.env.hello"
	data := []byte("hello=world")
	err := Write(target, data)

	if err != nil {
		t.Error(err)
	}
}
