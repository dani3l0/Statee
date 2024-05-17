package memory

import utils "statee/machine/utils"

type Memory struct {
	Available     int
	Cached        int
	Total         int
	SwapAvailable int
	SwapTotal     int
}

// Parse meminfo
func GetMemory() Memory {
	var memory Memory
	meminfo, _ := utils.Cat("/proc/meminfo")

	// RAM
	mem_available, _ := utils.GrepInt(meminfo, "MemAvailable")
	memory.Available = mem_available * 1024 / 1000
	mem_cached, _ := utils.GrepInt(meminfo, "Cached")
	memory.Cached = mem_cached * 1024 / 1000
	mem_total, _ := utils.GrepInt(meminfo, "MemTotal")
	memory.Total = mem_total * 1024 / 1000

	// Swap or zram
	swap_total, _ := utils.GrepInt(meminfo, "SwapTotal")
	swap_free, _ := utils.GrepInt(meminfo, "SwapFree")
	memory.SwapTotal = swap_total * 1024 / 1000
	memory.SwapAvailable = swap_free * 1024 / 1000

	return memory
}
