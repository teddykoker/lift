package models

import "github.com/jmoiron/sqlx"

// An Movement represents a single Movement to be performed
type Movement struct {
	id   int    `db:"id"`
	Name string `db:"name"`
}

var movementSchema = `
CREATE TABLE movement (
	int id,
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
	e := Movement{}
	err := store.DB.Get(&e, "SELECT * FROM Movement WHERE name=$1", name)
	return e, err
}

// List returns all Movements
func (store MovementStore) List() ([]Movement, error) {
	es := []Movement{}
	err := store.DB.Select(&es, "SELECT * FROM Movement")
	return es, err
}
