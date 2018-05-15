package models

import "github.com/jmoiron/sqlx"

// A Datastore can be used to access all of the models in the database
type Datastore struct {
	DB        *sqlx.DB
	Exercises *ExerciseStore
	Users     *UserStore
}

// NewDatastore creates a new datastore
func NewDatastore(db *sqlx.DB) *Datastore {
	d := &Datastore{DB: db}
	d.Exercises = &ExerciseStore{DB: db}
	d.Exercises.Init()
	d.Users = &UserStore{DB: db}
	d.Users.Init()
	return d
}
