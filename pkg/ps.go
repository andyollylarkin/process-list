package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

func ListAllProcess(lister ProcessLister) ([]Process, error) {
	return lister.ListProcess()
}

func FindProcessByNameContains(lister ProcessLister, namePath string) ([]Process, error) {
	res := make([]Process, 0)
	procs, err := ListAllProcess(lister)
	if err != nil {
		return nil, err
	}

	for _, p := range procs {
		if strings.Contains(p.Name, namePath) {
			res = append(res, p)
		}
	}

	return res, nil
}

func FindProcessByNameEqual(lister ProcessLister, name string) ([]Process, error) {
	res := make([]Process, 0)
	procs, err := ListAllProcess(lister)
	if err != nil {
		return nil, err
	}

	for _, p := range procs {
		if p.Name == name {
			res = append(res, p)
		}
	}

	return res, nil
}

func FindProcessByRegex(lister ProcessLister, regex regexp.Regexp) ([]Process, error) {
	res := make([]Process, 0)
	procs, err := ListAllProcess(lister)
	if err != nil {
		return nil, err
	}

	for _, p := range procs {
		if regex.MatchString(p.Name) {
			res = append(res, p)
		}
	}

	return res, nil
}

func FindProcessByPid(lister ProcessLister, pid int) (Process, error) {
	procs, err := ListAllProcess(lister)
	if err != nil {
		return Process{}, err
	}

	for _, p := range procs {
		if p.Pid == pid {
			return p, nil
		}
	}

	return Process{}, fmt.Errorf("process with pid %d not found", pid)
}
