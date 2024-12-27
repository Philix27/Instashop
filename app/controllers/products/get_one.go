package products

import (
	"instashop/infra/config"
	"instashop/infra/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductGetOneResponse struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
}

// @Tags		Product_GetOne
// @Success	200	{string}	string
// @Router		/products/{id} [get]
// @Param		id				path	string	true	"id of product"
// @Param		Authorization	header	string	true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
// @Success	200	{object}	ProductGetOneResponse	"success"
func ProductGetOne(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")

		num, err := strconv.Atoi(product_id)

		if err != nil {

			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid param, must be a number",
			})

		}
		product, err := ap.DbQueries.Products_GetOne(ap.Ctx, int64(num))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not fetch products",
			})
		}

		return c.JSON(http.StatusOK, ProductGetOneResponse{
			Id:          product.ID,
			Title:       product.Title,
			Description: product.Description.String,
			Price:       product.Price,
		})
	}
}
