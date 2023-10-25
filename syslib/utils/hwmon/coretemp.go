package hwmon

import (
	"errors"
	"os"
	"path"
	"statee/syslib/utils/sysfs"
	"strings"
)

type Coretemp struct {
	Labels    []string
	Temps     []int
	CritTemps []int
}

// Scans /sys/class/hwmon/hwmon{n} directors and parses file contents
func GetCoretemp() (Coretemp, error) {
	coretemp := Coretemp{}
	coretemp_path, err := Get("coretemp")
	if err != nil {
		return coretemp, errors.New("coretemp: no such device")
	}

	contents, _ := os.ReadDir(coretemp_path)

	for i := 0; i < len(contents); i++ {
		f := contents[i]
		name := f.Name()
		temp := strings.Split(name, "_")

		if len(temp) == 2 && strings.HasPrefix(name, "temp") {
			sensor_type := (strings.Split(name, "_"))[1]

			switch sensor_type {
			case "label":
				value, _ := sysfs.Cat(path.Join(coretemp_path, name))
				coretemp.Labels = append(coretemp.Labels, value)

			case "input":
				value, _ := sysfs.CatInt(path.Join(coretemp_path, name))
				if value > 200 {
					value /= 1000
				}
				coretemp.Temps = append(coretemp.Temps, value)

			case "crit":
				value, _ := sysfs.CatInt(path.Join(coretemp_path, name))
				if value > 200 {
					value /= 1000
				}
				coretemp.CritTemps = append(coretemp.CritTemps, value)
			}
		}
	}

	return coretemp, nil
}
