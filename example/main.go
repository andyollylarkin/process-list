package main

import (
	"fmt"
	"log"

	ps "gitlab.mindsw.io/migrate-core-libs/process-list"
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg/listers"
)

func main() {
	l := listers.NewLinuxSshProcessLister("core", "", nil, nil, "/Users/neekrasov/code/mind/files/id_rsa", "192.168.123.35", 22)

	p, err := ps.FindProcessByNameEqual(l, "mind_agent")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", p)

	// for _, pp := range p {
	// 	fmt.Printf("%s %d ", pp.Name, pp.Pid)
	// }
}
