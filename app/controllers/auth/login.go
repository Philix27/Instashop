package auth

import (
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	UserId      string `json:"UserId"`
	AccessToken string `json:"password"`
}

// @Tags		Auth_Login
// @Success	200		{object}	LoginResponse	"success"
// @Param		request	body		LoginInput		true	"Auth_Login"
// @Router		/auth/login [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthLogin(c echo.Context) error {
	dto := new(LoginInput)

	if err := validation.BindAndValidate(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: err.Error(),
		})
	}
	// todo: check db for user email and compare passwords
	return c.JSON(http.StatusCreated, dto)
}
