package handler

import (
	"log"
	"lv1/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	user := c.Get("user").(string)
	admin := c.Get("admin").(bool)

	// log.Printf("%s", user)
	// log.Printf("%v", admin)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))

	if err != nil {
		log.Printf("sign token error: %v", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginRes{
		Token: "Bearer " + t,
	})
}
