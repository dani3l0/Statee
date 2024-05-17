package machine

import (
	"encoding/json"
	"net/http"
	"statee/api/utils"
	"statee/machine/storage"
)

func GetStorage(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	rawData := storage.GetDisks()
	jsonData, _ := json.Marshal(rawData)

	w.Write(jsonData)
}
