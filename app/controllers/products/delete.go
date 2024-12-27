package products

import (
	"instashop/infra/config"
	"instashop/infra/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductDeleteResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

// @Tags		Product_Delete
// @Router		/products/{id} [delete]
// @Param		id				path		string					true	"id of product"
// @Param		Authorization	header		string					true	"Header must be set for valid response"
// @Success	200				{object}	ProductDeleteResponse	"success"
// @Accept		json
// @Produce	json
func ProductDelete(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")

		num, err := strconv.Atoi(product_id)

		if err != nil {

			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid param, must be a number",
			})

		}

		err = ap.DbQueries.Product_Delete(ap.Ctx, int64(num))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not delete product",
			})
		}

		return c.JSON(http.StatusOK, ProductDeleteResponse{
			Id:      product_id,
			Message: "Deleted successfully",
		})
	}
}
