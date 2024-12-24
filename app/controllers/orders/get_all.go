package orders

import (
	"github.com/labstack/echo/v4"
)

// @Tags		Order_GetAll
// @Success	200	{string}	string
// @Router		/orders/{id} [get]
func OrderGetAll(c echo.Context) error { return nil }
