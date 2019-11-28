package format

import (
	"encoding/json"
	"net/http"
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/* 
Send sends response of a request 
*/
func Send(response http.ResponseWriter, status int, data Resp) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(data)
}
/* 
Message format for response
*/
func Message(status bool, message string, data interface{}) Resp {
	return Resp{"status": status, "message": message, "data": data}
}
