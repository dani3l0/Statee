package disks

import (
	"path"
	"statee/machine/utils"
)

type Disk struct {
	Name       string
	Model      string
	Space      int
	Removable  bool
	ReadSpeed  int
	WriteSpeed int
	ReadDelay  int
	WriteDelay int
	ReadIOs    int
	WriteIOs   int
}

// Get disk with ids such as 'sda', 'sdb'
func GetDisk(id string) (Disk, error) {
	var disk = Disk{}
	disk.Name = id
	pathTo := path.Join("/sys/class/block", id)

	disk.Space, _ = utils.CatInt(pathTo, "size")
	removable, _ := utils.CatInt(pathTo, "removable")
	disk.Removable = removable != 0
	disk.Model, _ = utils.Cat(pathTo, "device/model")

	p := parseDiskStat(path.Join(pathTo, "stat"))
	disk.ReadSpeed = p.ReadSectors
	disk.WriteSpeed = p.WriteSectors
	disk.ReadDelay = p.ReadDelay
	disk.WriteDelay = p.WriteDelay
	disk.ReadIOs = p.ReadIOs
	disk.WriteIOs = p.WriteIOs

	return disk, nil
}
