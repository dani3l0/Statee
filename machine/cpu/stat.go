package cpu

import (
	"statee/machine/utils"
	"strconv"
	"strings"
	"time"
)

type CPUStat struct {
	Idle      int
	IdleLast  int
	Total     int
	TotalLast int
}

var cpuStat []CPUStat

// Gets /proc/stat and calculates the differences for each cpu
func GetCpuStat() []CPUStat {
	stats := getCPUStat()

	if len(cpuStat) == 0 {
		cpuStat = stats
		time.Sleep(time.Second)
		stats = getCPUStat()
	}

	for i := 0; i < len(stats); i++ {
		cpuStat[i].IdleLast = cpuStat[i].Idle
		cpuStat[i].TotalLast = cpuStat[i].Total
		cpuStat[i].Idle = stats[i].Idle
		cpuStat[i].Total = stats[i].Total
	}

	return cpuStat
}

// Parses /proc/stat lines and returns for each cpu
func getCPUStat() []CPUStat {
	var stats []CPUStat
	stat, _ := utils.Cat("/proc/stat")
	lines := strings.Split(stat, "\n")

	for _, line := range lines {
		stat := CPUStat{}
		entry := strings.Fields(line)
		device := entry[0]
		if len(device) <= 3 || !strings.Contains(device, "cpu") {
			continue
		}

		// cpu user nice system idle iowait irq softirq steal guest
		idle, _ := strconv.Atoi(entry[4])
		stat.Idle = idle
		stat.Total = 0

		for _, s := range entry {
			n, _ := strconv.Atoi(s)
			stat.Total += n
		}
		stats = append(stats, stat)
	}

	return stats
}
