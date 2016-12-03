package server

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/butterfli-go/models"
)



func GetAllPosts(c echo.Context) error {
	username := c.Param("username")
	posts, err := models.GetAllPosts(username)
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
	post_id := c.Param("post_id")
	err := models.DeletePost(post_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not able to remove the post..")
	}

	return c.JSON(http.StatusOK, "worked!!")
}
