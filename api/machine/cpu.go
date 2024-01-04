package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/cpu"
)

func GetCpu(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := cpu.GetCpu()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
