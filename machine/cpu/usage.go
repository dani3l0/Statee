package cpu

import (
	"os"
	"path"
	"regexp"
	"sort"
	"statee/machine/utils"
	"strconv"
	"strings"
)

type Usage struct {
	Id       int
	Core     int
	Load     float32
	Online   bool
	Freq     int
	BaseFreq int
	MinFreq  int
	MaxFreq  int
	Temp     float32
	Melt     float32
}

// Gets information for all cores for one cpu
func GetUsage() []Usage {
	usage := []Usage{}
	pattern := regexp.MustCompile(`^cpu\d+$`)
	stat, _ := utils.Cat("/proc/stat")
	sys := "/sys/devices/system/cpu"
	cpus, _ := os.ReadDir(sys)

	temps := getCoretemps()

	for _, v := range cpus {
		// Check if dir is named 'cpu[x]'
		name := v.Name()
		if !pattern.MatchString(name) {
			continue
		}

		// Runtime parameters
		u := Usage{}
		x := path.Join(sys, name)

		id := strings.Replace(name, "cpu", "", 1)
		u.Id = atoi(id)
		online, _ := utils.CatInt(x, "online")
		u.Online = online != 0

		// Online CPUs stuff
		if u.Online {
			u.Core, _ = utils.CatInt(x, "topology/core_id")
			u.Load = getLoadFor(stat, u.Id)

			// Frequencies
			u.Freq, _ = utils.CatInt(x, "cpufreq/scaling_cur_freq")
			u.BaseFreq, _ = utils.CatInt(x, "cpufreq/base_frequency")
			u.MinFreq, _ = utils.CatInt(x, "cpufreq/scaling_min_freq")
			u.MaxFreq, _ = utils.CatInt(x, "cpufreq/scaling_max_freq")
			u.Freq /= 1000
			u.BaseFreq /= 1000
			u.MinFreq /= 1000
			u.MaxFreq /= 1000
			u.Temp = temps.Temps[strconv.Itoa(u.Core)].Temp
			u.Melt = temps.Temps[strconv.Itoa(u.Core)].Melt
		}

		usage = append(usage, u)
	}

	// Sort ascending by Id
	sort.Slice(usage, func(i, j int) bool {
		return usage[i].Id < usage[j].Id
	})

	return usage
}
