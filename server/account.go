package server

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/butterfli-go/models"
	"fmt"

)

func GetAllAccountsByUsername(c echo.Context) error {
	username := c.Param("username")
	accounts, err := models.GetAllAccounts(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, accounts)
}

func CreateAccount(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")

	fmt.Print("creating an account for you")
	fmt.Print("with name: ")
	fmt.Print(title)

	account := models.NewAccount(username, title)
	err := account.Save()
	if err != nil {
		fmt.Print("account creation failure!")
		return c.JSON(http.StatusForbidden, "We're sorry! There's already an account with that name..")
	}
	fmt.Print("account creation success!")

	return c.JSON(http.StatusOK, account)
}




func CreateAccountCreds(c echo.Context) error {
	username := c.Param("username")
	accountId := c.Param("accountId")
	consumerKey := c.FormValue("consumerKey")
	consumerSecret := c.FormValue("consumerSecret")
	accessToken := c.FormValue("accessToken")
	accessTokenSecret := c.FormValue("accessTokenSecret")

	fmt.Print("creating account creds for you")

	account := models.NewAccountCreds(username, accountId, consumerKey, consumerSecret, accessToken, accessTokenSecret)

	err := account.Save()
	if err != nil {
		fmt.Print("account creation failure!")
		return c.JSON(http.StatusForbidden, "We're sorry! There's already an account with that name..")
	}
	fmt.Print("account creation success!")

	return c.JSON(http.StatusOK, account)
}


func GetAccountCreds(c echo.Context) error {
	accountId := c.Param("accountId")

	fmt.Print("getting the account creds for you")
	account, err := models.FindAccountCredsByAccountId(accountId)

	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! we couldn't find it....")
	}
	fmt.Print("found account creds!")

	return c.JSON(http.StatusOK, account)
}


func GetAccountByTitle(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")

	account, err := models.FindAccount(username, title)
	if err != nil {
		panic(err)
	}

	if account.Id != "" /*&& user.Username != "" */ {
		return c.JSON(http.StatusOK, account)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}


func GetAccountById(c echo.Context) error {
	account_id := c.Param("account_id")

	account, err := models.FindAccountById(account_id)
	if err != nil {
		panic(err)
	}

	if account.Id != "" /*&& user.Username != "" */ {
		return c.JSON(http.StatusOK, account)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}

