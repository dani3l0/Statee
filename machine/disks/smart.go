package disks

import (
	"fmt"
	"path"
	"strings"

	"github.com/anatol/smart.go"
)

type SMART struct {
	Device          string
	Name            string
	Read            uint64
	Written         uint64
	PowerCycles     uint16
	PowerOnHours    uint16
	UnsafeShutdowns uint16
}

func GetSmartData(device string) error {
	block := path.Join("/dev", device)

	if strings.HasPrefix(device, "sd") {
		dev, err := smart.OpenSata(block)
		defer dev.Close()
		sm, _ := dev.ReadSMARTData()
		fmt.Println(sm.Attrs)
		return err

	} else if strings.HasPrefix(device, "nvme") {
		dev, err := smart.OpenNVMe(block)
		defer dev.Close()
		sm, _ := dev.ReadSMART()
		fmt.Println(sm.AvailSpare)
		return err
	}

	dev, err := smart.Open(block)
	defer dev.Close()
	sm, _ := dev.ReadGenericAttributes()
	fmt.Println(sm)
	return err
}
