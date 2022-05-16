package api

import "net/http"

func renderJson(w http.ResponseWriter, res []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}
