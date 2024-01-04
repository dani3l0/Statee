package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/processes"
)

func GetProcesses(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := processes.GetProcesses()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
