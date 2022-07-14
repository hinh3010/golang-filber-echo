package handler

import (
	"lv1/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloAdmin(c echo.Context) error {
	x := &models.X{
		Text: "message",
	}
	return c.JSON(http.StatusOK, x)
}
