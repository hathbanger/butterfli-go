package models

import (
	//"encoding/json"

	//"labix.org/v2/mgo"
	"gopkg.in/mgo.v2/bson"
	//"log"
	//"sync"
	"fmt"
	//"time"
	"github.com/butterfli-go/store"
	"time"
)

type SearchTerm struct {
	Id 		bson.ObjectId 		`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       	`json:"time",bson:"time,omitempty"`
	Text		string           	`json:"text",bson:"text,omitempty"`
}

func NewSearchTerm(text string) *SearchTerm {
	s := new(SearchTerm)
	s.Id = bson.NewObjectId()
	s.Text = text

	return s
}

func (s *SearchTerm) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToSearchTermCollection(session, "searchTerms")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&SearchTerm{
		Id: s.Id,
		Timestamp: s.Timestamp,
		Text: s.Text})
	if err != nil {
		return err
	}

	return nil
}

func FindSearchTerm(text string) SearchTerm {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToSearchTermCollection(session, "searchTerms")
	if err != nil {
		panic(err)
	}

	searchTerm := SearchTerm{}
	err = collection.Find(bson.M{"text": text}).One(&searchTerm)
	if err != nil {
		fmt.Print("didn't work \n")
		fmt.Print(err)
		return searchTerm
	}

	return searchTerm
}