package auth

import (
	"instashop/app/models"
	"instashop/db"
	"instashop/infra/config"
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/utils"
	"instashop/infra/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags		Auth_Register
// @Success	201		{object}	models.RegisterResponse	"success"
// @Param		request	body		models.RegisterInput	true	"Auth_Register"
// @Router		/auth/register [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthRegister(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		if !utils.ValidateEmail(dto.Email) {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "Invalid email adderss",
			})
		}

		if !crypto.ValidateAndCompareClaimToken(ap.Env.JwtSecretKey, dto.Token, dto.Email) {
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

		user, err := ap.DbQueries.User_Create(ap.Ctx, db.User_CreateParams{
			Email:          dto.Email,
			HashedPassword: hashedPassword,
		})

		return c.JSON(http.StatusCreated, models.RegisterResponse{
			Message: "Account for " + user.Email + "created successfully",
		})
	}
}
