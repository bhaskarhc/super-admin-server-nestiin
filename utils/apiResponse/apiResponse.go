package apiresponse

import (
	"encoding/json"
	"net/http"
)

func Message(statusCode int, status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "statusCode": statusCode, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
