package server


import (
	"github.com/butterfli-go/models"
	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"

	"net/http"
	"fmt"
	"log"
	"net/url"
	"encoding/base64"
	"strconv"
	"io/ioutil"
	//"gopkg.in/mgo.v2/bson"
	//"labix.org/v2/mgo/bson"
	//"gopkg.in/mgo.v2/bson"
)

func SearchController(c echo.Context) error {

	socialNetwork := c.Param("socialNetwork")
	searchTerm := c.Param("searchTerm")
	accountId := c.Param("accountId")
	username := c.Param("username")

	//searchTermStruct := models.FindSearchTerm(searchTerm)
	fmt.Print("booooom \n")
	//fmt.Print(searchTermStruct.Text)

	results := Search(username, accountId, socialNetwork, searchTerm)

	for _, tweet := range results.Statuses {
		if len(tweet.Entities.Media) > 0 {
			fmt.Print("\n")
			imgurl :=  tweet.Entities.Media[0].Media_url
			fmt.Print(imgurl)
			post := models.NewPost(username, accountId, searchTerm, imgurl)
			err := post.Save()
			if err != nil {
				fmt.Print(" - - > Failure on this one.. Probably a duplicate.")
			}
		}
	}

	return c.JSON(http.StatusOK, results)
}

func PostTweet(c echo.Context) error {
	accountId := c.Param("account_id")
	postId := c.Param("postId")
	tweetText := c.Param("tweetText")
	results := PostMediaToTwitter(accountId, postId, tweetText)

	return c.JSON(http.StatusOK, results)
}


func Search(username string, accountId string, socialNetwork string, searchTerm string) anaconda.SearchResponse {
	switch socialNetwork {
	case "twitter":
		return SearchTwitter(username, accountId, searchTerm)
	default:
		panic("unrecognized escape character")
	}
}

func SearchTwitter(username string, accountId string, searchTerm string) anaconda.SearchResponse {
	fmt.Print("being seracdhTwitter")
	v := url.Values{}
	v.Set("count", "30")

	updatedSearch := searchTerm + " filter:twimg"
	fmt.Print(updatedSearch)
	accountCreds, err := models.FindAccountCredsByAccountId(accountId)
	fmt.Print("pastfindingacct")
	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)
	//api.EnableThrottling(10*time.Second, 5)
	search_result, err := api.GetSearch(updatedSearch, v)

	if err != nil {
		panic(err)
	}
	return search_result
}

func PostToTwitter(accountId string, text string) anaconda.Tweet {
	v := url.Values{}
	v.Set("count", "30")

	accountCreds, err := models.FindAccountCredsByAccountId(accountId)

	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)
	tweet, err := api.PostTweet(text, v)

	if err != nil {
		panic(err)
	}
	return tweet
}

func PostMediaToTwitter(accountId string, postId string, text string) anaconda.Tweet {

	post, err := models.FindPostById(postId)
	//data :=
	res, err := http.Get(post.Imgurl)
	if err != nil {
		fmt.Print(err)
	}

	//defer data.Body.Close()

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)


	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	fmt.Print("here we go w the account Creds: \n")
	fmt.Print(accountId)

	accountCreds, err := models.FindAccountCredsByAccountId(accountId)

	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)

	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	res.Body.Close()

	result, err := api.PostTweet(text, v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	return result
}
