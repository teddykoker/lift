package models

import (
	"github.com/globalsign/mgo"
)

// A Datastore can be used to access all of the models in the database
type Datastore struct {
	DB    *mgo.Database
	Users *UserStore
}

// NewDatastore creates a new datastore
func NewDatastore(db *mgo.Database) *Datastore {
	d := &Datastore{DB: db}
	d.Users = &UserStore{DB: db}
	return d
}
