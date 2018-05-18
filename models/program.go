package models

import "github.com/jmoiron/sqlx"

// A Program represents a workout program created by a user
type Program struct {
	ID        int    `db:"program_id"`
	Name      string `db:"name"`
	Exercises []*Exercise
}

var programSchema = `
CREATE TABLE IF NOT EXISTS program (
	program_id SERIAL primary key,
	name       VARCHAR,
	user_id    INT
);
`

// A ProgramStore is used for loading and saving Programs to the database
type ProgramStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store ProgramStore) Init() {
	store.DB.Exec(programSchema)
}

// Insert exercise into database
func (store ProgramStore) Insert(program *Program) error {
	return store.DB.QueryRow(`INSERT INTO program (name) VALUES ($1) RETURNING program_id`,
		&program.Name,
	).Scan(&program.ID)
}

// Get gets program with id
func (store ProgramStore) Get(id int) (Program, error) {
	p := Program{}
	err := store.DB.Get(&p, "SELECT * FROM program WHERE program_id=1$", id)
	return p, err
}

// List returns all programs
func (store ProgramStore) List() ([]Program, error) {
	ps := []Program{}
	err := store.DB.Select(&ps, "SELECT * FROM program")
	return ps, err
}
