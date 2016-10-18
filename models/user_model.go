package models

import (
	"time"
	"fmt"

	"labix.org/v2/mgo/bson"
	"github.com/user-base/store"
)

type User struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Password	string           `json:"password",bson:"password,omitempty"`
}

func NewUser(username string, password string) *User {
	u := new(User)
	u.Id = bson.NewObjectId()
	u.Username = username
	u.Password = password

	return u
}

func (u *User) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&User{
		Id: u.Id,
		Timestamp: u.Timestamp,
		Username: u.Username,
		Password: u.Password})
	if err != nil {
		return err
	}

	return nil
}

func FindUser(username string) (User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	user := User{}
	err = collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return user, err
	}

	return user, err
}