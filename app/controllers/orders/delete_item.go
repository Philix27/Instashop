package orders

import (
	"instashop/infra/config"
	"instashop/infra/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderItemDeleteResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

// @Tags		Order_Delete
// @Success	200	{string}	string
// @Router		/orders/{id} [delete]
// @Param		id	path	string	true	"id of order"
// @Param		Authorization	header		string			true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
func OrderDelete(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		order_id := c.Param("id")

		num, err := strconv.Atoi(order_id)

		if err != nil {

			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid param, must be a number",
			})

		}

		err = ap.DbQueries.OrderItem_DeleteById(ap.Ctx, int64(num))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not delete item",
			})
		}
		return c.JSON(http.StatusOK, OrderItemDeleteResponse{
			Id:      order_id,
			Message: "Deleted successfully",
		})
	}
}
