package disks

import (
	"path"
	"statee/syslib/utils/sysfs"
)

type Disk struct {
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
	var result = Disk{}
	pathTo := path.Join("/sys/class/block", id)

	result.Space, _ = sysfs.CatInt(path.Join(pathTo, "size"))
	removable, _ := sysfs.CatInt(path.Join(pathTo, "removable"))
	result.Removable = removable != 0

	p := parseDiskStat(path.Join(pathTo, "stat"))
	result.ReadSpeed = p.ReadSectors
	result.WriteSpeed = p.WriteSectors
	result.ReadDelay = p.ReadDelay
	result.WriteDelay = p.WriteDelay
	result.ReadIOs = p.ReadIOs
	result.WriteIOs = p.WriteIOs

	return result, nil
}
