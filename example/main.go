package main

import (
	"fmt"

	"github.com/andyollylarkin/process-list/pkg/net/iface"
	"github.com/andyollylarkin/process-list/pkg/readers"
)

func main() {
	r := readers.NewLocalDirReader()
	o := iface.NewNetSettingsObserver(r)
	ifs, _ := o.DefaultGatewayIface()
	fmt.Println(ifs.Name)
	fmt.Println(ifs.Addrs())
	// rr, err := o.RoutingTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, r := range rr {
	// 	_ = r
	// }
	// l := listers.NewLinuxLocalProcessLister()

	// p, err := ps.FindProcessByRegex(l, "^ostor-replica.+")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, pp := range p {
	// 	for _, ip := range pp.Net {
	// 		if ip.State == pkg.SocketStateListen {
	// 			fmt.Printf("PsName: %s, PsPID: %d, Addr: %s\t, Network: %s \t , RawState: %s\t ,State: %s\n",
	// 				pp.Name, pp.Pid, ip.Addr.String(), ip.Network, ip.State, ip.State.String())
	// 		}
	// 	}
	// }
}
