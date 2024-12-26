package middleware

import (
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
