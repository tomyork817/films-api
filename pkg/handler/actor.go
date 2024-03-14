package handler

import "net/http"

func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(userCtx).(int)
	newOkResponse(w, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(userCtx).(int)
	newOkResponse(w, map[string]interface{}{
		"id": id,
	})
}
