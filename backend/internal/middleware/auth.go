package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil || session.Values["authenticated"] == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if time.Now().Unix()-session.Values["last_activity"].(int64) > 1800 {
			delete(session.Values, "authenticated")
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		session.Values["last_activity"] = time.Now().Unix()
		session.Save(r, w)
		next.ServeHTTP(w, r)
	})
}