package products

import (
	"instashop/infra/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductGetAllResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

// @Tags		Product_GetAll
// @Success	200	{string}	string
// @Router		/products [get]
// @Param		Authorization	header	string	true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
// @Success	200	{array}	ProductGetAllResponse	"success"
func ProductGetAll(appState config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")
		return c.String(http.StatusOK, product_id)
	}
}
