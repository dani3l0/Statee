package network

import (
	"os"
	"path"
	"statee/machine/utils"
	"strings"
)

var NET_PATH = "/sys/class/net"

type Interface struct {
	Rx         int
	Tx         int
	Speed      int
	MacAddress string
	Wired      bool
	Enabled    bool
}

// Find all interfaces and fetch information
func GetNetwork() map[string]Interface {
	interfaces := make(map[string]Interface)
	ifaces, _ := os.ReadDir(NET_PATH)

	for _, x := range ifaces {
		name := x.Name()

		ignored := name == "lo" || // loopback
			strings.HasPrefix(name, "docker") || // docker
			strings.HasPrefix(name, "veth") || // virtual
			strings.HasPrefix(name, "virt") || // virtual
			strings.HasPrefix(name, "br") // bridge

		if ignored {
			continue
		}
		interfaces[name] = GetInterface(name)
	}
	return interfaces
}

// Get specified interface information
func GetInterface(name string) Interface {
	rx, _ := utils.CatInt(NET_PATH, name, "statistics/rx_bytes")
	tx, _ := utils.CatInt(NET_PATH, name, "statistics/tx_bytes")
	speed, _ := utils.CatInt(NET_PATH, name, "speed")
	mac_address, _ := utils.Cat(NET_PATH, name, "address")
	_, wireless := os.Stat(path.Join(NET_PATH, name, "wireless"))
	state, _ := utils.Cat(NET_PATH, name, "operstate")

	return Interface{
		Rx:         rx,
		Tx:         tx,
		Speed:      speed,
		MacAddress: mac_address,
		Wired:      wireless != nil,
		Enabled:    state == "up",
	}
}
