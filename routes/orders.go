package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OrderUpdate(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
func OrderCancel(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
func OrderDelete(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
func OrderCreate(c echo.Context) error { return nil }
func OrderGetAll(c echo.Context) error { return nil }
func OrderGetOne(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
