package orders

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags		Order_Delete
// @Success	200	{string}	string
// @Router		/orders/{id} [delete]
// @Param		id	path	string	true	"id of order"
// @Param		Authorization	header		string			true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
func OrderDelete(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
