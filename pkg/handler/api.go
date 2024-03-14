package handler

import "net/http"

func newApiHandler(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /actors", h.userIdentity(h.getAllActors))
	mux.HandleFunc("GET /actors/", h.userIdentity(h.getAllActors))

	mux.HandleFunc("POST /actors", h.userIdentity(h.createActor))
	mux.HandleFunc("POST /actors/", h.userIdentity(h.createActor))

	return mux
}
