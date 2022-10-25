package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var time = time.Now()
		try := fmt.Sprint(time.Weekday(), ", ", time.Day(), time.Month(), time.Year())
		data := []byte(try)
		w.Write(data)
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			a := "Hello there"
			w.Write([]byte(a))
		} else {
			a := fmt.Sprintf("Hello, %s!", name)
			w.Write([]byte(a))
		}

	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/time", TimeHandler())
	mux.Handle("/hello", SayHelloHandler())
	// TODO: answer here
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
