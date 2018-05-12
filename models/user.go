package models

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// A User represents a single user of the site
type User struct {
	ID       int    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

var userSchema = `
CREATE TABLE user (
	user_id integer primary key,
	username text,
	password text
);
`

// A UserStore is used for loading and saving Users to the database
type UserStore struct {
	DB *sqlx.DB
}

// Init initializes table schema
func (store UserStore) Init() {
	store.DB.Exec(userSchema)
}

// Insert user into database
func (store UserStore) Insert(user *User) error {
	var exists bool
	if err := store.DB.QueryRow("SELECT EXISTS (SELECT user_id FROM user WHERE username=$1)", user.Username).Scan(&exists); err != nil {
		return err
	}
	if exists {
		return errors.New("User exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(pass)
	_, err = store.DB.Exec("INSERT INTO user (username, password) VALUES ($1, $2)", user.Username, user.Password)

	// TODO write id
	return err
}

// Authenticate returns an error unless the username and password match those in the database
func (store UserStore) Authenticate(user *User) error {
	var hash string
	if err := store.DB.QueryRow("SELECT user_id, password FROM user WHERE username=$1", user.Username).Scan(&user.ID, &hash); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password)); err != nil {
		return err
	}
	user.Password = hash
	return nil
}

// Token provides a JWT token for the user
func (user *User) Token() (string, error) {
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
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.ID = claims["id"].(int)
		user.Username = claims["username"].(string)
		return nil
	}
	return errors.New("Invalid token")

}
