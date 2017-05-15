package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/butterfli-go/store"
	"time"
	"fmt"
)

type FavoriteTerm struct {
	Id 		bson.ObjectId 		`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       	`json:"time",bson:"time,omitempty"`
	Text		string           	`json:"text",bson:"text,omitempty"`
	Account		string           	`json:"account",bson:"account,omitempty"`
	PostCount	int           		`json:"postcount",bson:"postcount,omitempty"`
	SinceTweetId	int64 			`json:"sincetweetid",bson:"sincetweetid,omitempty"`
}

func NewFavoriteTerm(account string, text string) *FavoriteTerm {
	var sinceTweetId int64
	s := new(FavoriteTerm)
	s.Id = bson.NewObjectId()
	s.Text = text
	s.Account = account
	s.SinceTweetId = sinceTweetId

	return s
}

func (s *FavoriteTerm) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "favoriteTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	searchTerm := &SearchTerm{
		Id: s.Id,
		Timestamp: s.Timestamp,
		Text: s.Text,
		Account: s.Account,
		SinceTweetId: s.SinceTweetId}

	err = collection.Insert(searchTerm)

	if err != nil {panic(err)}
	return nil

}

func FindAllFavoriteTerms(accountId string) []*FavoriteTerm {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "favoriteTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	favoriteTerms := []*FavoriteTerm{}
	err = collection.Find(bson.M{"account": accountId}).All(&favoriteTerms)
	if err != nil {panic(err)}

	return favoriteTerms
}

func FindFavoriteTerm(account string, text string) (*FavoriteTerm, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "favoriteTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	favoriteTerm := FavoriteTerm{}
	err = collection.Find(bson.M{ "account": account, "text": text}).One(&favoriteTerm)

	if err != nil {fmt.Print("\nerror! couldn't find the searchTerm\n")}

	fmt.Print("\nfavorite term: ")
	fmt.Print(favoriteTerm)
	//FindAllSearchTerms(account)
	return &favoriteTerm, err
}


func UpdateFavoriteTerm(favoriteTerm *FavoriteTerm, sinceTweetId int64) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}


	collection, err := store.ConnectToCollection(session, "favoriteTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	fmt.Print(sinceTweetId)

	colQuerier := bson.M{"id": favoriteTerm.Id}
	change := bson.M{"$set": bson.M{"sincetweetid": sinceTweetId}}
	err = collection.Update(colQuerier, change)
	if err != nil {
		fmt.Print("\nissssues!\n")
	}

	return err
}

func AddPostCountToFavoriteTerm(favoriteTerm *FavoriteTerm, count int) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "favoriteTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	colQuerier := bson.M{"id": favoriteTerm.Id}
	change := bson.M{"$inc": bson.M{"postcount": count}}
	err = collection.Update(colQuerier, change)
	if err != nil {fmt.Print("\nissssues!\n")}

	return err
}