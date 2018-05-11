package models

import (
	"github.com/jmoiron/sqlx"
)

// An Exercise represents a single exercise to be performed
type Exercise struct {
	ID   int `db:"exercise_id"`
	Sets int `db:"sets"`
	Reps int `db:"reps"`

	MovementID int `db:"movement_id"`
	Movement   *Movement
}

var schema = `
CREATE TABLE exercise (
	exercise_id integer primary key,
	sets integer,
	reps integer,
	movement_id integer
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

// // Insert movement into database
// func (store MovementStore) Insert(movement *Movement) error {
// 	_, err := store.DB.NamedExec(`INSERT INTO movement (name) VALUES (:name)`, movement)
// 	// TODO: set id of movement
// 	return err
// }

// Insert exercise into database
func (store ExerciseStore) Insert(exercise *Exercise) error {
	_, err := store.DB.NamedExec(`INSERT INTO exercise (sets, reps, movement_id) VALUES (:sets, :reps, :movement_id)`, exercise)
	return err
}

// Get gets exercise with id
func (store ExerciseStore) Get(id int) (Exercise, error) {
	e := Exercise{}
	err := store.DB.Get(&e, "SELECT * FROM exercise WHERE exercise_id=$1", id)
	if err != nil {
		return e, err
	}
	e.Movement = &Movement{}
	err = store.DB.Get(e.Movement, "SELECT * FROM movement WHERE movement_id=$1", e.MovementID)
	return e, err
}

// List returns all exercises
func (store ExerciseStore) List() ([]Exercise, error) {
	es := []Exercise{}
	err := store.DB.Select(&es, "SELECT * FROM exercise")
	if err != nil {
		return es, err
	}
	for i := range es {
		es[i].Movement = &Movement{Name: "test"}
		err = store.DB.Get(es[i].Movement, "SELECT * FROM movement WHERE movement_id=$1", es[i].MovementID)
		if err != nil {
			break
		}
	}
	return es, err
}
