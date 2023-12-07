package disks

import (
	"statee/syslib/utils/sysfs"
	"strconv"
	"strings"
)

type DiskStat struct {
	ReadIOs      int
	WriteIOs     int
	ReadSectors  int
	WriteSectors int
	ReadDelay    int
	WriteDelay   int
}

type DiskStatHist struct {
	Now  DiskStat
	Last DiskStat
}

var statsHist DiskStatHist

func parseDiskStat(path string) DiskStat {
	var result DiskStat
	stats := DiskStat{}
	contents, _ := sysfs.Cat(path)
	values := strings.Fields(contents)

	readIOs, _ := strconv.Atoi(values[0])
	stats.ReadIOs = readIOs
	result.ReadIOs = readIOs - statsHist.Last.ReadIOs

	writeIOs, _ := strconv.Atoi(values[4])
	stats.WriteIOs = writeIOs
	result.WriteIOs = writeIOs - statsHist.Last.WriteIOs

	readSectors, _ := strconv.Atoi(values[2])
	stats.ReadSectors = readSectors
	result.ReadSectors = readSectors - statsHist.Last.ReadSectors

	writeSectors, _ := strconv.Atoi(values[6])
	stats.WriteSectors = writeSectors
	result.ReadIOs = writeSectors - statsHist.Last.WriteSectors

	readDelay, _ := strconv.Atoi(values[3])
	stats.ReadDelay = readDelay
	result.ReadDelay = readDelay - statsHist.Last.ReadDelay

	writeDelay, _ := strconv.Atoi(values[7])
	stats.WriteDelay = writeDelay
	result.WriteDelay = writeDelay - statsHist.Last.WriteDelay

	statsHist.Last = statsHist.Now
	statsHist.Now = stats

	return result
}
