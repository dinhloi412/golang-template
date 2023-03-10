package api

import (
	"encoding/json"
	"net/http"
)

func responseJsonData(jsonData string, w http.ResponseWriter) {
	data := map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"data":    jsonData,
	}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func responseJsonSuccessEx(message string, w http.ResponseWriter) {

	if len(message) < 1 {
		message = "Request succeeded"
	}

	data := map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"info":    message,
	}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func responseJsonSuccess(w http.ResponseWriter) {
	responseJsonSuccessEx("Request succeeded", w)
}

func responseJsonErrorEx(message string, status int, w http.ResponseWriter) {

	if status == 0 {
		status = http.StatusNotAcceptable
	}

	if len(message) < 1 {
		message = "Request not acceptable"
	}

	data := map[string]interface{}{
		"success": false,
		"status":  status,
		"error":   message,
	}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func responseJsonError(w http.ResponseWriter) {
	responseJsonErrorEx("Request not acceptable", http.StatusNotAcceptable, w)
}

func apiAuthenticate(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if len(auth) < 1 {
		return false
	}

	return false
}

//var jsonSuccess = map[string]string{"test": "test", "data": "test1"}

func JsonSuccess(message string, data ...interface{}) interface{} {
	return jsonSuccess(message, data)
}

func jsonSuccess(message string, data ...interface{}) interface{} {
	obj := map[string]interface{}{
		"success": true,
		"code":    200,
		"message": message,
		"data":    data,
	}

	if data == nil || len(data) < 1 {
		delete(obj, "data")
	} else {
		if len(data) == 1 {
			obj["data"] = data[0]
		}
	}

	return obj
}

func JsonError(code int, message string) interface{} {
	return jsonError(code, message)
}

func jsonError(code int, message string) interface{} {
	obj := map[string]interface{}{
		"success": false,
		"code":    code,
		"errors": []string{
			message,
		},
	}

	return obj
}
