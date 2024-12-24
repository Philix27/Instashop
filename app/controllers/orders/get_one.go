package orders

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags		Order_GetOne
// @Success	200	{string}	string
// @Router		/orders/{id} [get]
// @Param		id	path	string	true	"id of order"
func OrderGetOne(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
