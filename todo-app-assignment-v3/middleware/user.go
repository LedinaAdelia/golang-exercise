package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil || isExpired(db.Sessions[c.Name]) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Error: "http: named cookie not present",
			})
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}
