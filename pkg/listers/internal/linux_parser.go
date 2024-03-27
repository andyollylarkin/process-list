package internal

import (
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg"
)

func ParseLinux(reader DirReader) ([]pkg.Process, error) {
	res := make([]pkg.Process, 0)
	content, err := reader.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	for _, d := range content {
		pid, err := strconv.ParseInt(d.Name(), 10, 0)
		if err != nil {
			continue
		}
		path := filepath.Join("/proc", d.Name(), "comm")
		procName, err := reader.ReadFile(path)
		if err != nil {
			continue
		}

		res = append(res, pkg.Process{
			Pid:  int(pid),
			Name: strings.ReplaceAll(procName, "\n", ""),
		})
	}

	return res, nil
}
