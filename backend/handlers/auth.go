package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"todo-app/backend/models"
)

type AuthHandler struct {
	users map[string]models.User
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		users: make(map[string]models.User),
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, exists := h.users[user.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	h.users[user.Username] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	storedUser, exists := h.users[user.Username]
	if !exists || storedUser.Password != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storedUser)
}

func (h *AuthHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/signup", h.SignUp).Methods("POST")
	router.HandleFunc("/signin", h.SignIn).Methods("POST")
}