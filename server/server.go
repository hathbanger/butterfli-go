package server

import (
	"fmt"
	"log"


	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
	// "github.com/user-base/models"
	"golang.org/x/net/websocket"
)

var (
	Msg       = websocket.Message
	ActiveClients = make(map[ClientConn]int) // map containing clients
)

type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}

func hello() websocket.Handler {

	var clientMessage string
	return websocket.Handler(func(ws *websocket.Conn) {
		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			// websocket.Message.Send(ws, &msg)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", msg)	
			client := ws.Request().RemoteAddr

			log.Println("Client connected:", client)

			sockCli := ClientConn{ws, client}
			ActiveClients[sockCli] = 0
			log.Println("Number of clients connected ...", len(ActiveClients))			
			clientMessage = sockCli.clientIP + " Said: " + msg
			for cs, _ := range ActiveClients {
				if err = Msg.Send(cs.websocket, clientMessage); err != nil {
					// we could not send the message to a peer
					log.Println("Could not send message to ", cs.clientIP, err.Error())
				}
			}		
		}
	})
}


func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	//e.Use(middleware.Static("/static"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	 AllowOrigins: []string{"*"},
	 AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Static("/admin", "static")


// ROUTES
	e.GET("/", accessible)
	r.GET("", restricted)
	e.GET("/user/:username", GetUser)
	e.GET("/ws", standard.WrapHandler(hello()))
	e.POST("/user", CreateUser)
	e.POST("/login", Login)
	e.POST("/message", CreateMessage)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}

