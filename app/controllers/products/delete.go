package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//	@Tags		Product_Delete
//	@Success	200	{string}	string
//	@Router		/product/{id} [delete]
//	@Param		id	path	string	true	"id of product"
func ProductDelete(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
