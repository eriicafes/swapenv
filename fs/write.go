package fs

func Write(path string, data []byte) (int, error) {
	file, err := openFileWrite(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}
