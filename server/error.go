package server

import "encoding/json"

type httpErr struct {
	Message string `json:"message"`
	Info    map[string]interface{}
}

func prepareErrorResponse(message string, code int, info map[string]interface{}) []byte {
	res := &httpErr{
		Message: message,
		Info:    info,
	}
	resBytes, _ := json.Marshal(res)
	return resBytes
}
