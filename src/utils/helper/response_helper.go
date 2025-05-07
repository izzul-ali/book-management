package helper

import (
	"encoding/json"
	"net/http"

	"github.com/izzul-ali/book-management/src/model"
)

// Global Http Response
func ResponseJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(model.GlobalResponseModel{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// Global Http Response Error
func ResponseError(w http.ResponseWriter, status int, message string) {
	ResponseJSON(w, status, message, nil)
}
