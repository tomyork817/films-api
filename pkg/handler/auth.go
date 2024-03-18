package handler

import (
	"encoding/json"
	"github.com/bitbox228/vk-films-api"
	"net/http"
)

func newAuthHandler(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sign-in", h.signIn)
	mux.HandleFunc("POST /sign-in/", h.signIn)

	mux.HandleFunc("POST /sign-up", h.signUp)
	mux.HandleFunc("POST /sign-up/", h.signUp)

	return mux
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input vk_films.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" || input.Password == "" || (input.Role != vk_films.ADMIN && input.Role != vk_films.USER) {
		newErrorResponse(w, http.StatusBadRequest, "not all required fields are filled in")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"id":   id,
		"role": input.Role,
	})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input vk_films.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" || input.Password == "" {
		newErrorResponse(w, http.StatusBadRequest, "not all required fields are filled in")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Name, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponse(w, map[string]interface{}{
		"token": token,
	})
}
