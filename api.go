package main

import (
	"encoding/json"
	"lift/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Bind adds all routes to the application route
func (app *App) Bind() {
	app.Router.GET("/api/", app.home)
	app.Router.POST("/api/login", app.login)
	app.Router.POST("/api/users", app.signup)
}

func (app *App) home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("API home"))
}

func (app *App) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	if err := app.Store.Users.Authenticate(&user); err != nil {
		log.Println("invalid credential")
	}
	token, err := user.Token()
	if err != nil {
		log.Println(err)
	}
	resp, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (app *App) signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	log.Println(user)

	if err := app.Store.Users.Insert(&user); err != nil {
		// user exists
		log.Println(err)
		return
	}
	token, err := user.Token()
	if err != nil {
		log.Println(err)
	}
	resp, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
