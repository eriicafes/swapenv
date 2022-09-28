package fs

import "io"

func Copy(src, dst string) (int64, error) {
	srcFile, err := openFileRead(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := openFileWrite(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}