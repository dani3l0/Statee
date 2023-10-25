package cpu

import "statee/syslib/utils/stat"

// Return usages in %% for each cpu
func GetCpuUsage() []float32 {
	cpu_stat := stat.GetCpu()
	var cpu_usage []float32

	for _, core := range cpu_stat {
		diff_idle := core.Idle - core.IdleLast
		diff_total := core.Total - core.TotalLast
		core_usage := 100.0 * float32(diff_idle) / float32(diff_total)
		core_usage = 100.0 - core_usage
		cpu_usage = append(cpu_usage, core_usage)
	}

	return cpu_usage
}
