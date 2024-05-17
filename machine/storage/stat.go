package storage

import (
	"statee/machine/utils"
	"strconv"
	"strings"
	"time"
)

type DiskStat struct {
	ReadIOPS     int
	WriteIOPS    int
	ReadSectors  int
	WriteSectors int
	ReadDelay    int
	WriteDelay   int
}

type DiskStatHist struct {
	Now      DiskStat
	Last     DiskStat
	TimeLast time.Time
}

var statsHist DiskStatHist

// @TODO: last data timestamp and diff calculation
// to provide value accurately
func parseDiskStat(path string) DiskStat {
	var result DiskStat
	stats := DiskStat{}
	contents, _ := utils.Cat(path, "stat")
	values := strings.Fields(contents)
	time_diff := time.Since(statsHist.TimeLast).Seconds()

	// IOPS
	readIOPS, _ := strconv.Atoi(values[0])
	stats.ReadIOPS = readIOPS
	result.ReadIOPS = perSecond((readIOPS - statsHist.Last.ReadIOPS), time_diff)
	writeIOPS, _ := strconv.Atoi(values[4])
	stats.WriteIOPS = writeIOPS
	result.WriteIOPS = perSecond((writeIOPS - statsHist.Last.WriteIOPS), time_diff)

	// RW Speed
	readSectors, _ := strconv.Atoi(values[2])
	stats.ReadSectors = readSectors
	result.ReadSectors = perSecond((readSectors - statsHist.Last.ReadSectors), time_diff)
	writeSectors, _ := strconv.Atoi(values[6])
	stats.WriteSectors = writeSectors
	stats.WriteSectors = perSecond((writeSectors - statsHist.Last.WriteSectors), time_diff)

	// Response time
	readDelay, _ := strconv.Atoi(values[3])
	stats.ReadDelay = readDelay
	result.ReadDelay = perSecond((readDelay - statsHist.Last.ReadDelay), time_diff)
	writeDelay, _ := strconv.Atoi(values[7])
	stats.WriteDelay = writeDelay
	result.WriteDelay = perSecond((writeDelay - statsHist.Last.WriteDelay), time_diff)

	// Update cache
	statsHist.Last = statsHist.Now
	statsHist.TimeLast = time.Now()
	statsHist.Now = stats

	return result
}

func perSecond(raw_value int, seconds float64) int {
	return int(float64(raw_value) / seconds)
}
