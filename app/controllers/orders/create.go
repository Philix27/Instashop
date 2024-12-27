package orders

import (
	"instashop/infra/config"

	"github.com/labstack/echo/v4"
)

type OrderCreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type OrderCreateResponse struct {
	UserId      string `json:"UserId"`
	AccessToken string `json:"password"`
}

// @Tags		Order_Create
// @Router		/orders [post]
// @Param		Authorization	header		string				true	"Header must be set for valid response"
// @Param		request			body		OrderCreateInput	true	"Order_Create"
// @Success	200				{object}	OrderCreateResponse	"success"
// @Accept		json
// @Produce	json
func OrderCreate(appState config.AppState) echo.HandlerFunc { return nil }
