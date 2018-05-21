package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	portPtr := flag.Int("port", 3001, "Port to run application on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *portPtr)

	log.Printf("Listening on %s", addr)

	app := NewApp(os.Getenv("MONGODB_URI"))
	app.Run(addr)

	// store := app.Store

	// store.Exercises.Init()

	// err := store.Exercises.Insert(&models.Exercise{
	// 	Reps:      5,
	// 	Sets:      5,
	// 	Weight:    315.0,
	// 	Movement:  "Squat",
	// 	Sequence:  0,
	// 	ProgramID: 1,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// exercises, err := store.Exercises.List()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(exercises)

	// dude := &models.User{
	// 	Username: "dude",
	// 	Password: "password",
	// }
	// store.Users.Init()
	// err = store.Users.Insert(dude)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Print(*dude)
	// }

	// dude.Password = "password"
	// err = store.Users.Authenticate(dude)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("User Authenticated")
	// }

}
