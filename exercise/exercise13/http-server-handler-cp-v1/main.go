package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var time = time.Now()
		try := fmt.Sprint(time.Weekday(), ", ", time.Day(), time.Month(), time.Year())
		data := []byte(try)
		writer.Write(data)
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
