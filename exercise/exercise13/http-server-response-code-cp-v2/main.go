package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}
	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// parameter
		// body
		name := request.URL.Query().Get("name")
		check := IsNameExists(name)

		if request.Method != "GET" {
			a := "Method is not allowed"
			writer.WriteHeader(405)
			writer.Write([]byte(a))
		} else if check == true {
			a := "Name is exists"
			writer.WriteHeader(200)
			writer.Write([]byte(a))

		} else if check == false {
			a := "Data not found"
			writer.WriteHeader(404)
			writer.Write([]byte(a))
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/students", CheckStudentName())

	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
