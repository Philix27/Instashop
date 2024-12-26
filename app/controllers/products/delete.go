package products

import (
	"instashop/infra/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductDeleteResponse struct {
	Message string `json:"message"`
}

// @Tags		Product_Delete
// @Router		/product/{id} [delete]
// @Param		id				path		string					true	"id of product"
// @Param		Authorization	header		string					true	"Header must be set for valid response"
// @Success	200				{object}	ProductDeleteResponse	"success"
// @Accept		json
// @Produce	json
func ProductDelete(appState config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")
		return c.String(http.StatusOK, product_id)
	}
}
