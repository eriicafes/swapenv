package fs

func Write(path string, data []byte) error {
	file, err := openFileWrite(path)

	if err != nil {
		return err
	}
	
	defer file.Close()
	_, err = file.Write(data)
	return err
}
