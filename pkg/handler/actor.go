package handler

import (
	"encoding/json"
	vkfilms "github.com/bitbox228/vk-films-api"
	"net/http"
	"strconv"
)

func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.services.Actor.GetAll()
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, actors)
}

func (h *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	var input vkfilms.Actor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Actor.Create(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.Actor.Delete(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	var input vkfilms.UpdateActorInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Actor.Update(id, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"status": "ok",
	})
}
