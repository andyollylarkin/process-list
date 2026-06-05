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
	l := listers.NewLinuxLocalProcessLister()
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err = pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	processes, err := ps.FindProcessByRegex(l, "^(mds|cs).+")
	if err != nil {
		log.Fatal(err)
	}

	pprof.StopCPUProfile()

	for _, pp := range processes {
		for _, ip := range pp.Net {
			if ip.State == pkg.SocketStateListen {
				fmt.Printf("PsName: %s, PsPID: %d, PubAddr: %s\t, \t ListenAddr %d\t ,Network: %s \t , RawState: %s\t ,State: %s\n",
					pp.Name, pp.Pid, ip.PublicAddr, ip.ListenAddr, ip.Network, ip.State, ip.State.String())
			}
		}
	}
}
