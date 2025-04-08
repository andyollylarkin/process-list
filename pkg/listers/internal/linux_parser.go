package internal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg"
	"gitlab.mindsw.io/migrate-core-libs/process-list/utils"
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

		netPath := filepath.Join("/proc", strconv.Itoa(int(pid)), "net", "tcp")

		netFile, err := os.Open(netPath)
		if err != nil {
			continue
		}

		addresses, err := parseNetContent(netFile)
		if err != nil {
			continue
		}

		res = append(res, pkg.Process{
			Pid:  int(pid),
			Name: strings.ReplaceAll(procName, "\n", ""),
			Net:  addresses,
		})
	}

	return res, nil
}

func parseNetContent(content io.Reader) ([]net.Addr, error) {
	out := make([]net.Addr, 0)
	scanner := bufio.NewScanner(content)

	scanner.Scan() // skip first info line

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		ip, port, err := parseIpAndPort(fields[1])

		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(ip, port))
		if err != nil {
			continue
		}

		out = append(out, addr)
	}

	return out, nil
}

func parseIpAndPort(text string) (string, string, error) {
	parts := strings.Split(text, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid address format. Want ip:port in hex. Got %s", text)
	}

	ip, err := utils.CoventHexIpToAddress(parts[0])
	if err != nil {
		return "", "", err
	}
	port, err := utils.CoventHexPortToAddress(parts[1])
	if err != nil {
		return "", "", err
	}

	return ip, port, nil
}
