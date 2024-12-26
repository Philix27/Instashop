package orders

import (
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
func OrderGetAll(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
