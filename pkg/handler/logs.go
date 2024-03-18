package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func LogRequest(r *http.Request) {
	method := r.Method
	host := r.Host
	requestURL := r.URL.Path

	logrus.WithFields(logrus.Fields{
		"method":  method,
		"request": host + requestURL,
	}).Info("Received request")
}
