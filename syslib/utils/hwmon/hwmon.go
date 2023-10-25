package hwmon

import (
	"errors"
	"os"
	"path"
	"strings"
)

var BasePath string = "/sys/class/hwmon/"

func Get(target string) (string, error) {
	hwmons, _ := os.ReadDir(BasePath)

	for _, h := range hwmons {
		id := h.Name()
		full_path := path.Join(BasePath, id, "name")
		label, err := os.ReadFile(full_path)

		if err != nil {
			continue
		}

		name := string(label)
		name = strings.ReplaceAll(name, "\n", "")
		name = strings.Trim(name, "")

		if name == target {
			return path.Join(BasePath, id), nil
		}
	}

	return "", errors.New("hwmon: no such device")
}
