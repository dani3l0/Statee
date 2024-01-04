package cpu

import "statee/machine/utils/hwmon"

type Temperature struct {
	Name     string
	Now      int
	Meltdown int
}

// Parses hwmon values and returns available temperatures for CPU
func GetTemperatures() ([]Temperature, error) {
	var temperatures []Temperature
	hwmon_values, err := hwmon.GetCoretemp()
	if err != nil {
		return temperatures, err
	}

	cores := len(hwmon_values.Labels)

	// Kinda tricky as hwmon_values array lengths might be different
	for i := 0; i < cores; i++ {
		temperatures = append(temperatures, Temperature{
			Name:     hwmon_values.Labels[i],
			Now:      hwmon_values.Temps[i],
			Meltdown: hwmon_values.CritTemps[i],
		})
	}

	return temperatures, nil
}
