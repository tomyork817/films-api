package handler

import (
	"encoding/json"
	vkfilms "github.com/bitbox228/vk-films-api"
	"net/http"
	"strconv"
)

// @Summary GetAllActors
// @Security ApiKeyAuth
// @Tags actors
// @Description get all actors
// @ID get-all-actors
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/actors [get]
func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.services.Actor.GetAll()
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, actors)
}

// @Summary CreateActor
// @Security ApiKeyAuth
// @Tags actors
// @Description create actor
// @ID create-actor
// @Accept  json
// @Produce  json
// @Param input body vk_films.Actor true "actor info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/actors [post]
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

// @Summary DeleteActor
// @Security ApiKeyAuth
// @Tags actors
// @Description delete actor
// @ID delete-actor
// @Accept  json
// @Produce  json
// @Param id path integer true "actor id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/actors [delete]
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

// @Summary UpdateActor
// @Security ApiKeyAuth
// @Tags actors
// @Description update actor
// @ID update-actor
// @Accept  json
// @Produce  json
// @Param id path integer true "actor id"
// @Param input body vk_films.UpdateActorInput true "actor updated info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/actors [put]
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
