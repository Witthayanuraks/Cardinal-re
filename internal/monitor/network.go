package monitor

import (
	"github.com/shirou/gopsutil/net"
)

type NetworkStats struct {
	ActiveConnections int
	OpenPorts         []int
}

func GetNetworkStats() NetworkStats {
	connections, _ := net.Connections("all") // testing doang,  package gopsutil/net
	openPorts := []int{}                     // port yang di buka

	return NetworkStats{
		ActiveConnections: len(connections),
		OpenPorts:         openPorts,
	}
}
