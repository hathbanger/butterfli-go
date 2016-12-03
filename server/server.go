
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
	e.GET("/:username/search/:socialNetwork/:searchTerm", SearchController)
	e.GET("/:username/posts", GetAllPosts)
	e.GET("/:username/find/:postId", FindPost)
	e.GET("/user/:username", GetUser)
	e.POST("/post/approve/:postId", ApprovePost)
	e.POST("/post/disapprove/:postId", DisapprovePost)
	e.POST("/post/delete/:post_id", RemovePost)
	e.POST("/user", CreateUser)
	e.GET("/users", GetAllUsers)
	e.POST("/login", Login)

	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
