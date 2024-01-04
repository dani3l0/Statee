package osinfo

import (
	"statee/machine/utils"
	"strconv"
	"strings"
)

type OsInfo struct {
	Name     string
	Version  string
	Hostname string
	Arch     string
	Is64     bool
	Kernel   string
	Uptime   uint32
	LoadAvg  []float32
}

func GetOsInfo() OsInfo {
	// Gets distro name & version
	var name, version string
	os_release, _ := utils.Cat("/etc/os-release")

	for _, line := range strings.Split(os_release, "\n") {
		entry := strings.Split(line, "=")
		switch entry[0] {
		case "NAME":
			name = entry[1]
		case "VERSION_ID":
			version = entry[1]
		}
	}

	// Gets some other stuff
	hostname, _ := utils.Cat("/proc/sys/kernel/hostname")
	kversion, _ := utils.Cat("/proc/sys/kernel/osrelease")
	arch, _ := utils.Cat("/proc/sys/kernel/arch")
	is64bit := strings.Contains(arch, "64")

	uptime_raw, _ := utils.Cat("/proc/uptime")
	uptime_tmp, _ := strconv.ParseFloat(strings.Fields(uptime_raw)[0], 64)
	uptime := uint32(uptime_tmp)

	// Gets loadavg
	var loadavg []float32
	loadavg_raw, _ := utils.Cat("/proc/loadavg")
	for i, entry := range strings.Fields(loadavg_raw) {
		if i > 2 {
			break
		}
		parsed, _ := strconv.ParseFloat(entry, 64)
		loadavg = append(loadavg, float32(parsed))
	}

	// Return pretty info
	return OsInfo{
		Name:     name,
		Version:  version,
		Hostname: hostname,
		Arch:     arch,
		Is64:     is64bit,
		Kernel:   kversion,
		Uptime:   uptime,
		LoadAvg:  loadavg,
	}
}
