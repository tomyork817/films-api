package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	response := errorResponse{Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error parsing error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func newOkResponse(w http.ResponseWriter, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func newOkResponseJson(w http.ResponseWriter, entity interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(entity)
	w.Write(jsonData)
}
