package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	ps "github.com/andyollylarkin/process-list"
	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/listers"
)

func main() {
	// l := listers.NewLinuxSshProcessLister("root", "P@$$w0rd", []byte{}, []byte{}, "", "192.168.191.231", 22)
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	ll := listers.NewLinuxLocalProcessLister()

	p, err := ps.FindProcessByRegex(ll, "^mds.+")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StopCPUProfile()

	for _, pp := range p {
		for _, ip := range pp.Net {
			if ip.State == pkg.SocketStateListen {
				fmt.Printf("PsName: %s, PsPID: %d, Addr: %s\t, Network: %s \t , RawState: %s\t ,State: %s\n",
					pp.Name, pp.Pid, ip.Addr.String(), ip.Network, ip.State, ip.State.String())
			}
		}
	}
}
