
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
	// User Routes
	e.GET("/:username", GetUser)
	e.POST("/user", CreateUser)
	e.GET("/users", GetAllUsers)
	e.POST("/login", Login)

	e.GET("/:username/posts", FindPostsByUsername)
	e.GET("/:username/posts/:postId", FindPost)
	e.GET("/:username/posts/unapproved", FindAllUnapprovedPosts)
	e.GET("/:username/posts/approved", FindAllApprovedPosts)
	e.GET("/:username/posts/disapproved", FindAllDisapprovedPosts)

	e.GET("/:username/accounts", GetAllAccountsByUsername)
	e.GET("/:username/accounts/:account_id", GetAccountById)
	e.GET("/:username/accounts/:account_id/search-terms", GetAllSearchTerms)
	e.GET("/:username/accounts/find/:title", GetAccountByTitle)
	e.GET("/:username/accounts/:accountId/account-creds", GetAccountCreds)
	e.POST("/:username/accounts/:account_id/delete", RemoveAccount)

	e.GET("/:username/accounts/:account_id/posts", FindAllAccountPosts)
	e.GET("/:username/accounts/:account_id/posts/unapproved", FindAccountUnapprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/approved", FindAccountApprovedPosts)
	e.GET("/:username/accounts/:account_id/posts/disapproved", FindAccountDisapprovedPosts)
	e.POST("/post/approve/:postId", ApprovePost)
	e.POST("/post/disapprove/:postId", DisapprovePost)


	e.POST("/:username/accounts/:accountId/twitter/creds", CreateAccountCreds)
	e.POST("/:username/accounts/:account_id/post/:postId/upload/twitter/:tweetText", PostTweet)
	e.GET("/:username/accounts/:accountId/search/:socialNetwork/:searchTerm", SearchController)
	e.POST("/:username/accounts/:account_id/post/delete/:postId", RemovePost)


	// NOT TESTED
	e.POST("/:username/accounts/create/:title", CreateAccount)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
