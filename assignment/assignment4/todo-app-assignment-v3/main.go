package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Internal Server Error",
		})
	} else if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
	} else if _, ok := db.Users[user.Username]; ok {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username already exist",
		})
	} else {
		db.Users[user.Username] = user.Password
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{
			Username: user.Username,
			Message:  "Register Success",
		})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Internal Server Error",
		})
	} else if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
	} else if checkDB, ok := db.Users[user.Username]; !ok && checkDB != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Wrong User or Password!",
		})
	} else {
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   uuid.NewString(),
			Expires: time.Now().Add(5 * time.Hour),
			Path:    "/",
		}

		http.SetCookie(w, cookie)

		db.Sessions[cookie.Name] = model.Session{
			Username: user.Username,
			Expiry:   cookie.Expires,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{
			Username: user.Username,
			Message:  "Login Success",
		})

		fmt.Println("db: ", cookie.Value)
	}
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	var task model.Todo
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Internal Server Error",
		})
	} else {
		db.Task[db.Sessions["session_token"].Username] = append(db.Task[db.Sessions["session_token"].Username], model.Todo{
			Id:   uuid.NewString(),
			Task: task.Task,
			Done: task.Done,
		})
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.SuccessResponse{
			Username: db.Sessions["session_token"].Username,
			Message:  fmt.Sprintf("Task %s added!", task.Task),
		})
	}
	c, _ := r.Cookie("session_token")
	fmt.Println("add todo: ", c.Value)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	if len(db.Task) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Todolist not found!",
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db.Task[db.Sessions["session_token"].Username])
	}
	c, _ := r.Cookie("session_token")
	fmt.Println("list todo: ", c.Value)
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	db.Task[db.Sessions["session_token"].Username] = []model.Todo{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Username: db.Sessions["session_token"].Username,
		Message:  "Clear ToDo Success",
	})
	c, _ := r.Cookie("session_token")
	fmt.Println("clear todo: ", c.Value)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Username: db.Sessions["session_token"].Username,
		Message:  "Logout Success",
	})
	delete(db.Sessions, "session_token")
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

	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))

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
