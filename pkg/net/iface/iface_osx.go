//go:build darwin

package iface

import (
	"net"

	"github.com/andyollylarkin/process-list/pkg"
)

var defaultGateway = net.IPv4(0, 0, 0, 0).To4()

type Route struct {
	Iface   string
	Dest    *net.IPNet
	Gateway net.IP
	Flags   uint32
	Metric  uint32
	MTU     uint32
}

type NetSettingsObserver struct {
	dirReader pkg.DirReader
}

func NewNetSettingsObserver(dirReader pkg.DirReader) *NetSettingsObserver {
	return &NetSettingsObserver{
		dirReader: dirReader,
	}
}

func (nso *NetSettingsObserver) DefaultGatewayV4() (net.IP, error) {
	panic("not implemented. Tmp stub for build for OSX") // TODO: Implement this function
}

func (nso *NetSettingsObserver) DefaultGatewayV6() (net.IP, error) {
	panic("not implemented. Tmp stub for build for OSX") // TODO: Implement this function
}

func (nso *NetSettingsObserver) DefaultGatewayIface() (*net.Interface, error) {
	panic("not implemented. Tmp stub for build for OSX") // TODO: Implement this function
}

func (nso *NetSettingsObserver) RoutingTable() ([]Route, error) {
	panic("not implemented. Tmp stub for build for OSX") // TODO: Implement this function
}
