package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func updateUser(c echo.Context) error { return nil }
func deleteUser(c echo.Context) error { return nil }
func createUser(c echo.Context) error { return nil }

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
