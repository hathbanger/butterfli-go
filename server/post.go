package server

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/butterfli-go/models"
	"fmt"
)



func GetAllPosts(c echo.Context) error {
	username := c.Param("username")
	posts, err := models.GetAllPosts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAllAccountPosts(c echo.Context) error {
	accountId := c.Param("account_id")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(accountId)
	posts, err := models.GetAllAccountPosts(accountId)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAccountUnapprovedPosts(c echo.Context) error {
	accountId := c.Param("account_id")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(accountId)
	posts, err := models.GetAccountUnapprovedPosts(accountId)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAccountApprovedPosts(c echo.Context) error {
	accountId := c.Param("account_id")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(accountId)
	posts, err := models.GetAccountApprovedPosts(accountId)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAccountDisapprovedPosts(c echo.Context) error {
	accountId := c.Param("account_id")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(accountId)
	posts, err := models.GetAccountDisapprovedPosts(accountId)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAllUnapprovedPosts(c echo.Context) error {
	username := c.Param("username")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(username)
	posts, err := models.GetAllUnapprovedPosts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func FindAllApprovedPosts(c echo.Context) error {
	username := c.Param("username")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(username)
	posts, err := models.GetAllApprovedPosts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}


func FindAllDisapprovedPosts(c echo.Context) error {
	username := c.Param("username")
	fmt.Print("getting all posts for an acct!")
	fmt.Print(username)
	posts, err := models.GetAllDisapprovedPosts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, posts)
}


func FindPost(c echo.Context) error {
	postId := c.Param("postId")
	post, err := models.FindPostById(postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found!")
	}

	return c.JSON(http.StatusOK, post)
}

func ApprovePost(c echo.Context) error {
	postId := c.Param("postId")
	fmt.Print(postId)
	fmt.Print("\n that was postid")
	err := models.ApprovePostById(postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not approved")
	}
	return c.JSON(http.StatusOK, "approved!")
}

func DisapprovePost(c echo.Context) error {
	postId := c.Param("postId")
	err := models.DisapprovePostById(postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "failure")
	}
	return c.JSON(http.StatusOK, "disapproved!")
}


func RemovePost(c echo.Context) error {
	postId := c.Param("postId")
	err := models.DeletePost(postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not able to remove the post..")
	}

	return c.JSON(http.StatusOK, "worked!!")
}
