package readers

import (
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
