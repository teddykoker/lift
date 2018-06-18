package models

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
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
