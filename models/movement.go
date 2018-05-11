package models

import (
	"github.com/jmoiron/sqlx"
)

// An Movement represents a single Movement to be performed
type Movement struct {
	ID   int    `db:"movement_id"`
	Name string `db:"name"`
}

var movementSchema = `
CREATE TABLE movement (
	movement_id integer primary key,
	name text
);
`

// An MovementStore is used for loading and saving Movements to the database
type MovementStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store MovementStore) Init() {
	store.DB.Exec(movementSchema)
}

// GetByName gets a single Movement with given name
func (store MovementStore) GetByName(name string) (Movement, error) {
	m := Movement{}
	err := store.DB.Get(&m, "SELECT * FROM movement WHERE name=$1", name)
	return m, err
}

// List returns all Movements
func (store MovementStore) List() ([]Movement, error) {
	es := []Movement{}
	err := store.DB.Select(&es, "SELECT * FROM movement")
	return es, err
}

// Insert movement into Database
func (store MovementStore) Insert(movement *Movement) error {
	_, err := store.DB.NamedExec(`INSERT INTO movement (name) VALUES (:name)`, movement)
	// TODO: set id of movement
	return err
}
