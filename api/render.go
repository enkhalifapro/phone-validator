package api

import "net/http"

func renderJSON(w http.ResponseWriter, res []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}
