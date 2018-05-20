package main

import (
	"lift/models"
	"log"
	"net/http"
	"path"

	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// App struct where all variables will be defined
type App struct {
	Router *httprouter.Router
	Store  *models.Datastore
}

const entry = "client/build"

var static = path.Join(entry, "static")

// NewApp returns initialized struct
func NewApp(dbURL string) *App {

	info, err := mgo.ParseURL(dbURL)
	if err != nil {
		log.Fatalf("Error parsing database URL: %q", err)
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	db := session.DB("lift")

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir(static))

	// Catch all to serve index.html
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entry)
	})

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
