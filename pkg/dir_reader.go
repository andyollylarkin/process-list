package pkg

import "io/fs"

type DirReader interface {
	ReadDir(dirName string) ([]fs.DirEntry, error)
	ReadLink(name string) (string, error)
}
