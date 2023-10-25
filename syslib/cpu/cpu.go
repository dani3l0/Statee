package cpu

import (
	"os"
	"path"
	"statee/syslib/utils/sysfs"
	"strconv"
	"strings"
)

type Cpu struct {
	Model        string
	Cache        int
	CoresReal    int
	Flags        []string
	Cores        []Core
	Temperatures []Temperature
}

type Core struct {
	Usage   float32
	Online  bool
	Freq    int
	MinFreq int
	MaxFreq int
}

// Main function, get all cpu information
func Get() Cpu {
	var cpu Cpu
	cpu.Cores = GetCores()

	cpuinfo := GetCpuInfo()
	cpu.Model = cpuinfo.Model
	cpu.Flags = cpuinfo.SelectedFlags

	temperatures, _ := GetTemperatures()
	cpu.Temperatures = temperatures

	return cpu
}

// Gets information for all available cores
func GetCores() []Core {
	usage_all := GetCpuUsage()
	cores_count := len(usage_all)
	var cores = make([]Core, cores_count)

	// Usage in %%
	for i := 0; i < cores_count; i++ {
		cores[i].Usage = usage_all[i]
	}

	// Stuff from sysfs
	sys_data, _ := os.ReadDir("/sys/devices/system/cpu")
	for i, entry := range sys_data {
		_cpuid := strings.Replace(entry.Name(), "cpu", "", 1)
		cpuid, err := strconv.Atoi(_cpuid)
		if err != nil {
			continue
		}

		cpu_path := path.Join("/sys/devices/system/cpu", "cpu"+strconv.Itoa(cpuid))

		online_, _ := sysfs.CatInt(path.Join(cpu_path, "online"))
		cores[i].Online = online_ != 0

		freq_now, _ := sysfs.CatInt(path.Join(cpu_path, "cpufreq/scaling_cur_freq"))
		freq_min, _ := sysfs.CatInt(path.Join(cpu_path, "cpufreq/scaling_min_freq"))
		freq_max, _ := sysfs.CatInt(path.Join(cpu_path, "cpufreq/scaling_max_freq"))
		freq_now /= 1000
		freq_min /= 1000
		freq_max /= 1000
		cores[i].Freq = freq_now
		cores[i].MinFreq = freq_min
		cores[i].MaxFreq = freq_max
	}

	return cores
}
