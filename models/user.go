package models

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

const collection = "users"

// A User represents a single user of the site
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Username string        `json:"username"`
	Password string        `json:"password,omitempty"`
	Token    string        `bson:"-" json:"token"`
	Programs []*Program    `json:"programs"`
}

// A Program represents a workout program
type Program struct {
	Name    string     `json:"name"`
	Workout []*Workout `json:"workout"`
}

// A Workout represents one workout, which is part of a program,
// and consists of a sequence of exercises.
type Workout struct {
	Exercises []*Exercise `json:"exercises"`
}

// An Exercise represents a single exercise to be performed
type Exercise struct {
	Sets     int     `json:"sets"`
	Reps     int     `json:"reps"`
	Weight   float64 `json:"weight"`
	Movement string  `json:"movement"`
}

// A UserStore is used for loading and saving Users to the database
type UserStore struct {
	DB *mgo.Database
}

// Insert user into database
func (store UserStore) Insert(user *User) error {
	count, err := store.DB.C(collection).Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("user with username exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(pass)
	user.ID = bson.NewObjectId()
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

// FindByID returns the user with the given id
func (store UserStore) FindByID(id bson.ObjectId) (*User, error) {
	user := &User{}
	err := store.DB.C(collection).FindId(id).One(user)
	return user, err
}

// Update updates given user using its id
func (store UserStore) Update(user *User) error {
	return store.DB.C(collection).UpdateId(user.ID, user)
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
		// TODO: handle possibitly of runtime panic if claims doesn't contain
		// valid items
		user.ID = bson.ObjectIdHex(claims["id"].(string))
		user.Username = claims["username"].(string)
		return nil
	}
	return errors.New("Invalid token")

}
