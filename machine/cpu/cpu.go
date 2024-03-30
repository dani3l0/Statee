package cpu

import (
	"statee/machine/utils"
	"strings"
)

type Cpu struct {
	Model   string
	Cores   int
	Threads int
	Cache   int
	Flags   []string
	Usage   []Usage
}

// Main function, get all cpu information
func GetCpu() Cpu {
	var cpu Cpu

	// Basic info
	cpuinfo, _ := utils.Cat("/proc/cpuinfo")
	cpu.Model, _ = utils.Grep(cpuinfo, "model name")
	cpu.Cache, _ = utils.GrepInt(cpuinfo, "cache size")

	// Flags & capabilities
	_flags, _ := utils.Grep(cpuinfo, "flags")
	flags := strings.Fields(_flags)
	selected_flags := "avx avx2 avx512f aes vmx ht svm smt"
	for _, flag := range flags {
		if strings.Contains(selected_flags, flag) {
			cpu.Flags = append(cpu.Flags, flag)
		}
	}

	// Load, temperatures, frequencies per each thread/core
	cpu.Usage = GetUsage()
	cpu.Threads = len(cpu.Usage)
	cores := 0
	for _, v := range cpu.Usage {
		if v.Core > cores {
			cores = v.Core
		}
	}
	cpu.Cores = 1 + cores

	return cpu
}
