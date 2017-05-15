package models

import (
	"time"
	// "fmt"

	"labix.org/v2/mgo/bson"
	"github.com/butterfli-go/store"
	//"github.com/go-blog/models"
	//"github.com/labstack/gommon/log"

)

type User struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Password	string           `json:"password",bson:"password,omitempty"`
	PostIds		[]string 	`json:"post_ids",bson:"post_ids,omitempty"`
	posts 		[]*Post 	`json:"posts",bson:"posts,omitempty"`
}

func (u *User) GetAllPosts() error {
	for _, post_id := range(u.PostIds) {
		post, err := FindPostById(post_id)
		if err != nil {
			return err
		}

		u.posts = append(u.posts, post)
	}

	return nil
}

func NewUser(username string, password string) *User {
	u := new(User)
	u.Id = bson.NewObjectId()
	u.Username = username
	u.Password = password
	u.PostIds = []string{}

	return u
}

func (u *User) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "users", []string{"users"})
	if err != nil {panic(err)}

	err = collection.Insert(&User{
		Id: u.Id,
		Timestamp: u.Timestamp,
		Username: u.Username,
		Password: u.Password,
		PostIds: u.PostIds})
	if err != nil {return err}

	return nil
}

func FindUser(username string) (User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "users", []string{"users"})
	if err != nil {panic(err)}

	user := User{}
	err = collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return user, err
	}

	return user, err
}


func GetAllUsers() ([]*User, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "users", []string{"users"})
	if err != nil {panic(err)}

	users := []*User{}
	err = collection.Find(nil).All(&users)

	return users, err
}