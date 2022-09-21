package fs

import (
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func Read(path string) ([]byte, error) {
	
	file, err := openFileRead(path);
	if err != nil {
		return []byte{}, nil
	}
	defer file.Close()
	
	buffer := make([]byte, 4 * 1024)
	
}


