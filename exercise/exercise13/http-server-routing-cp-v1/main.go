package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var time = time.Now()
		try := fmt.Sprint(time.Weekday(), ", ", time.Day(), time.Month(), time.Year())
		data := []byte(try)
		w.Write(data)
	})
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			a := "Hello there"
			w.Write([]byte(a))
		} else {
			a := fmt.Sprintf("Hello, %s!", name)
			w.Write([]byte(a))
		}
	})
}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
