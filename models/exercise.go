package models

import "github.com/jmoiron/sqlx"

// An Exercise represents a single exercise to be performed
type Exercise struct {
	id   int `db:"id"`
	Sets int `db:"sets"`
	Reps int `db:"reps"`

	movementID int `db:"movement_id"`
	Movement   Movement
}

var schema = `
CREATE TABLE exercise (
	int id,
	int sets,
	int reps,
	int movement_id
);
`

// An ExerciseStore is used for loading and saving Exercises to the database
type ExerciseStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store ExerciseStore) Init() {
	store.DB.Exec(schema)
}

// Get gets exercise with id
func (store ExerciseStore) Get(id int) {
	e := Exercise{}

}

// List returns all exercises
func (store ExerciseStore) List() ([]Exercise, error) {
	es := []Exercise{}
	err := store.DB.Select(&es, "SELECT * FROM exercise")
	return es, err
}
