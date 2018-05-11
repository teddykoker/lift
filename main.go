package main

import (
	"fmt"
	"lift/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// portPtr := flag.Int("port", 3001, "Port to run application on")
	// flag.Parse()

	// addr := fmt.Sprintf(":%d", *portPtr)

	// log.Printf("Listening on %s", addr)

	// app := NewApp("./app.db")
	// app.Run(addr)

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}
	store := models.NewDatastore(db)
	store.Movements.Init()

	err = store.Movements.Insert(&models.Movement{Name: "Bench"})
	if err != nil {
		log.Fatalln(err)
	}
	store.Movements.Insert(&models.Movement{Name: "Squat"})
	store.Movements.Insert(&models.Movement{Name: "Deadlift"})

	movements, err := store.Movements.List()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(movements)

	store.Exercises.Init()

	err = store.Exercises.Insert(&models.Exercise{
		Reps:       5,
		Sets:       5,
		MovementID: movements[0].ID,
	})

	if err != nil {
		log.Fatalln(err)
	}
	exercises, err := store.Exercises.List()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(exercises)
}
