package memory

import sysfs "statee/syslib/utils/sysfs"

type Memory struct {
	Used          int
	Available     int
	Total         int
	SwapUsed      int
	SwapAvailable int
	SwapTotal     int
}

func GetMemory() Memory {
	var memory Memory
	meminfo, _ := sysfs.Cat("/proc/meminfo")

	mem_total, _ := sysfs.GrepInt(meminfo, "MemTotal")
	mem_available, _ := sysfs.GrepInt(meminfo, "MemAvailable")
	memory.Total = mem_total / 1000
	memory.Available = mem_available / 1000
	memory.Used = memory.Total - memory.Available

	swap_total, _ := sysfs.GrepInt(meminfo, "SwapTotal")
	swap_free, _ := sysfs.GrepInt(meminfo, "SwapFree")
	memory.SwapTotal = swap_total / 1000
	memory.SwapAvailable = swap_free / 1000
	memory.SwapUsed = memory.SwapTotal - memory.SwapAvailable

	return memory
}
