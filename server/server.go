
package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
)


func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted Access
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// CHECKOUT!! e.Group()

	// ROUTES
	e.GET("/", accessible)
	r.GET("", restricted)

	e.GET("/:username/accounts/:accountId/search/:socialNetwork/:searchTerm", SearchController)
	e.GET("/:username/posts", GetAllPosts)

	e.GET("/:username/accounts/:account_id/posts", FindAllAccountPosts)
	e.GET("/:username/accounts/:account_id/posts/unapproved", FindAccountUnapprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/Approved", FindAccountApprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/Disapproved", FindAccountDisapprovedPosts)

	e.GET("/:username/posts/unapproved", FindAllUnapprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/approved", FindAllApprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/disapproved", FindAllDisapprovedPosts)

	e.GET("/:username/find/:postId", FindPost)

	e.GET("/user/:username", GetUser)
	e.POST("/user", CreateUser)
	e.GET("/users", GetAllUsers)

	e.GET("/:username/accounts", GetAllAccountsByUsername)
	e.GET("/:username/accounts/:account_id", GetAccountById)
	e.POST("/:username/accounts/:accountId/twitter/creds", CreateAccountCreds)
	e.GET("/:username/accounts/:accountId/account-creds", GetAccountCreds)
	e.GET("/:username/accounts/find/:title", GetAccountByTitle)
	e.POST("/:username/accounts/create/:title", CreateAccount)

	e.POST("/post/upload/twitter/:postId/:tweetText", PostTweet)
	e.POST("/post/approve/:postId", ApprovePost)
	e.POST("/post/disapprove/:postId", DisapprovePost)
	e.POST("/:username/accounts/:account_id/post/delete/:postId", RemovePost)
	e.POST("/login", Login)

	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
