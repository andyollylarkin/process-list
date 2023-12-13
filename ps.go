package ps

import "gitlab.mindsw.io/migrate-core-libs/process-list/pkg"

func ListProcess() ([]*pkg.Process, error) {
	return pkg.ListProcess()
}

func FindProcessByNameContains(namePath string) ([]*pkg.Process, error) {
	return pkg.FindProcessByNameContains(namePath)
}

func FindProcessByNameEqual(name string) ([]*pkg.Process, error) {
	return pkg.FindProcessByNameEqual(name)
}

func FindProcessByPid(pid int) (*pkg.Process, error) {
	return pkg.FindProcessByPid(pid)
}
