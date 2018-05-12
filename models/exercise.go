package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// An Exercise represents a single exercise to be performed
type Exercise struct {
	ID        int     `db:"exercise_id"`
	Sets      int     `db:"sets"`
	Reps      int     `db:"reps"`
	Weight    float64 `db:"weight"`
	Movement  string  `db:"movement"`
	Sequence  int     `db:"sequence"`
	ProgramID int     `db:"program_id"`
}

var exerciseSchema = `
CREATE TABLE exercise (
	exercise_id integer primary key,
	sets        integer,
	reps        integer,
	weight      real,
	movement    text,
	sequence    integer,
	program_id  integer
);
`

// An ExerciseStore is used for loading and saving Exercises to the database
type ExerciseStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store ExerciseStore) Init() {
	res, err := store.DB.Exec(exerciseSchema)
	fmt.Println(res, err)
}

// Insert exercise into database
func (store ExerciseStore) Insert(exercise *Exercise) error {
	_, err := store.DB.NamedExec(
		`INSERT INTO exercise (sets, reps, weight, movement, sequence, program_id) VALUES (
			:sets,
			:reps,
			:weight,
			:movement,
			:sequence,
			:program_id)`, exercise)
	return err
}

// Get gets exercise with id
func (store ExerciseStore) Get(id int) (Exercise, error) {
	e := Exercise{}
	err := store.DB.Get(&e, "SELECT * FROM exercise WHERE exercise_id=$1", id)
	return e, err
}

// List returns all exercises
func (store ExerciseStore) List() ([]Exercise, error) {
	es := []Exercise{}
	err := store.DB.Select(&es, "SELECT * FROM exercise")
	return es, err
}
