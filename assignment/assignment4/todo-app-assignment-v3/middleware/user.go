package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Error: "http: named cookie not present",
				})
				return
			}
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}
