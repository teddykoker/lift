package datastore

import (
	"lift/datastore/mongostore"
	"lift/models"

	"github.com/globalsign/mgo/bson"
)

// Datastore provides methods to interact with a database
type Datastore interface {
	// users
	GetUser(bson.ObjectId) (*models.User, error)
	GetUserByUsername(string) (*models.User, error)
	InsertUser(*models.User) error
	UpdateUser(*models.User) error
}

// New creates a new datastore
func New(dbURL string) Datastore {
	return mongostore.New(dbURL)
}
