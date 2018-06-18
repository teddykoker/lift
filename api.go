package main

import (
	"encoding/json"
	"lift/models"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

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
	var login models.User
	json.NewDecoder(r.Body).Decode(&login)
	u, err := app.Store.GetUserByUsername(login.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(login.Password)) != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect username or password")
		return
	}
	token, err := u.GenerateToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}
	u.Token = token
	u.Password = ""
	resp, err := json.Marshal(u)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error marshalling json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (app *App) signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)
	if err := app.Store.InsertUser(&u); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusConflict, "Username taken")
		return
	}
	token, err := u.GenerateToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}
	u.Token = token
	u.Password = ""
	resp, err := json.Marshal(u)
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
	user, err := app.Store.GetUser(user.ID)
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
	if err := app.Store.UpdateUser(user); err != nil {
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
