package cpu

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"statee/machine/utils"
	"strings"
)

// Some data structs
type coretemp struct {
	Temp float32
	Melt float32
}
type coretemps struct {
	Ok      bool
	PerCore bool
	Temps   map[string]coretemp
}

// Gets CPU temperature for each core
func getCoretemps() coretemps {
	result := coretemps{}
	hwpath := "/sys/class/hwmon"
	hwmons, _ := os.ReadDir(hwpath)
	found := ""

	// Find proper hwmon id for CPU
	for _, hwmon := range hwmons {
		p := path.Join(hwpath, hwmon.Name())
		name, err := utils.Cat(p, "name")
		if err != nil || name != "coretemp" {
			continue
		}
		found = p
		break
	}

	if found == "" {
		return result
	}

	pattern := regexp.MustCompile(`^Core \d+$`)
	result.Temps = make(map[string]coretemp)
	sensors, _ := os.ReadDir(found)

	// Read temperatures, core by core
	for _, v := range sensors {
		_name := v.Name()
		name := strings.Split(_name, "_")
		fmt.Println(name)
		if len(name) != 2 || name[1] != "label" {
			continue
		}
		label, _ := utils.Cat(found, _name)
		if !pattern.MatchString(label) {
			continue
		}
		result.Ok = true
		result.PerCore = true
		temp_raw, _ := utils.CatInt(found, name[0]+"_input")
		melt_raw, _ := utils.CatInt(found, name[0]+"_crit")
		key := strings.Replace(label, "Core ", "", 1)
		result.Temps[key] = coretemp{
			Temp: float32(temp_raw) / 1000,
			Melt: float32(melt_raw) / 1000,
		}
	}

	return result
}
