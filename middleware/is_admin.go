package middleware

import "github.com/labstack/echo/v4"

func IsAdmin(username, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}
