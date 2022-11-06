package main

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}

	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	} else if _, exists := db.Users[user.Username]; exists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username already exist",
		})
		return
	} else if user.Username != "" && user.Password != "" {
		db.Users[user.Username] = user.Password
		json.NewEncoder(w).Encode(model.SuccessResponse{
			Username: user.Username,
			Message:  "Register Success",
		})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}
	v, exists := db.Users[user.Username]
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	} else if !exists || (v != user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Wrong User or Password!",
		})
		return
	} else if exists {
		id := uuid.New()

		http.SetCookie(w, &http.Cookie{
			Name:   "session_token",
			Value:  id.String(),
			MaxAge: 18000,
			Path:   "/",
		})

		db.Sessions["session_token"] = model.Session{
			Username: user.Username,
			Expiry:   time.Unix(18000, 0),
		}

		json.NewEncoder(w).Encode(model.SuccessResponse{
			Username: user.Username,
			Message:  "Login Success",
		})
	}
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	uuid := uuid.New()
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todo)

	username := db.Sessions["session_token"].Username
	db.Task[username] = append(db.Task[username], model.Todo{
		Id:   uuid.String(),
		Task: todo.Task,
		Done: todo.Done,
	})

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Username: db.Sessions["session_token"].Username,
		Message:  fmt.Sprintf("Task %s added!", todo.Task),
	})
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todo)
	username := db.Sessions["session_token"].Username
	if len(db.Task) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Todolist not found!",
		})
	} else if len(db.Task[username]) >= 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db.Task[username])
	}
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	username := db.Sessions["session_token"].Username
	db.Task[username] = []model.Todo{}
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Username: db.Sessions["session_token"].Username,
		Message:  "Clear ToDo Success",
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Username: db.Sessions["session_token"].Username,
		Message:  "Logout Success",
	})
	delete(db.Sessions, "session_token")
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	// username := db.Sessions["session_token"].Username
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
	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(http.HandlerFunc(ClearToDo)))
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
