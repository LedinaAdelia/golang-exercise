package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// file, _ := api.usersRepo.ReadUser()
	// cred := model.Credentials{}
	// json.Unmarshal(file, &cred)
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	}
	r.ParseForm()
	creds := model.Credentials{
		Username: r.FormValue("username"),
		Password: r.Form.Get("password"),
	} // TODO: replace this

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here
	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	}

	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	}
	// Read usernmae and password request with FormValue.
	r.ParseForm()
	creds := model.Credentials{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	} // TODO: replace this

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here
	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Username or Password empty",
		})
		return
	}
	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
	// TODO: answer here
	id := uuid.New()

	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  id.String(),
		MaxAge: 18000,
		Path:   "/",
	})

	session := model.Session{
		Token:    id.String(),
		Username: creds.Username,
		Expiry:   time.Now().Add(5 * time.Hour),
	} // TODO: replace this

	err = api.sessionsRepo.AddSessions(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session_token")
	sessionToken := c.Value
	api.sessionsRepo.DeleteSessions(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:   "",
		Value:  "",
		MaxAge: 0,
		Path:   "/",
	})

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
