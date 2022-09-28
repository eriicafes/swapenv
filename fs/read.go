package fs

import (
	"errors"
	"io"
)
func Read(path string) ([]byte, error) {
	
	file, err := openFileRead(path);
	if err != nil {
		return []byte{}, nil
	}
	defer file.Close()
	

	// loop through the file and read it into a byte array

	var bytes []byte
	buffer := make([]byte, 4 * 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return bytes, err
		}
		bytes = append(bytes, buffer[:n]...)
	}

	return bytes, err
}


