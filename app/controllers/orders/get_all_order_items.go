package orders

import (
	"instashop/infra/config"
	"instashop/infra/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderGetAllResponse struct {
	Id       string `json:"id"`
	Product1 string `json:"product1"`
	Product2 string `json:"product2"`
	Product3 string `json:"product3"`
}

// @Tags		Order_GetAll
// @Success	200	{array}	OrderGetAllResponse	"success"
// @Router		/orders [get]
// @Param		Authorization	header	string	true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
func OrderGetAll(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		orderItems, err := ap.DbQueries.OrderItem_GetAll(ap.Ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could fetch product",
			})
		}

		return c.JSON(http.StatusOK, orderItems)
	}
}
