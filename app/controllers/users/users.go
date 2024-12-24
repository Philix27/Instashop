package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserUpdate(c echo.Context) error {
	return c.String(http.StatusOK, "updateUser")
}

func UserDelete(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUser")
}

func UserCreate(c echo.Context) error {
	return c.String(http.StatusOK, "createUser")
}

// e.GET("/users/:id", getUser)
func UserGet(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/users/:id", getUser)
func UserGetAll(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
