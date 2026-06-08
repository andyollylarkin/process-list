package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/net/iface"
	"github.com/andyollylarkin/process-list/utils"
)

func ParseLinux(reader pkg.DirReader, matchCondition func(int, string) bool) ([]pkg.Process, error) {
	res := make([]pkg.Process, 0)
	content, err := reader.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	v4ip, err := (iface.NewNetSettingsObserver(reader)).DefaultGatewayV4()
	if err != nil {
		return nil, err
	}

	netSnapshot, err := parseAllNetFiles(reader)
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

		if matchCondition != nil {
			if !matchCondition(int(pid), strings.ReplaceAll(procName, "\n", "")) {
				continue
			}
		}

		cmdlineContent, err := reader.ReadFile(filepath.Join("/proc", d.Name(), "cmdline"))
		if err != nil {
			continue
		}

		fds, err := iterFdDir(reader, int(pid))
		if err != nil {
			continue
		}

		socketInodes := filterSocketsFds(reader, fds, pid)
		allAddresses := lookupByInodes(netSnapshot, socketInodes)

		for i := range allAddresses {
			allAddresses[i].PublicAddr = v4ip
		}

		res = append(res, pkg.Process{
			Pid:     int(pid),
			Name:    strings.ReplaceAll(procName, "\n", ""),
			Net:     allAddresses,
			Fds:     fds,
			Cmdline: string(bytes.ReplaceAll([]byte(cmdlineContent), []byte{0}, []byte(" "))),
		})
	}

	return res, nil
}

func filterSocketsFds(reader pkg.DirReader, fds []int, pid int64) []int {
	out := make([]int, 0)
	for _, fd := range fds {
		realName, err := reader.ReadLink(fmt.Sprintf("/proc/%d/fd/%d", pid, fd))
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

func iterFdDir(reader pkg.DirReader, pid int) ([]int, error) {
	path := filepath.Join("/proc", strconv.Itoa(pid), "fd")

	out := make([]int, 0)

	d, err := reader.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, dir := range d {
		if dir.IsDir() {
			continue
		}
		fdNo, errN := strconv.ParseInt(dir.Name(), 10, 64)
		if errN != nil {
			continue
		}
		out = append(out, int(fdNo))
	}

	return out, nil
}

func parseAllNetFiles(reader pkg.DirReader) (map[int]pkg.NetworkState, error) {
	result := make(map[int]pkg.NetworkState)

	protocols := []struct {
		path    string
		network string
	}{
		{"/proc/self/net/tcp", "tcp"},
		{"/proc/self/net/tcp6", "tcp6"},
		{"/proc/self/net/udp", "udp"},
		{"/proc/self/net/udp6", "udp6"},
	}

	for _, p := range protocols {
		f, err := reader.Open(p.path)
		if err != nil {
			continue
		}
		entries, err := parseNetFile(f, p.network)
		_ = f.Close()
		if err != nil {
			continue
		}
		for inode, entry := range entries {
			result[inode] = entry
		}
	}

	return result, nil
}

func parseNetFile(content io.Reader, network string) (map[int]pkg.NetworkState, error) {
	fullContent, err := io.ReadAll(content)
	if err != nil {
		return nil, err
	}

	out := make(map[int]pkg.NetworkState)
	scanner := bufio.NewScanner(bytes.NewBuffer(fullContent))
	scanner.Scan() // skip header line

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 12 {
			continue
		}

		ip, port, err := parseIpAndPort(fields[1])
		if err != nil {
			continue
		}

		inode, err := strconv.Atoi(fields[9])
		if err != nil {
			continue
		}

		if inode == 0 {
			continue
		}

		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(ip, port))
		if err != nil {
			continue
		}

		out[inode] = pkg.NetworkState{
			ListenAddr: addr,
			State:      pkg.SocketState(fields[3]),
			Network:    network,
		}
	}

	return out, nil
}

func lookupByInodes(snapshot map[int]pkg.NetworkState, inodes []int) []pkg.NetworkState {
	out := make([]pkg.NetworkState, 0, len(inodes))
	for _, inode := range inodes {
		if entry, ok := snapshot[inode]; ok {
			out = append(out, entry)
		}
	}
	return out
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
