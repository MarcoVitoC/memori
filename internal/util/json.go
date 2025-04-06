package util

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&data)
}

func WriteJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}

func WriteResponse(w http.ResponseWriter, code int, data any, errors any) error {
	type response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data	 any		`json:"data,omitempty"`
		Errors any    `json:"errors,omitempty"`
	}

	return WriteJSON(w, code, response{
		Code: code,
		Status: http.StatusText(code),
		Data: data,
		Errors: errors,
	})
}
