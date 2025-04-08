package main

import (
	"fmt"
	"log"

	ps "github.com/andyollylarkin/process-list"
	"github.com/andyollylarkin/process-list/pkg/listers"
)

func main() {
	l := listers.NewLinuxLocalProcessLister()

	p, err := ps.ListProcess(l)
	if err != nil {
		log.Fatal(err)
	}

	for _, pp := range p {
		fmt.Printf("%s: %d\n", pp.Name, pp.Pid)
		for _, ip := range pp.Net {
			fmt.Printf("PsName: %s, PsPID: %d, Addr: %s\n", pp.Name, pp.Pid, ip.String())
		}
	}
}
