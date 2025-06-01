package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var password = os.Getenv("APP_PASSWORD")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var input map[string]string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if input["password"] == password {
			http.SetCookie(w, &http.Cookie{
				Name:    "session",
				Value:   "authenticated",
				Expires: time.Now().Add(30 * time.Minute),
			})
			http.Redirect(w, r, "/todo", http.StatusSeeOther)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "frontend/index.html")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}