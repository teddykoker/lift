package main

import (
	"database/sql"
)

// An Exercise represents a single exercise to be performed
type Exercise struct {
	Name string
}

type ExerciseStore struct {
	DB *sql.DB
}

func (store ExerciseStore) Get(int id) (Exercise, error) {
	var e Exercise
	err := store.DB.QueryRow("SELECT name FROM exercises WHERE id=$1", id).Scan(&e.Name)
	return (e, err)
}

func (store ExerciseStore) List() ([]Exercise, error) {
	rows, err := store.DB.Query("SELECT name FROM exercises")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var es []Exercise
	for rows.Next() {
		var e Exercise
		if err := rows.Scan(&e.Name); err != nil {
			return nil, err
		}
		es = append(es, e)
	}
	return es, nil
}
