package network

import (
	"os"
	"statee/machine/utils"
	"strings"
)

var NET_PATH = "/sys/class/net"

type Interface struct {
	Name       string
	Rx         int
	Tx         int
	Speed      int
	MacAddress string
	Wired      bool
	State      string
}

// Find all interfaces and fetch information
func GetNetwork() []Interface {
	var interfaces []Interface
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
		interfaces = append(interfaces, GetInterface(name))
	}
	return interfaces
}

// Get specified interface information
func GetInterface(name string) Interface {
	rx, _ := utils.CatInt(NET_PATH, name, "statistics/rx_bytes")
	tx, _ := utils.CatInt(NET_PATH, name, "statistics/tx_bytes")
	speed, _ := utils.CatInt(NET_PATH, name, "speed")
	mac_address, _ := utils.Cat(NET_PATH, name, "address")
	carrier, _ := utils.CatInt(NET_PATH, name, "carrier")
	state, _ := utils.Cat(NET_PATH, name, "operstate")

	return Interface{
		Name:       name,
		Rx:         rx,
		Tx:         tx,
		Speed:      speed,
		MacAddress: mac_address,
		Wired:      carrier == 0,
		State:      state,
	}
}
