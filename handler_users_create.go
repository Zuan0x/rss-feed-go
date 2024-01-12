package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Zuan0x/rss-feed-go/internal/database"
)

// handlerCreateUser is the handler for "POST /users" endpoint.

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name    string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("failed to decode request body: %v", err)
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		log.Printf("name is required")
		respondWithError(w, http.StatusBadRequest, "name is required")
		return
	}


	ctx := r.Context()

	user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		Name: req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Printf("failed to create user: %v", err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
