package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProductUpdate(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
func ProductDelete(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
func ProductCreate(c echo.Context) error { return nil }
func ProductGetAll(c echo.Context) error { return nil }
func ProductGetOne(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
