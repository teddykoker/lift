package main

import (
	"lift/datastore"
	"log"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

// App struct where all variables will be defined
type App struct {
	Router *httprouter.Router
	Store  datastore.Datastore
}

const entry = "client/build"

var static = path.Join(entry, "static")

// NewApp returns initialized struct
func NewApp(dbURL string) *App {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir(static))

	// Catch all to serve index.html
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entry)
	})

	app := &App{
		Store:  datastore.New(dbURL),
		Router: router,
	}
	app.Bind()

	return app
}

// Run hosts the application at the given address
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
