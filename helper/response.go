package helper

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, code int, response map[string]any) {
	dataJson, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJson)
}
