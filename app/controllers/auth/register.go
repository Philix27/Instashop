package auth

import (
	"instashop/app/models"
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

//	@Tags		Auth_Register
//	@Success	200		{object}	models.RegisterResponse	"success"
//	@Param		request	body		models.RegisterInput	true	"Auth_Register"
//	@Router		/auth/register [POST]
//	@Accept		json
//	@Produce	json
//	@Failure	400	{object}	types.ErrMsg	"error"
func AuthRegister(c echo.Context) error {
	dto := new(models.RegisterInput)

	if err := validation.BindAndValidate(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: err.Error(),
		})
	}

	if dto.Password != dto.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: "Please confirm password",
		})
	}

	if !crypto.ValidateAndCompareClaimToken(dto.Token, dto.Email) {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: "Invalid token",
		})
	}

	hashedPassword, err := crypto.HashPassword(dto.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrMsg{
			Error: "Invalid token",
		})
	}

	// todo: create user on db
	println("hashedPassword", hashedPassword)
	return c.JSON(http.StatusCreated, models.RegisterResponse{
		Message: "Account for " + dto.Email + "created successfully",
	})
}
