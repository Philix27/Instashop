package middleware

import "github.com/labstack/echo/v4"

// func IsAdmin(username, password string, c echo.Context) (bool, error) {
// 	if username == "joe" && password == "secret" {
// 		return true, nil
// 	}
// 	return false, nil
// }

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
