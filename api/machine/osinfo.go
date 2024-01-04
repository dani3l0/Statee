package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/osinfo"
)

func GetOsInfo(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := osinfo.GetOsInfo()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
