package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func ErrorJSON(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError := &JSONError{
		Status: http.StatusText(code),
		Error:  error,
	}
	e := json.NewEncoder(w)
	e.Encode(jsonError)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	ErrorJSON(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
}

func UnsupportedMediaType(w http.ResponseWriter, r *http.Request, contentTypes []string) {
	ErrorJSON(w, fmt.Sprintf("content type must be one of %v", contentTypes), http.StatusUnsupportedMediaType)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorJSON(w, "page not found", http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter) {
	ErrorJSON(w, "internal server error", http.StatusInternalServerError)
}

func DecodingError(w http.ResponseWriter) {
	ErrorJSON(w, "body decoding error", http.StatusBadRequest)
}
