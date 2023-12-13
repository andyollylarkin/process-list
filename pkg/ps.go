package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Process struct {
	*os.Process
	Name string
}

func ListProcess() ([]*Process, error) {
	res := make([]*Process, 0)
	content, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, d := range content {
		pid, err := strconv.ParseInt(d.Name(), 10, 0)
		if err != nil {
			continue
		}
		path := filepath.Join("/proc", d.Name(), "exe")
		exeName, err := os.Readlink(path)
		if err != nil {
			continue
		}
		proc, err := os.FindProcess(int(pid))
		if err != nil {
			continue
		}

		res = append(res, &Process{
			Process: proc,
			Name:    filepath.Base(exeName),
		})
	}

	return res, nil
}

func FindProcessByNameContains(namePath string) ([]*Process, error) {
	res := make([]*Process, 0)
	procs, err := ListProcess()
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

func FindProcessByNameEqual(name string) ([]*Process, error) {
	res := make([]*Process, 0)
	procs, err := ListProcess()
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

func FindProcessByPid(pid int) (*Process, error) {
	procs, err := ListProcess()
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
