package handler

import (
	"github.com/bitbox228/vk-films-api/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-in", signIn)

	return mux
}
