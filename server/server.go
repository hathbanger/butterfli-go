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
	 AllowOrigins: []string{"*"},
	 AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))



// ROUTES
	e.GET("/", accessible)
	r.GET("", restricted)
	e.GET("/user/:username", GetUser)
	e.POST("/user", CreateUser)
	e.GET("/users", GetAllUsers)
	e.POST("/login", Login)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}

