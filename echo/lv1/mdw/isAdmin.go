package mdw

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAdminMdw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		isAdmin := claims["admin"].(bool)

		if isAdmin {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
