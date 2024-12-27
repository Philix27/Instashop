package products

import (
	"instashop/infra/config"
	"instashop/infra/types"
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
func ProductGetAll(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {

		products, err := ap.DbQueries.Products_GetAll(ap.Ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not delete product",
			})
		}

		return c.JSON(http.StatusOK, products)
	}
}
