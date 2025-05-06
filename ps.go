package ps

import (
	"regexp"

	"github.com/andyollylarkin/process-list/pkg"
)

func ListProcess(lister pkg.ProcessLister) ([]pkg.Process, error) {
	return pkg.ListAllProcess(lister)
}

func FindProcessByNameContains(lister pkg.ProcessLister, namePath string) ([]pkg.Process, error) {
	return pkg.FindProcessByNameContains(lister, namePath)
}

func FindProcessByRegex(lister pkg.ProcessLister, regex string) ([]pkg.Process, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	return pkg.FindProcessByRegex(lister, *r)
}

func FindProcessByNameEqual(lister pkg.ProcessLister, name string) ([]pkg.Process, error) {
	return pkg.FindProcessByNameEqual(lister, name)
}

func FindProcessByPid(lister pkg.ProcessLister, pid int) (pkg.Process, error) {
	return pkg.FindProcessByPid(lister, pid)
}
