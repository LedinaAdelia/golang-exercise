package main

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorRes := model.ErrorResponse{
			Error: "Username or Password empty",
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	} else if v, exists := db.Users[user.Username]; exists {
		w.WriteHeader(http.StatusConflict)
		errorRes := model.ErrorResponse{
			Error: "Username already exist",
		}
		json.NewEncoder(w).Encode(errorRes)
		fmt.Println(v)
		return
	} else if user.Username != "" && user.Password != "" {
		db.Users[user.Username] = user.Password
		response := model.SuccessResponse{
			Username: user.Username,
			Message:  "Register Success",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	v, exists := db.Users[user.Username]

	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorRes := model.ErrorResponse{
			Error: "Username or Password empty",
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	} else if !exists || (v != user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		errorRes := model.ErrorResponse{
			Error: "Wrong User or Password!",
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	} else if exists {
		id := uuid.New()
		cookie := &http.Cookie{
			Name:   "session_token",
			Value:  id.String(),
			Path:   "/",
			MaxAge: 18000,
		}

		a := model.Session{
			Username: user.Username,
			Expiry:   time.Unix(18000, 0),
		}
		db.Sessions[cookie.Name] = a
		http.SetCookie(w, cookie)

		response := model.SuccessResponse{
			Username: user.Username,
			Message:  "Login Success",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// for _, c := range r.Cookies() {
	// 	if c.Name != "session_token" {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// errorRes := model.ErrorResponse{
	// 	Error: "http: named cookie not present",
	// }
	// json.NewEncoder(w).Encode(errorRes)
	// return
	// 	}
	// }
	fmt.Println("3")
	// c, err := r.Cookie("session_token")
	// fmt.Println(c, err)
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		errorRes := model.ErrorResponse{
	// 			Error: "http: named cookie not present",
	// 		}
	// 		json.NewEncoder(w).Encode(errorRes)
	// 		return
	// 	}
	// }
	// sessionToken := c.Value
	// fmt.Println(sessionToken)
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))
	// TODO: answer here

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
