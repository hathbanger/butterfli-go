package server

import (
	"fmt"
	"github.com/user-boiler/models"
	"github.com/labstack/echo"
)

func CreateMessage(c echo.Context) error {
	username := c.FormValue("username")
	message := c.FormValue("message")

	mes := models.NewMessage(username, message)
	err := mes.Save()
	if err != nil {
		fmt.Println(err)
	}

	return err
}
