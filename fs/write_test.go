package fs

import (
	"testing"
)

// Write writes data to a file named by filename
func TestWrite(t *testing.T) {
	target := "./example/.env.write"
	data := []byte("hello=world")
	n, err := Write(target, data)
	

	if err != nil {
		t.Error(err)
	}

	

	if n != len(data) {
		t.Errorf("Write %d bytes, expected %d bytes", n, len(data))
	}


}