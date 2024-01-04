package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/disks"
)

func GetDisks(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := disks.GetDisks()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
