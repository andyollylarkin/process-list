package main

import (
	"fmt"
	"log"

	ps "github.com/andyollylarkin/process-list"
	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/listers"
)

func main() {
	l := listers.NewLinuxLocalProcessLister()

	p, err := ps.FindProcessByRegex(l, "^mds.+")
	if err != nil {
		log.Fatal(err)
	}

	for _, pp := range p {
		for _, ip := range pp.Net {
			if ip.State == pkg.SocketStateListen {
				fmt.Printf("PsName: %s, PsPID: %d, Addr: %s\t, Network: %s \t , RawState: %s\t ,State: %s\n",
					pp.Name, pp.Pid, ip.Addr.String(), ip.Network, ip.State, ip.State.String())
			}
		}
	}
}
