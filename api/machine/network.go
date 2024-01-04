package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/network"
)

func GetNetwork(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := network.GetNetwork()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
