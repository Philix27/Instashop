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

type OrderUpdateInput struct {
	Status db.Orderstatus
}
type OrderUpdateResponse struct {
	Message string `json:"message"`
}

// @Tags		Order_Update
// @Success	200	{string}	string
// @Router		/orders/{id} [put]
// @Param		id	path	string	true	"id of order"
// @Param		Authorization	header		string			true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
func OrderUpdate(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")
		num, err := strconv.Atoi(product_id)

		dto := new(OrderUpdateInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		err = ap.DbQueries.Orders_UpdateStatus(ap.Ctx, db.Orders_UpdateStatusParams{
			ID:          int64(num),
			OrderStatus: dto.Status,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not create product",
			})
		}

		// return c.String(http.StatusOK, product_id)
		return c.JSON(http.StatusOK, OrderUpdateResponse{
			Message: "Updated",
		})
	}
}
