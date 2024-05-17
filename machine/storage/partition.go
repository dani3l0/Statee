package storage

import (
	"errors"
	"path"
	"statee/machine/utils"
	"strings"
	"syscall"
)

type Partition struct {
	Name       string
	Free       int
	Total      int
	Filesystem string
	Mountpoint string
	ReadOnly   bool
}

// Get full info about specified partition
func GetPartition(id string, sector_size int) (Partition, error) {
	partition := Partition{}
	pathTo := path.Join("/sys/class/block/", id)

	props, exists := systemMounts[id]
	if !exists {
		return partition, errors.New("no such partition")
	}

	// Space data
	statfs := syscall.Statfs_t{}
	syscall.Statfs(props.Mountpoint, &statfs)
	partition.Free = int(statfs.Bfree*uint64(statfs.Bsize)) / 1000
	partition.Total = int(statfs.Blocks*uint64(statfs.Bsize)) / 1000

	// If exists, fetch info about specified partition
	pretty_name := strings.ToTitle(path.Base(props.Mountpoint))
	partition.Name = pretty_name
	partition.Mountpoint = props.Mountpoint
	partition.Filesystem = props.Filesystem
	ro, _ := utils.CatInt(pathTo, "ro")
	partition.ReadOnly = ro == 1

	return partition, nil
}
