package fs

import (
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func ReadFileContents(path string) []byte {
	bytes, err := os.ReadFile(path);
	if err != nil {
    		log.Fatal(err);
	}
	return bytes
}


func GetDirectoryContents(dir string) []string {
	entries, err := os.ReadDir(dir);
	if err != nil {
		// log.Fatal(err);
		// Return empty slice
		return make([]string, 0)
	}

	// Create a slice to hold the names of the files
	names := make([]string, 0, len(entries))

    // Iterate through the directory entities and print out their name.
	for _, entry := range(entries) {
		names = append(names, entry.Name())
	}
	return names
}