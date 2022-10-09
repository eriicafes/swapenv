package fs

import "testing"

func TestOpenRead(t *testing.T) {
	file, err := openFileRead("../example/.env.hello")

	if err != nil {
		t.Error("could not open file")
	}

	bufferSize := 1

	n, err := file.Read(make([]byte, bufferSize))

	if err != nil || n != bufferSize {
		t.Error("could not read file")
	}
}

func TestOpenWrite(t *testing.T) {
	file, err := openFileWrite("../example/.env.hello")

	if err != nil {
		t.Error("could not open file")
	}

	content := []byte("hello=world")

	_, err = file.Write(content)

	if err != nil {
		t.Error("could not write to file")
	}
}
