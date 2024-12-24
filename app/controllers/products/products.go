package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


// ProductUpdate  Product Update
//	@Description	Product update
//	@Tags			ProductUpdate
//	@Success		200	{string} string
//	@Router			/product/{id} [patch]
//	@Param			id path string true "id of product"
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
