package main

import (
	"lift/models"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// App struct where all variables will be defined
type App struct {
	Router *httprouter.Router
	Store  *models.Datastore
}

// Entry is the path of the production client to host
const Entry string = "client/build"

// NewApp returns initialized struct
func NewApp(dbURL string) *App {

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := httprouter.New()

	// Catch all to serve static files
	router.NotFound = http.FileServer(http.Dir("client/build"))

	app := &App{
		Store:  models.NewDatastore(db),
		Router: router,
	}
	app.Bind()

	return app
}

// Run hosts the application at the given address
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
