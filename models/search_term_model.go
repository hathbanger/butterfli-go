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
	AccountId	string           	`json:"accountId",bson:"accountId,omitempty"`
}

func NewSearchTerm(accountId string, text string) *SearchTerm {
	s := new(SearchTerm)
	s.Id = bson.NewObjectId()
	s.Text = text
	s.AccountId = accountId

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
		Text: s.Text,
		AccountId: s.AccountId})
	if err != nil {
		return err
	}
	fmt.Print("Saved the SearchTerm!!!!!! \n")
	return nil
}

func FindAllSearchTerms(accountId string) []*SearchTerm {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToSearchTermCollection(session, "searchTerms")
	if err != nil {
		panic(err)
	}

	searchTerms := []*SearchTerm{}
	err = collection.Find(bson.M{"accountId": accountId}).All(&searchTerms)
	if err != nil {
		fmt.Print("didn't work \n")
		fmt.Print(err)
		return searchTerms
	}

	return searchTerms
}
func FindSearchTerm(accountId string, text string) SearchTerm {
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
	err = collection.Find(bson.M{"text": text, "accountId": accountId}).One(&searchTerm)
	if err != nil {
		fmt.Print("didn't work \n")
		fmt.Print(err)
		return searchTerm
	}

	return searchTerm
}