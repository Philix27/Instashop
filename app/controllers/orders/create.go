package orders

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

type OrderCreateInput struct {
	Quantity  int32 `json:"quantity"`
	ProductID int32 `json:"productId"`
	OrderId   int32 `json:"orderId"`
}
type OrderCreateResponse struct {
	Message string `json:"message"`
}

// @Tags		Order_Create
// @Router		/orders [post]
// @Param		Authorization	header		string				true	"Header must be set for valid response"
// @Param		request			body		OrderCreateInput	true	"Order_Create"
// @Success	200				{object}	OrderCreateResponse	"success"
// @Accept		json
// @Produce	json
func OrderCreate(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(OrderCreateInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}
		userId := c.Get("userId").(string)
		userIdNum, err := strconv.Atoi(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{
				Error: "Invalid user id",
			})
		}

		//! Check if an order exist
		orders, err := ap.DbQueries.Orders_GetAllByUserId(ap.Ctx, int32(userIdNum))

		if len(orders) < 1 {
			//! Create new order
			order, err := ap.DbQueries.Orders_Create(ap.Ctx, int32(userIdNum))

			if err != nil {
				return c.JSON(http.StatusInternalServerError, types.ErrMsg{
					Error: "Could not create order",
				})
			}

			//! Add an order item
			err = ap.DbQueries.OrderItem_Create(ap.Ctx, db.OrderItem_CreateParams{
				Quantity: dto.Quantity,
				OrderID: pgtype.Int4{
					Int32: int32(order.ID),
				},
				ProductID: pgtype.Int4{
					Int32: dto.ProductID,
				},
				UserID: pgtype.Int4{
					Int32: int32(userIdNum),
				},
			})
		}

		//! Add an order item
		err = ap.DbQueries.OrderItem_Create(ap.Ctx, db.OrderItem_CreateParams{
			Quantity: dto.Quantity,
			OrderID: pgtype.Int4{
				Int32: int32(dto.OrderId),
			},
			ProductID: pgtype.Int4{
				Int32: dto.ProductID,
			},
			UserID: pgtype.Int4{
				Int32: int32(userIdNum),
			},
		})

		return c.JSON(http.StatusCreated, OrderCreateResponse{
			Message: "Order created successfully",
		})
	}
}
