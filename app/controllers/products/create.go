package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductCreateInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int `json:"price" validate:"required"`
}

type ProductCreateResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

//	@Tags		Product_Create
//	@Success	200	{string}	string
//	@Router		/products [post]
//	@Success	200				{object}	ProductCreateResponse	"success"
//	@Param		request			body		ProductCreateInput		true	"Product_Create"
//	@Param		Authorization	header		string					true	"Header must be set for valid response"
//	@Accept		json
//	@Produce	json
func ProductCreate(c echo.Context) error {
	product_id := c.Param("id")
	return c.String(http.StatusOK, product_id)
}
