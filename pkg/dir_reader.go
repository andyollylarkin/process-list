package pkg

import "io/fs"

//go:generate mockgen -destination=mock_dir_reader.go -package=pkg . DirReader
type DirReader interface {
	Open(name string) (fs.File, error)
	ReadDir(dirName string) ([]fs.DirEntry, error)
	ReadLink(name string) (string, error)
	ReadFile(filePath string) (string, error)
}
