package internal

import "io/fs"

type DirReader interface {
	ReadDir(dirName string) ([]fs.DirEntry, error)
	ReadLink(name string) (string, error)
	ReadFile(filePath string) (string, error)
}
