package handler

import (
	"github.com/bitbox228/vk-films-api/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/auth/", http.StripPrefix("/auth", newAuthHandler(h)))
	mux.Handle("/api/", http.StripPrefix("/api", newApiHandler(h)))

	return mux
}
