package cpu

import (
	"statee/machine/utils"
	"strings"
)

type CpuInfo struct {
	Model         string
	CoresReal     int
	CacheSize     int
	SelectedFlags []string
}

// Parses /proc/cpuinfo and gets basic information about installed CPU
func GetCpuInfo() CpuInfo {
	var cpu_info CpuInfo
	const selected_flags = "avx2 aes vmx svm"

	cpuinfo_raw, _ := utils.Cat("/proc/cpuinfo")

	cpu_info.Model, _ = utils.Grep(cpuinfo_raw, "model name")
	cpu_info.CacheSize, _ = utils.GrepInt(cpuinfo_raw, "cache size")
	processors, _ := utils.Grep(cpuinfo_raw, "processor")
	cpu_info.CoresReal = len(processors)

	flags, _ := utils.Grep(cpuinfo_raw, "flags")
	for _, flag := range strings.Fields(flags) {
		if strings.Contains(selected_flags, flag) {
			cpu_info.SelectedFlags = append(cpu_info.SelectedFlags, flag)
		}
	}

	return cpu_info
}
