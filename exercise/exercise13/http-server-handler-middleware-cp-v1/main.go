package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Welcome to Student page"))
		}
	}) // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
