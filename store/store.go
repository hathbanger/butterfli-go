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

func ConnectToPostsCollection(session *mgo.Session, collection_str string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)
	index := mgo.Index{
		Key:        []string{"imgurl"},
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

func ConnectToAccountsCollection(session *mgo.Session, collection_str string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)
	index := mgo.Index{
		Key: []string{"title", "username"},
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

func ConnectToAccountCredCollection(session *mgo.Session, collection_str string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)
	index := mgo.Index{
		Key: []string{"account", "username"},
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

func ConnectToSearchTermCollection(session *mgo.Session, collection_str string) (*mgo.Collection, error) {
	collection := session.DB("test").C(collection_str)
	index := mgo.Index{
		Key:        []string{"text"},
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

