package osinfo

import (
	"statee/machine/utils"
	"strconv"
	"strings"
)

type OsInfo struct {
	Name       string
	Version    string
	Hostname   string
	Arch       string
	Is64       bool
	Kernel     string
	Uptime     uint32
	LoadAvg    []float32
	LoadAvgMax int
}

func GetOsInfo() OsInfo {
	osinfo := OsInfo{}

	// Gets distro name & version
	os_release, _ := utils.Cat("/etc/os-release")
	for _, line := range strings.Split(os_release, "\n") {
		line = strings.ReplaceAll(line, "\"", "")
		entry := strings.Split(line, "=")
		switch entry[0] {
		case "NAME":
			osinfo.Name = entry[1]
		case "VERSION_ID":
			osinfo.Version = entry[1]
		}
	}

	// Gets some other stuff
	osinfo.Hostname, _ = utils.Cat("/proc/sys/kernel/hostname")
	osinfo.Kernel, _ = utils.Cat("/proc/sys/kernel/osrelease")
	osinfo.Arch, _ = utils.Cat("/proc/sys/kernel/arch")
	osinfo.Is64 = strings.Contains(osinfo.Arch, "64")
	uptime_raw, _ := utils.Cat("/proc/uptime")
	uptime_tmp, _ := strconv.ParseFloat(strings.Fields(uptime_raw)[0], 64)
	osinfo.Uptime = uint32(uptime_tmp)

	// Gets loadavg
	loadavg_raw, _ := utils.Cat("/proc/loadavg")
	for i, entry := range strings.Fields(loadavg_raw) {
		if i == 3 {
			break
		}
		parsed, _ := strconv.ParseFloat(entry, 64)
		osinfo.LoadAvg = append(osinfo.LoadAvg, float32(parsed))
	}

	// Return pretty info
	return osinfo
}
