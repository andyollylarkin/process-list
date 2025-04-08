package pkg

import "net"

type Process struct {
	Pid  int
	Name string
	Net  []net.Addr
}
