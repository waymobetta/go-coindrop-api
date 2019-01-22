package utils

import (
	"encoding/json"
	"net/http"
)

// Message handles setting universal format of messaging for API responses
func Message(status bool, message interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond handles setting universal format of structure for API responses
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
