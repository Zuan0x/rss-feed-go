package main

import (
	"net/http"
	"strings"
)

// handlerGetUser is the handler for GET /users
func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithError(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		respondWithError(w, http.StatusUnauthorized, "invalid Authorization header")
		return
	}

	apiKey := authHeaderParts[1]
	
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
