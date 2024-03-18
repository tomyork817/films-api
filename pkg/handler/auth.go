package handler

import (
	"encoding/json"
	"github.com/bitbox228/vk-films-api"
	_ "github.com/swaggo/http-swagger"
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

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body vk_films.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body vk_films.SignInUserInput true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input vk_films.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.ValidateSignIn(); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
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
