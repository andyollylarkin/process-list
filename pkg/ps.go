package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

func ListAllProcess(lister ProcessLister) ([]Process, error) {
	return lister.ListProcess(nil)
}

func FindProcessByNameContains(lister ProcessLister, namePath string) ([]Process, error) {
	res := make([]Process, 0)
	procs, err := lister.ListProcess(func(i int, currentPsName string) bool {
		return strings.Contains(currentPsName, namePath)
	})

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

	procs, err := lister.ListProcess(func(pid int, currentPsName string) bool {
		return strings.EqualFold(name, currentPsName)
	})
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
	procs, err := lister.ListProcess(func(i int, currentPsName string) bool {
		return regex.MatchString(currentPsName)
	})
	if err != nil {
		return nil, err
	}

	return procs, nil
}

func FindProcessByPid(lister ProcessLister, pid int) (Process, error) {
	procs, err := lister.ListProcess(func(i int, currentPsName string) bool {
		return i == pid
	})
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
