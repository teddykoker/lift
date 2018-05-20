package models

import (
	"errors"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const collection = "users"

// A User represents a single user of the site
type User struct {
	ID       int
	Username string
	Password string
	Token    string
	Programs []*Program
}

// A Program represents a workout program
type Program struct {
	Name    string
	Workout []*Workout
}

// A Workout represents one workout, which is part of a program,
// and consists of a sequence of exercises.
type Workout struct {
	Exercises []*Exercise
}

// An Exercise represents a single exercise to be performed
type Exercise struct {
	ID       int
	Sets     int
	Reps     int
	Weight   float64
	Movement string
}

// A UserStore is used for loading and saving Users to the database
type UserStore struct {
	DB *mgo.Database
}

// Insert user into database
func (store UserStore) Insert(user *User) error {
	// TODO: Check if user exists
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(pass)
	return store.DB.C(collection).Insert(user)
}

// Authenticate returns an error unless the username and password match those in the database
func (store UserStore) Authenticate(user *User) error {
	verify := &User{}
	if err := store.DB.C(collection).Find(bson.M{"username": user.Username}).One(verify); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(verify.Password), []byte(user.Password)); err != nil {
		return err
	}
	user = verify
	return nil
}

// GenerateToken provides a JWT token for the user
func (user *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
	})
	return token.SignedString([]byte("secret"))
}

// FromToken returns an error if the token cannot be parsed or is invalid
func (user *User) FromToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	user.Token = tokenString
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.ID = int(claims["id"].(float64)) // numbers in claims are always float64
		user.Username = claims["username"].(string)
		return nil
	}
	return errors.New("Invalid token")

}
