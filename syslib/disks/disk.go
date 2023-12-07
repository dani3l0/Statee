package disks

import (
	"path"
	"statee/syslib/utils/sysfs"
)

type Disk struct {
	Name       string
	Space      int
	Removable  bool
	ReadSpeed  int
	WriteSpeed int
	ReadDelay  int
	WriteDelay int
	ReadIOs    int
	WriteIOs   int
}

func GetDisk(id string) (Disk, error) {
	var disk = Disk{}
	disk.Name = id
	pathTo := path.Join("/sys/class/block", id)

	disk.Space, _ = sysfs.CatInt(path.Join(pathTo, "size"))
	removable, _ := sysfs.CatInt(path.Join(pathTo, "removable"))
	disk.Removable = removable != 0

	p := parseDiskStat(path.Join(pathTo, "stat"))
	disk.ReadSpeed = p.ReadSectors
	disk.WriteSpeed = p.WriteSectors
	disk.ReadDelay = p.ReadDelay
	disk.WriteDelay = p.WriteDelay
	disk.ReadIOs = p.ReadIOs
	disk.WriteIOs = p.WriteIOs

	return disk, nil
}
