package handler

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) userIdentity(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		_, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		next(w, r)
	}
}

func (h *Handler) adminIdentity(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userRole, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		if userRole != vkfilms.ADMIN {
			newErrorResponse(w, http.StatusUnauthorized, "user is not admin")
			return
		}

		next(w, r)
	}
}
