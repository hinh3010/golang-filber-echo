package handler

import (
	"fmt"
	"lv1/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userName := claims["user"].(string)
	isAdmin := claims["admin"].(bool)

	message := fmt.Sprintf("hello %s is admin %v", userName, isAdmin)

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}
