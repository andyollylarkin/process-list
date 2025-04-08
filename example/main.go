package main

import (
	"fmt"
	"log"

	ps "gitlab.mindsw.io/migrate-core-libs/process-list"
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg/listers"
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
