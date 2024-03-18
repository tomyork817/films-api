package handler

import (
	_ "github.com/swaggo/http-swagger"
	"net/http"
)

func newApiHandler(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /actors", h.userIdentity(h.getAllActors))
	mux.HandleFunc("GET /actors/", h.userIdentity(h.getAllActors))

	mux.HandleFunc("POST /actors", h.adminIdentity(h.createActor))
	mux.HandleFunc("POST /actors/", h.adminIdentity(h.createActor))

	mux.HandleFunc("DELETE /actors", h.adminIdentity(h.deleteActor))

	mux.HandleFunc("PUT /actors", h.adminIdentity(h.updateActor))

	mux.HandleFunc("GET /films", h.userIdentity(h.getAllFilms))
	mux.HandleFunc("GET /films/", h.userIdentity(h.getAllFilms))

	mux.HandleFunc("GET /films/sort", h.userIdentity(h.getAllFilmsSorted))

	mux.HandleFunc("POST /films", h.adminIdentity(h.createFilm))
	mux.HandleFunc("POST /films/", h.adminIdentity(h.createFilm))

	mux.HandleFunc("DELETE /films", h.adminIdentity(h.deleteFilm))

	mux.HandleFunc("PUT /films", h.adminIdentity(h.updateFilm))

	mux.HandleFunc("GET /films/search", h.userIdentity(h.searchFilms))

	return mux
}
