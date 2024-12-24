package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


// @Tags		Product_Create
//	@Success	200	{string}	string
//	@Router		/products [post]
func ProductCreate(c echo.Context) error { return nil }


// @Tags		Product_GetAll
//	@Success	200	{string}	string
//	@Router		/products [get]
func ProductGetAll(c echo.Context) error { return nil }

// @Tags		Product_GetOne
//	@Success	200	{string}	string
//	@Router		/products/{id} [get]
//	@Param		id	path	string	true	"id of product"
func ProductGetOne(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
