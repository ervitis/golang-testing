package helpers

import (
	"encoding/json"
	"net/http"
	"time"
)

type errorResponse struct {
	Date    string `json:"date"`
	Message string `json:"message"`
}

func ResponseWithError(w http.ResponseWriter, code int, errorMessage string) {
	Response(w, code, &errorResponse{Date: time.Now().Format(time.RFC3339), Message: errorMessage})
}

func Response(w http.ResponseWriter, code int, data interface{}) {
	b, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
