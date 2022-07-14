package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Age  int
}

var listUsers = []User{
	{
		Name: "adu",
		Age:  20,
	},
	{
		Name: "adu2",
		Age:  22,
	},
	{
		Name: "adu3",
		Age:  23,
	},
}

func CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "api get create")
}

func GetUsers(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(c.Response())
	for _, user := range listUsers {
		if err := encoder.Encode(user); err != nil {
			return err
		}
		c.Response().Flush()
		time.Sleep(time.Second * 1)
	}
	return nil
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "api get user by "+id)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "api update user by "+id)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "api delete user by "+id)
}
