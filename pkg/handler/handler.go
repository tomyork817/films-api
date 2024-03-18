package handler

import (
	_ "github.com/bitbox228/vk-films-api/docs"
	"github.com/bitbox228/vk-films-api/pkg/service"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
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

	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/swagger/doc.json")))

	return mux
}
