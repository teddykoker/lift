package mongostore

import (
	"errors"
	"lift/models"

	"github.com/globalsign/mgo/bson"
)

const collection = "users"

// GetUser returns the user with the given id
func (db *Mongostore) GetUser(id bson.ObjectId) (*models.User, error) {
	user := &models.User{}
	err := db.C(collection).FindId(id).One(user)
	return user, err
}

// GetUserByUsername returns the user with the given username
func (db *Mongostore) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := db.C(collection).Find(bson.M{"username": username}).One(user)
	return user, err
}

// InsertUser inserts a user into the store
func (db *Mongostore) InsertUser(user *models.User) error {
	count, err := db.C(collection).Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("user with username exists")
	}
	user.ID = bson.NewObjectId()
	return db.C(collection).Insert(user)
}

// UpdateUser updates given user using its id
func (db *Mongostore) UpdateUser(user *models.User) error {
	return db.C(collection).UpdateId(user.ID, user)
}
