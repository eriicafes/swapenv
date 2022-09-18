package fs

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(path string){
	filePtr, err := os.Create(path);
	if err != nil {
		log.Fatal(err);
	}
	defer filePtr.Close(); // close the file
	// We can read from and write to the file
}

func CopyFile(source, target string) error {
	return copyFileMode(source, target, 0644)
}

func copyFileMode(src, tgt string, perm os.FileMode) (err error) {
	if _, err := os.Stat(src); err != nil {
		return err
	}

	source, err := os.Open(filepath.Clean(src))
	if err != nil {
		return err
	}
	defer source.Close()

	target, err := os.OpenFile(tgt, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer func() {
		if r := target.Close(); r != nil && err == nil {
			err = r
		}
	}()

	_, err = io.Copy(target, source)
	return err
}

func WriteFileContents(path string, text string) {
	content := text;
  
  	/* os.WriteFile takes in file path, a []byte of the file content, 
     	and permission bits in case file doesn't exist */
  
	err := os.WriteFile(path, []byte(content), 0666);
	if err != nil {
		log.Fatal(err);
	}
}


func AppendToFile(path string, text string) {
	content := "\n" + text;
	file, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, 0644);
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();
	file.Write([]byte(content));
}

func RemoveFile(path string) error {
	err := os.Remove(path);
	if err != nil{
		log.Fatal(err);
	}
	return nil
}

func MakeDir(dir string) error{
	err := os.Mkdir(dir, 0755);
	if err != nil {
		log.Fatal(err);
	}
	return nil
  }

func RemoveDir(dir string) error {
	err := os.Remove(dir);
	if err != nil{
		log.Fatal(err);
	}
	return nil
}