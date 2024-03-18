package handler

import (
	"encoding/json"
	vkfilms "github.com/bitbox228/vk-films-api"
	"net/http"
	"strconv"
)

// @Summary GetAllFilmsSorted
// @Security ApiKeyAuth
// @Tags films
// @Description get all films sorted
// @ID get-all-films-sorted
// @Accept  json
// @Produce  json
// @Param type query string false "type of sort (by rating/name/date)"
// @Param order query string false "order sort (asc/desc)"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films/sort [get]
func (h *Handler) getAllFilmsSorted(w http.ResponseWriter, r *http.Request) {
	var sort vkfilms.Sort
	LogRequest(r)

	sort.Type = vkfilms.SortType(r.URL.Query().Get("type"))
	sort.Order = vkfilms.SortOrder(r.URL.Query().Get("order"))

	films, err := h.services.Film.GetAllSorted(sort)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, films)
}

// @Summary GetAllFilmsSorted
// @Security ApiKeyAuth
// @Tags films
// @Description get all films
// @ID get-all-films
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films [get]
func (h *Handler) getAllFilms(w http.ResponseWriter, r *http.Request) {
	var sort vkfilms.Sort
	LogRequest(r)

	films, err := h.services.Film.GetAllSorted(sort)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseJson(w, films)
}

// @Summary CreateFilm
// @Security ApiKeyAuth
// @Tags films
// @Description create film
// @ID create-film
// @Accept  json
// @Produce  json
// @Param input body vk_films.CreateFilmInput true "film info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films [post]
func (h *Handler) createFilm(w http.ResponseWriter, r *http.Request) {
	var input vkfilms.CreateFilmInput
	LogRequest(r)

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

// @Summary DeleteFilm
// @Security ApiKeyAuth
// @Tags films
// @Description delete film
// @ID delete-film
// @Accept  json
// @Produce  json
// @Param id query integer true "film id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films [delete]
func (h *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)

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

// @Summary UpdateFilm
// @Security ApiKeyAuth
// @Tags films
// @Description update film
// @ID update-film
// @Accept  json
// @Produce  json
// @Param input body vk_films.UpdateFilmInput true "film updated info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films [put]
func (h *Handler) updateFilm(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)

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

// @Summary SearchFilms
// @Security ApiKeyAuth
// @Tags films
// @Description search films by film/actor name
// @ID search-films
// @Accept  json
// @Produce  json
// @Param fragment query string true "fragment to search"
// @Param type query string false "search by (film/actor)"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/films/search [get]
func (h *Handler) searchFilms(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)

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
