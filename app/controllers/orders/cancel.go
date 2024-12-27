package orders

import (
	"instashop/db"
	"instashop/infra/config"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderUpdateStatusInput struct {
	Status string `json:"status" enums:"CANCEL"`
}

type OrderCancelResponse struct {
	Message string `json:"message"`
}

// @Tags		Order_UpdateStatus
// @Router		/orders/cancel/{id} [patch]
// @Param		id		path		string					true	"id of order"
// @Success	200		{object}	OrderCancelResponse		"success"
// @Param		request	body		OrderUpdateStatusInput	true	"Order_UpdateStatus"
// @Accept		json
// @Produce	json
// @Failure	400				{object}	types.ErrMsg	"error"
// @Param		Authorization	header		string			true	"Header must be set for valid response"
func OrderCancel(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		order_id := c.Param("id")
		num, err := strconv.Atoi(order_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid param, must be a number",
			})
		}

		dto := new(OrderUpdateStatusInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		anOrder, err := ap.DbQueries.Orders_GetById(ap.Ctx, int64(num))

		if anOrder.OrderStatus == db.OrderstatusCOMPLETED {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "Order has already been completed",
			})
		}
		if anOrder.OrderStatus == db.OrderstatusCANCEL {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "Order has already been canceled",
			})
		}

		if dto.Status != "CANCEL" {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "A valid state needed",
			})
		}

		err = ap.DbQueries.Orders_UpdateStatus(ap.Ctx, db.Orders_UpdateStatusParams{
			ID:          int64(num),
			OrderStatus: db.OrderstatusCANCEL,
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "Could not update order " + err.Error(),
			})
		}
		return c.JSON(http.StatusOK, OrderCancelResponse{
			Message: "updated successfully",
		})

	}
}
