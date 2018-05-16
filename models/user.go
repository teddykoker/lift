package models

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// A User represents a single user of the site
type User struct {
	ID       int    `db:"user_id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password,omitempty"`
	Token    string `json:"token"`
}

var userSchema = `
CREATE TABLE IF NOT EXISTS users (
	user_id  SERIAL PRIMARY KEY,
	username VARCHAR,
	password VARCHAR
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
	if err := store.DB.QueryRow("SELECT EXISTS (SELECT user_id FROM users WHERE username=$1)", user.Username).Scan(&exists); err != nil {
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
	return store.DB.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING user_id", user.Username, user.Password).Scan(&user.ID)
}

// Authenticate returns an error unless the username and password match those in the database
func (store UserStore) Authenticate(user *User) error {
	var hash string
	if err := store.DB.QueryRow("SELECT user_id, password FROM users WHERE username=$1", user.Username).Scan(&user.ID, &hash); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password)); err != nil {
		return err
	}
	user.Password = hash
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
