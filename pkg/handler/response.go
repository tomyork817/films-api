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

func newOkResponse(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(response)
}
