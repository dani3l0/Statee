package network

import (
	"os"
	"path"
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
	ipath := path.Join(NET_PATH, name)

	rx, _ := utils.CatInt(path.Join(ipath, "statistics/rx_bytes"))
	tx, _ := utils.CatInt(path.Join(ipath, "statistics/tx_bytes"))
	speed, _ := utils.CatInt(path.Join(ipath, "speed"))
	mac_address, _ := utils.Cat(path.Join(ipath, "address"))
	wired, _ := utils.CatInt(path.Join(ipath, "carrier"))
	state, _ := utils.Cat(path.Join(ipath, "operstate"))

	return Interface{
		Name:       name,
		Rx:         rx,
		Tx:         tx,
		Speed:      speed,
		MacAddress: mac_address,
		Wired:      wired == 1,
		State:      state,
	}
}
