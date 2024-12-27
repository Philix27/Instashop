package products

import (
	"instashop/db"
	"instashop/infra/config"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ProductUpdateInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int32  `json:"price" validate:"required"`
	Stock       int32  `json:"stock" validate:"required"`
	ImageUrl    string `json:"imageUrl"`
}
type ProductUpdateResponse struct {
	Message string `json:"message"`
}

//	@Tags		Product_Update
//	@Router		/products/{id} [patch]
//	@Param		id				path	string	true	"id of product"
//	@Param		Authorization	header	string	true	"Header must be set for valid response"
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	ProductUpdateResponse	"success"
//	@Param		request	body		ProductUpdateInput		true	"Product_Update"
func ProductUpdate(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		product_id := c.Param("id")
		num, err := strconv.Atoi(product_id)

		dto := new(ProductUpdateInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		err = ap.DbQueries.Product_Update(ap.Ctx, db.Product_UpdateParams{
			ID:    int64(num),
			Title: dto.Title,
			Description: pgtype.Text{
				String: dto.Description,
				Valid:  true,
			},
			ImageUrl: pgtype.Text{
				String: dto.ImageUrl,
				Valid:  true,
			},
			Stock: int32(dto.Stock),
			Price: dto.Price,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not create product",
			})
		}

		// return c.String(http.StatusOK, product_id)
		return c.JSON(http.StatusOK, ProductUpdateResponse{
			Message: "Updated",
		})
	}
}
