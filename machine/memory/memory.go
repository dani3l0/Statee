package memory

import utils "statee/machine/utils"

type Memory struct {
	Used          int
	Available     int
	Cached        int
	Total         int
	SwapUsed      int
	SwapAvailable int
	SwapTotal     int
}

// Parse meminfo
func GetMemory() Memory {
	var memory Memory
	meminfo, _ := utils.Cat("/proc/meminfo")

	// RAM
	mem_available, _ := utils.GrepInt(meminfo, "MemAvailable")
	memory.Available = mem_available / 1000
	mem_cached, _ := utils.GrepInt(meminfo, "Cached")
	memory.Cached = mem_cached / 1000
	mem_total, _ := utils.GrepInt(meminfo, "MemTotal")
	memory.Total = mem_total / 1000
	memory.Used = memory.Total - memory.Available

	// Swap or zram
	swap_total, _ := utils.GrepInt(meminfo, "SwapTotal")
	swap_free, _ := utils.GrepInt(meminfo, "SwapFree")
	memory.SwapTotal = swap_total / 1000
	memory.SwapAvailable = swap_free / 1000
	memory.SwapUsed = memory.SwapTotal - memory.SwapAvailable

	return memory
}
