package orders

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderCancelInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type OrderCancelResponse struct {
	UserId      string `json:"UserId"`
	AccessToken string `json:"password"`
}

//	@Tags		Order_cancel
//	@Router		/orders/cancel/{id} [patch]
//	@Param		id		path		string				true	"id of order"
//	@Param		request	body		OrderCancelInput	true	"Order_cancel"
//	@Success	200		{object}	OrderCancelResponse	"success"
//	@Accept		json
//	@Produce	json
//	@Failure	400				{object}	types.ErrMsg	"error"
//	@Param		Authorization	header		string			true	"MyHeader must be set for valid response"
func OrderCancel(c echo.Context) error {
	order_id := c.Param("id")
	return c.String(http.StatusOK, order_id)
}
