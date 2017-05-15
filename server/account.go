package server

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/butterfli-go/models"
)

func GetAllAccountsByUsername(c echo.Context) error {
	username := c.Param("username")
	accounts, err := models.GetAllAccounts(username)
	if err != nil {panic(err)}

	return c.JSON(http.StatusOK, accounts)
}

func CreateAccount(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")
	account := models.NewAccount(username, title)
	err := account.Save()
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There's already an account with that name..")
	}
	return c.JSON(http.StatusOK, account)
}


func RemoveAccount(c echo.Context) error {
	accountId := c.Param("account_id")
	err := models.DeleteAccount(accountId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not able to remove the account..")
	}

	return c.JSON(http.StatusOK, "worked!!")
}


func CreateAccountCreds(c echo.Context) error {
	username := c.Param("username")
	accountId := c.Param("accountId")
	consumerKey := c.FormValue("consumerKey")
	consumerSecret := c.FormValue("consumerSecret")
	accessToken := c.FormValue("accessToken")
	accessTokenSecret := c.FormValue("accessTokenSecret")
	account := models.NewAccountCreds(username, accountId, consumerKey, consumerSecret, accessToken, accessTokenSecret)
	err := account.Save()
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There's already an account with that name..")
	}

	return c.JSON(http.StatusOK, account)
}


func GetAccountCreds(c echo.Context) error {
	accountId := c.Param("accountId")
	account, err := models.FindAccountCredsByAccountId(accountId)
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! we couldn't find it....")
	}
	return c.JSON(http.StatusOK, account)
}


func GetAccountByTitle(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")
	account, err := models.FindAccount(username, title)
	if err != nil {
		panic(err)
	}
	if account.Id != "" {
		return c.JSON(http.StatusOK, account)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}


func GetAccountById(c echo.Context) error {
	account_id := c.Param("account_id")
	account, err := models.FindAccountById(account_id)
	if err != nil {panic(err)}
	if account.Id != "" {
		return c.JSON(http.StatusOK, account)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}

