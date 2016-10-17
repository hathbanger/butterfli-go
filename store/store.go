package store

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

func ConnectToDb() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		//return err
		if err != nil {
			panic(err)
		}
	}

	return session, err
}

func ConnectToCollection(session *mgo.Session, collection_str string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	return collection, err
}