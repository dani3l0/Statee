package cpu

import (
	"strconv"
	"strings"
)

// Stat line
// cpu[x](0)  user(1)  nice{2}  system(3)  idle(4)  iowait(5)  irq(6)  softirq(7)  steal(8)  guest(9)
// cpu5	      13184    3        3630       405391   160        759     326         0         0

// {"cpu0": [27424, 3228]} etc.
var statCache = make(map[string][]int)

func getLoadFor(stat string, cpuId int) float32 {
	for _, line := range strings.Split(stat, "\n") {
		fields := strings.Fields(line)
		name := fields[0]
		if len(fields) > 0 && name == "cpu"+strconv.Itoa(cpuId) {
			idle := statDiff(name, 0, atoi(fields[4]))
			load := statDiff(name, 1, statUsed(fields))
			total := idle + load
			return 100.0 * float32(load) / float32(total)
		}
	}
	return -127
}

// Auto-compute difference between current and previous results
// for live data
// Sum stat fields considered as an actual load
func statDiff(cpuName string, mapId int, newValue int) int {
	if _, ok := statCache[cpuName]; !ok {
		statCache[cpuName] = []int{0, 0}
	}
	sth := newValue
	sthLast := statCache[cpuName][mapId]
	sthDiff := sth - sthLast
	statCache[cpuName][mapId] = sth
	return sthDiff
}

// Sum stat fields considered as an actual load
func statUsed(fields []string) int {
	used := 0
	used += atoi(fields[1])
	used += atoi(fields[2])
	used += atoi(fields[3])
	used += atoi(fields[6])
	used += atoi(fields[7])
	return used
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	return i
}
