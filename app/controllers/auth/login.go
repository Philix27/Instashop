package auth

import (
	"instashop/app/models"
	"instashop/infra/config"
	"instashop/infra/crypto"
	"instashop/infra/rbac/roles"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//	@Tags			Auth_Login
//	@Description	"You can either login as a USER or an ADMIN. Case-sensitive roles"
//	@Success		200		{object}	models.LoginResponse	"success"
//	@Param			request	body		models.LoginInput		true	"Auth_Login"
//	@Router			/auth/login [POST]
//	@Accept			json
//	@Produce		json
//	@Failure		400	{object}	types.ErrMsg	"error"
func AuthLogin(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(models.LoginInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		user, err := ap.DbQueries.User_GetByEmail(ap.Ctx, dto.Email)

		if err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "user not found",
			})
		}

		if !crypto.PasswordMatch(user.HashedPassword, dto.Password) {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: "Invalid email and password",
			})
		}
		userId := strconv.Itoa(int(user.ID))

		token, err := crypto.CreateJWTToken(ap.Env.JwtSecretKey, userId, assignRole(dto.Role), time.Hour*24)

		// todo: check db for user email and compare passwords
		return c.JSON(http.StatusOK, models.LoginResponse{
			AccessToken: token,
			UserId:      userId,
		})
	}

}

func assignRole(role string) string {
	if role == roles.Admin {
		return roles.Admin
	} else if role == roles.User {
		return roles.User
	} else {
		return roles.Guest
	}
}
