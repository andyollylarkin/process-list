package pkg

import "net"

type NetworkState struct {
	ListenAddr net.Addr
	PublicAddr net.IP
	State      SocketState
	Network    string
}

type Process struct {
	Pid  int
	Name string
	Fds  []int
	Net  []NetworkState
}
