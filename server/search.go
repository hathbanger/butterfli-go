package server


import (
	"github.com/butterfli-go/models"
	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"
	"net/http"
	"fmt"

	"net/url"
	//"encoding/base64"
	"strconv"
	//"io/ioutil"
	"strings"
	"time"
	//"github.com/smartystreets/goconvey/web/server/api"
)


func SearchController(c echo.Context) error {
	socialNetwork := c.Param("socialNetwork")
	searchTermString := c.Param("searchTerm")
	accountId := c.Param("accountId")
	username := c.Param("username")
	searchTerm, err := models.FindSearchTerm(accountId, searchTermString)
	if err != nil {
		fmt.Print("WOAH! NEW TERM")
		searchTerm = models.NewSearchTerm(accountId, searchTermString)
		searchTerm.Save()
	}
	results := Search(username, accountId, socialNetwork, *searchTerm)
	postedResults := CreatePostFromResults(username, accountId, searchTerm, socialNetwork, results)

	return c.JSON(http.StatusOK, postedResults)
}

func SearchAndFavorite(c echo.Context) error {

	searchTermString := c.Param("searchTerm")
	accountId := c.Param("accountId")

	api := AuthTwitter(accountId)
	favoriteTerm, err := models.FindFavoriteTerm(accountId, searchTermString)
	if err != nil {
		favoriteTerm = models.NewFavoriteTerm(accountId, searchTermString)
		favoriteTerm.Save()
	}
	//results := Search(username, accountId, socialNetwork, *favoriteTerm)
	v := url.Values{}
	s := strconv.FormatInt(favoriteTerm.SinceTweetId, 10)
	v.Set("since_id", s)
	v.Add("count", "30")
	updatedSearch := favoriteTerm.Text
	search_result, err := api.GetSearch(updatedSearch, v)
	if err != nil {panic(err)}


	var succeses = 0
	var failures = 0
	for _, tweet := range search_result.Statuses {
		res, err := api.Favorite(tweet.Id)
		if res.Id != 0  {
			succeses = succeses + 1
			fmt.Print(" Success!")
		}
		if err != nil {
			failures = failures + 1
			fmt.Print("error!")
			fmt.Print(err)
		}
	}
	fmt.Print(succeses)

	return c.JSON(http.StatusOK, fmt.Sprintf("AccountId %s just favorited %v new tweets, and failed %v times", accountId, succeses, failures))
}


func Search(username string, accountId string, socialNetwork string, searchTerm models.SearchTerm) anaconda.SearchResponse {
	switch socialNetwork {
	case "twitter":
		return SearchTwitter(username, accountId, searchTerm)
	default:
		panic("unrecognized escape character")
	}
}

func AuthTwitter(accountId string) *anaconda.TwitterApi {
	accountCreds, err := models.FindAccountCredsByAccountId(accountId)
	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)

	if err != nil {panic(err)}

	return api
}

func SearchTwitter(username string, accountId string, searchTerm models.SearchTerm) anaconda.SearchResponse {
	v := url.Values{}
	s := strconv.FormatInt(searchTerm.SinceTweetId, 10)
	v.Set("since_id", s)
	v.Add("count", "30")
	updatedSearch := searchTerm.Text + " filter:twimg"
	api := AuthTwitter(accountId)
	search_result, err := api.GetSearch(updatedSearch, v)
	if err != nil {panic(err)}

	fmt.Print("search_result:")
	fmt.Print(search_result)

	return search_result
}


//func FavoriteTwitter(username string, accountId string, searchTerm models.SearchTerm) anaconda.SearchResponse {
//	v := url.Values{}
//	s := strconv.FormatInt(searchTerm.SinceTweetId, 10)
//	v.Set("since_id", s)
//	v.Add("count", "30")
//	updatedSearch := searchTerm.Text + " filter:twimg"
//	api := AuthTwitter(accountId)
//	search_result, err := api.GetSearch(updatedSearch, v)
//	if err != nil {panic(err)}
//
//	fmt.Print("search_result:")
//	fmt.Print(search_result)
//
//	return search_result
//}


// func FavoriteTweets(c echo.Context) error {

// }



func GetAllSearchTerms(c echo.Context) error {
	accountId := c.Param("accountId")
	searchTerms := models.FindAllSearchTerms(accountId)
	return c.JSON(http.StatusOK, searchTerms)
}

func BotnetFavoriteTweet(c echo.Context) error {
	tweetId := c.Param("tweetId")
	accountsArray := c.Param("accountsArray")
	tweetId64, err := strconv.ParseInt(tweetId, 10, 64)
	if err != nil {
		panic(err)
	}
	accountsSlice := strings.Split(accountsArray, "+")
	for _, accountId := range accountsSlice {
		api := AuthTwitter(accountId)
		api.Favorite(tweetId64)
		api.EnableThrottling(10*time.Second, 5)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just liked tweetId %s with %v accounts", tweetId, len(accountsSlice)))
}


func BotnetFollowAccountId(c echo.Context) error {
	followAccountId := c.Param("accountId")
	accountsArray := c.Param("accountsArray")
	followAccountId64, err := strconv.ParseInt(followAccountId, 10, 64)
	if err != nil {
		panic(err)
	}
	accountsSlice := strings.Split(accountsArray, "+")
	for _, accountId := range accountsSlice {
		api := AuthTwitter(accountId)
		api.FollowUserId(followAccountId64, nil)
		api.EnableThrottling(10*time.Second, 5)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just followed accountId %s with %v accounts", followAccountId, len(accountsSlice)))
}

func BotnetFollowAccountName(c echo.Context) error {
	followAccountName := c.Param("accountName")
	accountsArray := c.Param("accountsArray")

	accountsSlice := strings.Split(accountsArray, "+")
	for _, accountId := range accountsSlice {
		api := AuthTwitter(accountId)
		api.FollowUser(followAccountName)
		api.EnableThrottling(10*time.Second, 5)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just followed %s with %v accounts", followAccountName, len(accountsSlice)))
}