package store

import (
	"labix.org/v2/mgo"
)

func ConnectToDb() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {panic(err)}

	return session, err
}

func ConnectToCollection(session *mgo.Session, collection_str string, keyString []string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)

	index := mgo.Index{
		Key:        keyString,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {panic(err)}

	return collection, err
}
