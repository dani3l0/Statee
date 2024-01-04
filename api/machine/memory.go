package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/memory"
)

func GetMemory(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := memory.GetMemory()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
