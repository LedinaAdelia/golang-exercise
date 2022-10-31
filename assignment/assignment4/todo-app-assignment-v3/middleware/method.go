package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorRes := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			json.NewEncoder(w).Encode(errorRes)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorRes := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			json.NewEncoder(w).Encode(errorRes)
			return
		} else if r.Body == http.NoBody {
			w.WriteHeader(http.StatusBadRequest)
			errorRes := model.ErrorResponse{
				Error: "Internal Server Error",
			}
			json.NewEncoder(w).Encode(errorRes)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorRes := model.ErrorResponse{
				Error: "Method is not allowed!",
			}
			json.NewEncoder(w).Encode(errorRes)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}
