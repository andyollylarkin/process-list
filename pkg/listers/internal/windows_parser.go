package internal

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg"
)

func ParseWindows(executor Executor) ([]pkg.Process, error) {
	out, err := executor.Exec()
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(out)

	scanner := bufio.NewScanner(buff)

	outPs := make([]pkg.Process, 0)

	for scanner.Scan() {
		oneLine := scanner.Text()
		parts := strings.Fields(oneLine)

		if len(parts) < 2 {
			continue
		}

		pid, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			continue
		}

		outPs = append(outPs, pkg.Process{
			Name: strings.Trim(parts[1], " \n\t"),
			Pid: int(pid),
		})
	}

	return outPs, nil
}
