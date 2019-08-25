package utils

import (
	"encoding/json"
	"net/http"
)

// SendResponse takes care of api response
func SendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
