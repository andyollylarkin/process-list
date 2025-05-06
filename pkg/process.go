package pkg

import "net"

type NetworkState struct {
	Addr    net.Addr
	State   SocketState
	Network string
}

type Process struct {
	Pid  int
	Name string
	Fds  []int
	Net  []NetworkState
}
