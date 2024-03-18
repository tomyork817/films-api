package handler

import (
	"encoding/json"
	vkfilms "github.com/bitbox228/vk-films-api"
	"net/http"
	"strconv"
)

func (h *Handler) getAllFilmsSorted(w http.ResponseWriter, r *http.Request) {
	var sort vkfilms.Sort
	sort.Type = vkfilms.SortType(r.URL.Query().Get("type"))
	sort.Order = vkfilms.SortOrder(r.URL.Query().Get("order"))

	films, err := h.services.Film.GetAllSorted(sort)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, films)
}

func (h *Handler) getAllFilms(w http.ResponseWriter, r *http.Request) {
	var sort vkfilms.Sort

	films, err := h.services.Film.GetAllSorted(sort)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, films)
}

func (h *Handler) createFilm(w http.ResponseWriter, r *http.Request) {
	var input vkfilms.CreateFilmInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Film.Create(input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.Film.Delete(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) updateFilm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	var input vkfilms.UpdateFilmInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Film.Update(id, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) searchFilms(w http.ResponseWriter, r *http.Request) {
	var search vkfilms.Search
	search.Type = vkfilms.SearchType(r.URL.Query().Get("type"))
	search.Fragment = r.URL.Query().Get("fragment")

	films, err := h.services.Film.GetSearch(search)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, films)
}
