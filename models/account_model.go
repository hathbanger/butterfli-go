package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/butterfli-go/store"
	"labix.org/v2/mgo"
)

type Account struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Title		string           	`json:"title",bson:"title,omitempty"`
	Username	string           	`json:"username",bson:"username,omitempty"`
	SearchTerms 	[]*SearchTerm		 `json:"searchterms",bson:"searchterms,omitempty"`
	FavoriteTerms 	[]*FavoriteTerm		 `json:"favoriteterms",bson:"favoriteterms,omitempty"`
	Posts 		[]*Post		 	`json:"posts",bson:"posts,omitempty"`
	AccountCreds    []AccountCreds		`json:"accountCreds",bson:"accountCreds,omitempty"`
}

func NewAccount(username string, title string) *Account {
	a := new(Account)
	a.Id = bson.NewObjectId()
	a.Timestamp = time.Now()
	a.Username = username
	a.Title = title

	return a
}

func (a *Account) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "accounts", []string{"title", "username"})
	if err != nil {
		panic(err)
	}
	err = collection.Insert(&Account{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Title: a.Title,
		Username: a.Username,
	})
	if err != nil {
		return err
	}
	return nil
}


func DeleteAccount(accountId string) error {
	session, err := store.ConnectToDb()
	if err != nil {
		panic(err)
	}
	collection := ConnectAccounts(session)
	err = collection.Remove(bson.M{"id": bson.ObjectIdHex(accountId)})
	if err != nil {
		panic(err)
	}
	return nil
}



func FindAccount(username string, title string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection := ConnectAccounts(session)
	account := Account{}
	err = collection.Find(bson.M{"username": username, "title": title}).One(&account)
	if err != nil {
		panic(err)
	}
	return &account, err
}

func FindAccountById(account_id string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection := ConnectAccounts(session)
	if err != nil {
		panic(err)
	}
	account := Account{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(account_id)}).One(&account)
	if err != nil {
		panic(err)
	}
	return &account, err
}

func (u *User) FindAccountByTitle(title string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "accounts", []string{"username", "title"})
	if err != nil {
		//panic(err)
		return &Account{}, err
	}
	account := Account{}
	err = collection.Find(bson.M{"username": u.Username, "title": title}).One(&account)
	if err != nil {
		panic(err)
	}

	return &account, err
}

func GetAllAccounts(username string) ([]*Account, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection := ConnectAccounts(session)
	accounts := []*Account{}
	err = collection.Find(bson.M{"username": username}).All(&accounts)
	return accounts, err
}


func ConnectAccounts(session *mgo.Session) *mgo.Collection{
	collection, err := store.ConnectToCollection(session, "accounts", []string{"username", "title"})
	if err != nil {
		panic(err)
	}
	return collection
}