package server


import (
	//"fmt"
	//"time"
	//"labix.org/v2/mgo/bson"
	"github.com/butterfli-go/models"
	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"

	//"net/url"
	//"net/url"
	"net/http"
	//"fmt"
	"fmt"
	"net/url"
)

func SearchController(c echo.Context) error {
	socialNetwork := c.Param("socialNetwork")
	searchTerm := c.Param("searchTerm")
	username := c.Param("username")

	searchTermStruct := models.FindSearchTerm(searchTerm)
	fmt.Print("booooom \n")
	fmt.Print(searchTermStruct.Text)

	results := Search(username, socialNetwork, searchTerm)

	for _, tweet := range results.Statuses {
		if len(tweet.Entities.Media) > 0 {
			fmt.Print("\n")
			imgurl :=  tweet.Entities.Media[0].Media_url
			fmt.Print(imgurl)
			post := models.NewPost(username, searchTerm, imgurl)
			err := post.Save()
			if err != nil {
				fmt.Print(" - - > Failure on this one.. Probably a duplicate.")
			}
		}
	}
	return c.JSON(http.StatusOK, results)
}


func Search(username string, socialNetwork string, searchTerm string) anaconda.SearchResponse {
	switch socialNetwork {
	case "twitter":
		return SearchTwitter(username, searchTerm)
	default:
		panic("unrecognized escape character")
	}
}

func SearchTwitter(username string, searchTerm string) anaconda.SearchResponse {
	v := url.Values{}
	v.Set("count", "30")

	updatedSearch := searchTerm + " filter:twimg"
	anaconda.SetConsumerKey("32lvF7IqHpkZwJDDey4f160fT")
	anaconda.SetConsumerSecret("nXcawfuDxew7gAdYyi2J3CDQQWEIWIsCPHSNO8kOlqEuSDMDGN")
	api := anaconda.NewTwitterApi("28226407-zLDSNIDqXDEtK9YnDqBH1agA45BXjFkOkA03aZYsf", "SHz6fDl31gLYGXYQdiSbxrDiZFZTH6ewIJN4Kp2DDcjiI")
	//api.EnableThrottling(10*time.Second, 5)
	search_result, err := api.GetSearch(updatedSearch, v)
	fmt.Print("SearchTwitter success")
	fmt.Print("\n")

	if err != nil {
		panic(err)
	}
	return search_result
}
