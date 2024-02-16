package main

import (
	"fmt"
	"log"

	ps "gitlab.mindsw.io/migrate-core-libs/process-list"
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg/listers"
)

func main() {
	l := listers.NewWindowsSmbProcessLister("/opt/mind/bin/winexe", "admin", "Qwerty12344", "192.168.122.229", 445)

	p, err := ps.ListProcess(l)
	if err != nil {
		log.Fatal(err)
	}

	for _, pp := range p {
		fmt.Println(pp.Name, pp.Pid)
	}
}
