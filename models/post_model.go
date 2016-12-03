package models

import (
	//"encoding/json"

	//"labix.org/v2/mgo"
	"gopkg.in/mgo.v2/bson"
	//"log"
	//"sync"
	"fmt"
	"time"
	"github.com/butterfli-go/store"
)


type Post struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Imgurl		string           `json:"imgurl",bson:"imgurl,omitempty"`
	SearchTerm	string           `json:"searchterm",bson:"searchterm,omitempty"`
	Approved	bool           `json:"approved",bson:"approved,omitempty"`
	Rated	bool           `json:"rated",bson:"rated,omitempty"`
}

func NewPost(username string, searchTerm string, imgUrl string) *Post {
	p := new(Post)
	p.Id = bson.NewObjectId()
	p.Timestamp = time.Now()
	p.Username = username
	p.SearchTerm = searchTerm
	p.Imgurl = imgUrl
	p.Approved = false
	p.Rated = false

	return p
}

func (p *Post) Save() error {
	fmt.Print("saving! from the top ")
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToPostsCollection(session, "posts")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&Post{
		Id: p.Id,
		Timestamp: p.Timestamp,
		Username: p.Username,
		SearchTerm: p.SearchTerm,
		Imgurl: p.Imgurl,
		Approved: p.Approved,
		Rated: p.Rated})

	//fmt.Print(Post)

	if err != nil {
		return err
	}

	return nil
}


func FindPostById(postId string) (*Post, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToPostsCollection(session, "posts")
	if err != nil {
		//panic(err)
		return &Post{}, err
	}

	post := Post{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(postId)}).One(&post)
	if err != nil {
		panic(err)
		//return &post, err
	}

	return &post, err
}

func GetAllPosts(username string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")

	posts := []*Post{}

	err = collection.Find(bson.M{}).All(&posts)

	return posts, err
}

func DeletePost(id string) error {
	session, err := store.ConnectToDb()

	collection := session.DB("test").C("posts")
	fmt.Println("id:")
	fmt.Println(id)

	err = collection.Remove(bson.M{"id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Println("fack")
		fmt.Println(err)
	}
	return nil
}


func ApprovePostById(postId string) error {
	session, err := store.ConnectToDb()

	collection := session.DB("test").C("posts")
	fmt.Println("id:")
	fmt.Println(postId)

	post, err := FindPostById(postId)

	fmt.Print(post)
	fmt.Print(" and then ")
	fmt.Print(postId)

	colQuerier := bson.M{"id": bson.ObjectIdHex(postId)}
	change := bson.M{"$set": bson.M{ "approved": true, "rated": true }}
	err = collection.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
	return nil
}

func DisapprovePostById(postId string) error {
	session, err := store.ConnectToDb()

	collection := session.DB("test").C("posts")
	fmt.Println("id:")
	fmt.Println(postId)

	post, err := FindPostById(postId)

	fmt.Print(post)
	fmt.Print(" and then ")
	fmt.Print(postId)

	colQuerier := bson.M{"id": bson.ObjectIdHex(postId)}
	change := bson.M{"$set": bson.M{ "approved": false, "rated": true }}
	err = collection.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
	return nil
}