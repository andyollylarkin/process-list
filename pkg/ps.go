package pkg

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func ListProcess(reader DirReader) ([]*Process, error) {
	res := make([]*Process, 0)
	content, err := reader.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, d := range content {
		pid, err := strconv.ParseInt(d.Name(), 10, 0)
		if err != nil {
			continue
		}
		path := filepath.Join("/proc", d.Name(), "exe")
		exeName, err := reader.ReadLink(path)
		if err != nil {
			continue
		}

		res = append(res, &Process{
			Pid:  int(pid),
			Name: filepath.Base(exeName),
		})
	}

	return res, nil
}

func FindProcessByNameContains(reader DirReader, namePath string) ([]*Process, error) {
	res := make([]*Process, 0)
	procs, err := ListProcess(reader)
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

func FindProcessByNameEqual(reader DirReader, name string) ([]*Process, error) {
	res := make([]*Process, 0)
	procs, err := ListProcess(reader)
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

func FindProcessByPid(reader DirReader, pid int) (*Process, error) {
	procs, err := ListProcess(reader)
	if err != nil {
		return nil, err
	}

	for _, p := range procs {
		if p.Pid == pid {
			return p, nil
		}
	}

	return nil, fmt.Errorf("process with pid %d not found", pid)
}
