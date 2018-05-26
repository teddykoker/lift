package main

import (
	"encoding/json"
	"lift/models"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Bind adds all routes to the application route
func (app *App) Bind() {
	app.Router.GET("/api/", app.home)
	app.Router.POST("/api/login", app.login)
	app.Router.POST("/api/signup", app.signup)
	app.Router.GET("/api/user", app.user)
	app.Router.POST("/api/program", app.createProgram)
}

// respondWithError is a helper to respond with error messages in JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func (app *App) home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("API home"))
}

func (app *App) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse request")
		return
	}
	defer r.Body.Close()
	if err := app.Store.Users.Authenticate(&user); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect username or password")
		return
	}
	token, err := user.GenerateToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}
	user.Token = token
	user.Password = ""
	resp, err := json.Marshal(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (app *App) signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse request")
		return
	}
	defer r.Body.Close()

	if err := app.Store.Users.Insert(&user); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusConflict, "Username taken")
		return
	}
	token, err := user.GenerateToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}
	user.Token = token
	user.Password = ""
	resp, err := json.Marshal(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *App) user(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	log.Printf("received token: %s\n", tokenString)
	var user models.User
	if err := user.FromToken(tokenString); err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	resp, err := json.Marshal(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *App) createProgram(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	log.Printf("received token: %s\n", tokenString)
	var user *models.User
	if err := user.FromToken(tokenString); err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := app.Store.Users.FindByID(user.ID)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	var program models.Program
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&program); err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse request")
		return
	}
	defer r.Body.Close()
	user.Programs = append(user.Programs, &program)
	if err := app.Store.Users.Update(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not update user")
	}
	resp, err := json.Marshal(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
