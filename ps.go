package ps

import "gitlab.mindsw.io/migrate-core-libs/process-list/pkg"

func ListProcess(reader pkg.DirReader) ([]*pkg.Process, error) {
	return pkg.ListProcess(reader)
}

func FindProcessByNameContains(reader pkg.DirReader, namePath string) ([]*pkg.Process, error) {
	return pkg.FindProcessByNameContains(reader, namePath)
}

func FindProcessByNameEqual(reader pkg.DirReader, name string) ([]*pkg.Process, error) {
	return pkg.FindProcessByNameEqual(reader, name)
}

func FindProcessByPid(reader pkg.DirReader, pid int) (*pkg.Process, error) {
	return pkg.FindProcessByPid(reader, pid)
}
