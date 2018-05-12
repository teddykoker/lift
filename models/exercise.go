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
CREATE TABLE IF NOT EXISTS exercise (
	exercise_id SERIAL primary key,
	sets        INT,
	reps        INT,
	weight      REAL,
	movement    VARCHAR,
	sequence    INT,
	program_id  INT
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
	return store.DB.QueryRow(`INSERT INTO exercise (sets, reps, weight, movement, sequence, program_id) VALUES
		($1, $2, $3, $4, $5, $6) RETURNING exercise_id`,
		&exercise.Sets,
		&exercise.Reps,
		&exercise.Weight,
		&exercise.Movement,
		&exercise.Sequence,
		&exercise.ProgramID,
	).Scan(&exercise.ID)
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
