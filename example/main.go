package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	ps "github.com/andyollylarkin/process-list"
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

	start := time.Now()

	processes, err := ps.FindProcessByRegex(l, "load-example")
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	pprof.StopCPUProfile()

	fmt.Printf("Found %d processes in %s\n", len(processes), elapsed.String())
}
