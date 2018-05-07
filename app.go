package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

// App struct where all variables will be defined
type App struct {
	Router *httprouter.Router
	DB     *sql.DB
}

// Entry is the path of the production client to host
const Entry string = "client/build"

// NewApp returns initialized struct
func NewApp(dbPath string) *App {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := httprouter.New()

	// Catch all to serve static files
	router.NotFound = http.FileServer(http.Dir("client/build"))

	app := &App{
		DB:     db,
		Router: router,
	}

	return app
}

// Run hosts the application at the given address
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
