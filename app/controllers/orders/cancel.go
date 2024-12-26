package orders

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderCancelResponse struct {
	OrderId string `json:"orderId"`
}

// @Tags		Order_Cancel
// @Router		/orders/cancel/{id} [patch]
// @Param		id	path		string				true	"id of order"
// @Success	200	{object}	OrderCancelResponse	"success"
// @Accept		json
// @Produce	json
// @Failure	400				{object}	types.ErrMsg	"error"
// @Param		Authorization	header		string			true	"Header must be set for valid response"
func OrderCancel(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
