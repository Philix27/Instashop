package auth

import (
	"instashop/app/models"
	"instashop/infra/config"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags		Auth_Login
// @Success	200		{object}	models.LoginResponse	"success"
// @Param		request	body		models.LoginInput		true	"Auth_Login"
// @Router		/auth/login [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthLogin(appState config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(models.LoginInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}
		// todo: check db for user email and compare passwords
		return c.JSON(http.StatusCreated, dto)
	}

}
