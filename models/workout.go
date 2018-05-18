package models

import "github.com/jmoiron/sqlx"

// A Workout represents one workout, which is part of a program,
// and consists of a sequence of exercises.
type Workout struct {
	ID        int `db:"workout_id"`
	Sequence  int `db:"sequence"`
	Exercises []*Exercise
	ProgramID int `db:"program_id"`
}

var workoutSchema = `
CREATE TABLE IF NOT EXISTS workout (
	workout_id SERIAL primary key,
	sequence   INT,
	program_id INT
);
`

// A WorkoutStore is used for loading and saving Workouts to the database
type WorkoutStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store WorkoutStore) Init() {
	store.DB.Exec(workoutSchema)
}

// Insert workout into database
func (store WorkoutStore) Insert(workout *Workout) error {
	return store.DB.QueryRow(`INSERT INTO exercise (sequence, program_id)
		VALUES ($1, $2) RETURNING exercise_id`, &workout.Sequence, &workout.ProgramID).Scan(&workout.ID)
}
