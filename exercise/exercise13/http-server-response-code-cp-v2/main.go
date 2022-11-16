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
		name := request.URL.Query().Get("name")
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method is not allowed"))
		} else if !IsNameExists(name) {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Data not found"))
		} else {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("Name is exists"))
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
