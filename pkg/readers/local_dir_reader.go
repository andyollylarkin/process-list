package readers

import (
	"io"
	"io/fs"
	"os"
)

type LocalDirReader struct{}

func NewLocalDirReader() *LocalDirReader {
	return &LocalDirReader{}
}

func (r *LocalDirReader) ReadDir(dirName string) ([]fs.DirEntry, error) {
	return os.ReadDir(dirName)
}

func (r *LocalDirReader) ReadLink(name string) (string, error) {
	return os.Readlink(name)
}

func (r *LocalDirReader) ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
