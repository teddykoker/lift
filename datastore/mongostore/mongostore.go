package mongostore

import (
	"log"

	"github.com/globalsign/mgo"
)

// Mongostore is a Datastore that uses MongoDB
type Mongostore struct {
	*mgo.Database
}

// New creates a new database connection
func New(dbURL string) *Mongostore {
	info, err := mgo.ParseURL(dbURL)
	if err != nil {
		log.Fatalf("Error parsing database URL: %s", err)
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	db := session.DB("")
	store := &Mongostore{db}
	log.Printf("Connected to database: %s", dbURL)
	return store
}
