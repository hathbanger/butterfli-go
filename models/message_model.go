package models

import (
	"time"

	"labix.org/v2/mgo/bson"
	"github.com/user-boiler/store"
)

type Message struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Msg	string           `json:"msg",bson:"msg,omitempty"`
}

func NewMessage(username string, message string) *Message {
	m := new(Message)
	m.Id = bson.NewObjectId()
	m.Username = username
	m.Msg = message

	return m
}

func (m *Message) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "messages")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&Message{Id: m.Id,
		Timestamp: m.Timestamp,
		Username: m.Username,
		Msg: m.Msg})
	if err != nil {
		return err
	}

	return nil
}
