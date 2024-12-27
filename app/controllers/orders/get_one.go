package orders

import (
	"instashop/infra/config"
	"instashop/infra/types"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type OrderGetOneResponse struct {
	ID        int64 `json:"id"`
	Quantity  int32 `json:"quantity"`
	ProductID int32 `json:"productId"`
}

// @Tags		Order_GetOne
// @Success	200	{array}	OrderGetOneResponse	"success"
// @Router		/orders/{id} [get]
// @Param		id	path	string	true	"id of order"
//
//	@Param		Authorization	header		string			true	"Header must be set for valid response"
//	@Accept		json
//	@Produce	json
func OrderGetOne(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		order_id := c.Param("id")

		num, err := strconv.Atoi(order_id)

		if err != nil {

			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid param, must be a number",
			})

		}

		orderItems, err := ap.DbQueries.OrderItem_GetAllOrderId(ap.Ctx, pgtype.Int4{
			Int32: int32(num),
		})

		return c.JSON(http.StatusOK, orderItems)
	}
}
