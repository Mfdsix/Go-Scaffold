package helpers

import (
	"encoding/json"
	"go-scaffold/app/models"
	"net/http"
)

func HttpOk(w http.ResponseWriter, data interface{}, code int) {
	response := models.MResponse{
		Success: false,
		Status:  "success",
		Data:    data,
	}
	resp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func HttpOops(w http.ResponseWriter, message string, code int) {
	response := models.MResponse{
		Success: false,
		Status:  message,
	}
	resp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
