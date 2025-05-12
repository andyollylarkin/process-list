//go:build linux

package iface

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"slices"

	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/utils"
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

func (nso *NetSettingsObserver) DefaultGatewayIface() (*net.Interface, error) {
	rtable, err := nso.RoutingTable()
	if err != nil {
		return nil, err
	}

	for _, route := range rtable {
		if slices.Equal(route.Dest.IP, defaultGateway) {
			iface, err := net.InterfaceByName(route.Iface)
			if err != nil {
				return nil, err
			}

			return iface, nil
		}
	}

	return nil, fmt.Errorf("default gateway not found")
}

func (nso *NetSettingsObserver) RoutingTable() ([]Route, error) {
	rFile, err := nso.dirReader.Open("/proc/net/route")
	if err != nil {
		return nil, err
	}

	routes := make([]Route, 0)

	scanner := bufio.NewScanner(rFile)

	scanner.Scan() // Skip the header line

	/**
	Example routing table:
	Iface   Destination     Gateway         Flags   RefCnt  Use     Metric  Mask            MTU     Window  IRTT
	ens4    00000000        FEBFA8C0        0003    0       0       101     00000000        0       0       0
	**/

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 11 {
			continue
		}

		iface := fields[0]

		strDest, err := utils.CoventHexIpToAddress(fields[1])
		if err != nil {
			continue
		}

		strMask, err := utils.CoventHexIpToAddress(fields[7])
		if err != nil {
			continue
		}

		dest := &net.IPNet{
			IP:   net.ParseIP(strDest).To4(),
			Mask: net.IPMask(net.ParseIP(strMask).To4()),
		}

		strGateway, err := utils.CoventHexIpToAddress(fields[2])
		if err != nil {
			continue
		}

		flags, err := strconv.ParseUint(fields[3], 10, 0)
		if err != nil {
			continue
		}

		metric, err := strconv.ParseUint(fields[6], 10, 0)
		if err != nil {
			continue
		}

		mtu, err := strconv.ParseUint(fields[8], 10, 0)
		if err != nil {
			continue
		}

		routes = append(routes, Route{
			Iface:   iface,
			Dest:    dest,
			Gateway: net.ParseIP(strGateway).To4(),
			Flags:   uint32(flags),
			Metric:  uint32(metric),
			MTU:     uint32(mtu),
		})
	}

	return routes, nil
}
