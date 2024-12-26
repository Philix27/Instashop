package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductUpdateResponse struct {
	Message string `json:"message"`
}

// @Tags		Product_Update
// @Router		/product/{id} [patch]
// @Param		id				path	string	true	"id of product"
// @Param		Authorization	header	string	true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
// @Success	200	{object}	ProductUpdateResponse	"success"
func ProductUpdate(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
