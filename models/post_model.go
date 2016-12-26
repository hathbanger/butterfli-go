package models

import (
	"labix.org/v2/mgo/bson"
	"fmt"
	"time"
	"github.com/butterfli-go/store"
)


type Post struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Account	string           `json:"account",bson:"account,omitempty"`
	Imgurl		string           `json:"imgurl",bson:"imgurl,omitempty"`
	SearchTerm	string           `json:"searchterm",bson:"searchterm,omitempty"`
	Approved	bool           `json:"approved",bson:"approved,omitempty"`
	Rated	bool           `json:"rated",bson:"rated,omitempty"`
}

func NewPost(username string, account string, searchTerm string, imgUrl string) *Post {
	p := new(Post)
	p.Id = bson.NewObjectId()
	p.Timestamp = time.Now()
	p.Username = username
	p.Account = account
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

	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})
	if err != nil {
		panic(err)
	}
	post := &Post{
		Id: p.Id,
		Timestamp: p.Timestamp,
		Username: p.Username,
		Account: p.Account,
		SearchTerm: p.SearchTerm,
		Imgurl: p.Imgurl,
		Approved: p.Approved,
		Rated: p.Rated}


	err = collection.Insert(post)


	if err != nil {
		return err
	}

	return nil
}


func FindPostById(post_id string) (*Post, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	fmt.Print("\nfind post by id commence: ")
	fmt.Print(post_id)
	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})
	//collection, err := store.ConnectToPostsCollection(session, "posts")
	if err != nil {
		//panic(err)
		fmt.Print("\nfind post byid error\n")
		return &Post{}, err
	}

	post := Post{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(post_id)}).One(&post)
	if err != nil {

		fmt.Print("\n not found!\n\n\n")
		//panic(err)

		return &post, err
	}

	return &post, err
}

func GetAllPosts(username string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})

	posts := []*Post{}

	err = collection.Find(bson.M{"username": username}).All(&posts)

	return posts, err
}



func GetAccountUnapprovedPosts(accountId string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})

	posts := []*Post{}

	err = collection.Find(bson.M{"account": accountId, "rated": false}).All(&posts)

	return posts, err
}


func GetAccountApprovedPosts(accountId string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})

	posts := []*Post{}

	err = collection.Find(bson.M{"account": accountId, "approved": true}).All(&posts)

	return posts, err
}

func GetAccountDisapprovedPosts(accountId string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})

	posts := []*Post{}

	err = collection.Find(bson.M{"account": accountId, "approved": false}).All(&posts)

	return posts, err
}



func GetAllUnapprovedPosts(username string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")

	posts := []*Post{}

	err = collection.Find(bson.M{"username": username, "rated": false}).All(&posts)

	return posts, err
}

func GetAllApprovedPosts(username string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")

	posts := []*Post{}

	err = collection.Find(bson.M{"username": username, "approved": true}).All(&posts)

	return posts, err
}

func GetAllDisapprovedPosts(username string) ([]*Post, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")

	posts := []*Post{}

	err = collection.Find(bson.M{"username": username, "approved": false}).All(&posts)

	return posts, err
}

func GetAllAccountPosts(accountId string) ([]*Post, error){
	fmt.Print("\n\nthis is the acct ID " + accountId)

	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("posts")


	posts := []*Post{}

	//account, err := FindAccountById(accountId)

	err = collection.Find(bson.M{"account": accountId}).All(&posts)

	//fmt.Print(account.Posts)

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
		fmt.Print(" issues ")
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