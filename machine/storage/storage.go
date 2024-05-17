package storage

import (
	"os"
	"strings"
)

// Only those devices are considered physical disks
var whitelist = []string{
	"nvme",   // NVMe drives (nvme0n1, nvme0n2, ...)
	"sd",     // SATA, SCSI, USB drives (sda, sdb, ...)
	"mmcblk", // SD Cards, eMMCs, etc. (mmcblk0, mmcblk1, ...)
	"vd",     // QEMU virtual disk devices (vda, vdb, ...)
}

// Find all appropriate disks
func GetDisks() []Disk {
	disks := []Disk{}
	devices, _ := os.ReadDir("/sys/block")
	GetMounts()

	for _, v := range devices {
		x := v.Name()
		for _, y := range whitelist {
			if strings.HasPrefix(x, y) {
				disk, err := GetDisk(x)
				if err != nil {
					continue
				}
				disks = append(disks, disk)
			}
		}
	}

	return disks
}
