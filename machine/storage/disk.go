package storage

import (
	"os"
	"path"
	"statee/machine/utils"
	"strings"
)

type Disk struct {
	Name       string
	Model      string
	Space      int
	Removable  bool
	ReadSpeed  int
	WriteSpeed int
	IOPS       int
	AvgResp    int
	Partitions map[string]Partition
}

// Get disk with ids such as 'sda', 'sdb'
func GetDisk(id string) (Disk, error) {
	var disk = Disk{}
	disk.Name = id
	pathTo := path.Join("/sys/block", id)

	// Disk parameters
	sectors, _ := utils.CatInt(pathTo, "size")
	sector_size, _ := utils.CatInt(pathTo, "queue/hw_sector_size")
	disk.Space = sectors * sector_size / 1000
	removable, _ := utils.CatInt(pathTo, "removable")
	disk.Removable = removable != 0
	disk.Model, _ = utils.Cat(pathTo, "device/model")

	// Available partitions
	parts, _ := os.ReadDir("/sys/class/block")
	disk.Partitions = make(map[string]Partition)
	for _, part := range parts {
		if strings.HasPrefix(part.Name(), id) && len(part.Name()) > len(id) {
			part_info, err := GetPartition(part.Name(), sector_size)
			if err != nil {
				continue
			}
			disk.Partitions[part.Name()] = part_info
		}
	}

	// IO Stats
	p := parseDiskStat(pathTo)
	disk.AvgResp = (p.ReadDelay + p.WriteDelay) / 2
	disk.ReadSpeed = p.ReadSectors * sector_size / 1000
	disk.WriteSpeed = p.WriteSectors * sector_size / 1000
	disk.IOPS = p.ReadIOPS + p.WriteIOPS
	return disk, nil
}
