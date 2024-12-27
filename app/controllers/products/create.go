package products

import (
	"instashop/db"
	"instashop/infra/config"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ProductCreateInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	ImageUrl    string `json:"imageUrl"`
}

type ProductCreateResponse struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Price       int32  `json:"price"`
}

// @Tags		Product_Create
// @Success	200	{string}	string
// @Router		/products [post]
// @Success	200				{object}	ProductCreateResponse	"success"
// @Param		request			body		ProductCreateInput		true	"Product_Create"
// @Param		Authorization	header		string					true	"Header must be set for valid response"
// @Accept		json
// @Produce	json
func ProductCreate(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {

		dto := new(ProductCreateInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		product, err := ap.DbQueries.Products_Create(ap.Ctx, db.Products_CreateParams{
			Title: dto.Title,
			Description: pgtype.Text{
				String: dto.Description,
				Valid:  true,
			},
			ImageUrl: pgtype.Text{
				String: dto.ImageUrl,
				Valid:  true,
			},
			Price: int32(dto.Price),
			Stock: int32(dto.Stock),
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Could not create product",
			})
		}

		// return c.String(http.StatusOK, product_id)
		return c.JSON(http.StatusCreated, ProductCreateResponse{
			Id:          product.ID,
			Title:       product.Title,
			Price:       product.Price,
			Description: product.Description.String,
			ImageUrl:    product.ImageUrl.String,
		})
	}
}
