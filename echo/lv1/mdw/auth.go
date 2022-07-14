package mdw

import "github.com/labstack/echo/v4"

func BasicAuth(user string, pass string, c echo.Context) (bool, error) {
	if user == "admin" && pass == "adu" {
		c.Set("user", user)
		c.Set("admin", true)
		return true, nil
	}

	if user == "adu" && pass == "adu" {
		c.Set("user", user)
		c.Set("admin", false)
		return true, nil
	}

	return false, nil
}
