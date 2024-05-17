package storage

import (
	"path"
	"statee/machine/utils"
	"strings"
)

type Mount struct {
	Path       string
	Mountpoint string
	Filesystem string
}

var systemMounts map[string]Mount

func GetMounts() {
	mounts := map[string]Mount{}
	mounts_raw, _ := utils.Cat("/proc/mounts")

	// [0]/dev/nvme0n1p3 [1]/home [2]ext4 [3]rw,seclabel,relatime [4]0 [5]0
	for _, mnt_raw := range strings.Split(mounts_raw, "\n") {
		mount := strings.Fields(mnt_raw)
		if !strings.HasPrefix(mount[0], "/dev") {
			continue
		}
		mounts[path.Base(mount[0])] = Mount{
			Path:       mount[0],
			Mountpoint: mount[1],
			Filesystem: mount[2],
		}
	}

	systemMounts = mounts
}
