package internal

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/utils"
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

		netPathTcp4 := filepath.Join("/proc", strconv.Itoa(int(pid)), "net", "tcp")

		netFile, err := os.Open(netPathTcp4)
		if err != nil {
			continue
		}

		defer netFile.Close()

		netPathTcp6 := filepath.Join("/proc", strconv.Itoa(int(pid)), "net", "tcp6")

		netFile6, err := os.Open(netPathTcp6)
		if err != nil {
			continue
		}

		defer netFile6.Close()

		netPathUdp4 := filepath.Join("/proc", strconv.Itoa(int(pid)), "net", "udp")

		netFileUdp4, err := os.Open(netPathUdp4)
		if err != nil {
			continue
		}

		defer netFileUdp4.Close()

		netPathUdp6 := filepath.Join("/proc", strconv.Itoa(int(pid)), "net", "udp6")

		netFileUdp6, err := os.Open(netPathUdp6)
		if err != nil {
			continue
		}

		defer netFileUdp6.Close()

		fds, err := iterFdDir(int(pid))
		if err != nil {
			continue
		}

		addresses4, err := parseNetContent(netFile, filterSocketsFds(fds, pid), "tcp")
		if err != nil {
			continue
		}

		addresses6, err := parseNetContent(netFile6, filterSocketsFds(fds, pid), "tcp6")
		if err != nil {
			continue
		}

		addressesUdp4, err := parseNetContent(netFileUdp4, filterSocketsFds(fds, pid), "udp")

		if err != nil {
			continue
		}
		addressesUdp6, err := parseNetContent(netFileUdp6, filterSocketsFds(fds, pid), "udp6")
		if err != nil {
			continue
		}

		allAddresses := append(append(addresses4, addresses6...), append(addressesUdp4, addressesUdp6...)...)

		res = append(res, pkg.Process{
			Pid:  int(pid),
			Name: strings.ReplaceAll(procName, "\n", ""),
			Net:  allAddresses,
			Fds:  fds,
		})
	}

	return res, nil
}

func filterSocketsFds(fds []int, pid int64) []int {
	out := make([]int, 0)
	for _, fd := range fds {
		realName, err := os.Readlink(fmt.Sprintf("/proc/%d/fd/%d", pid, fd))
		if err != nil {
			continue
		}

		if strings.Contains(realName, "socket:") {
			sockFd, err := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(realName, "socket:["), "]"), 10, 0)
			if err != nil {
				continue
			}
			out = append(out, int(sockFd))
		}
	}

	return out
}

func iterFdDir(pid int) ([]int, error) {
	path := filepath.Join("/proc", strconv.Itoa(pid), "fd")

	out := make([]int, 0)

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		fdNo, errN := strconv.ParseInt(d.Name(), 10, 64)
		if errN != nil {
			return nil
		}

		out = append(out, int(fdNo))

		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

func parseNetContent(content io.Reader, fds []int, network string) ([]pkg.NetworkState, error) {
	out := make([]pkg.NetworkState, 0)
	scanner := bufio.NewScanner(content)

	scanner.Scan() // skip first info line

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 12 {
			continue
		}

		ip, port, err := parseIpAndPort(fields[1])
		if err != nil {
			return nil, err
		}

		socketFd, err := strconv.Atoi(fields[9])
		if err != nil {
			return nil, err
		}

		if socketFd == 0 {
			continue
		}

		if !contains(fds, socketFd) {
			continue
		}

		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(ip, port))
		if err != nil {
			continue
		}

		state := pkg.SocketState(fields[3])

		out = append(out, pkg.NetworkState{
			Addr:    addr,
			State:   pkg.SocketState(state),
			Network: network,
		})
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

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
