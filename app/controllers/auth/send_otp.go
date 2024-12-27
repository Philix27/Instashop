package auth

import (
	"instashop/app/models"
	"instashop/infra/config"
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// @Tags		Auth_SendEmailOtp
// @Success	200		{object}	models.SendEmailResponse	"success"
// @Param		request	body		models.SendEmailOtpInput	true	"Auth_SendEmailOtp"
// @Router		/auth/send-otp [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthSendEmailOtp(ap config.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(models.SendEmailOtpInput)

		if err := validation.BindAndValidate(c, dto); err != nil {
			return c.JSON(http.StatusBadRequest, types.ErrMsg{
				Error: err.Error(),
			})
		}

		otpValue := crypto.GenerateOTP()
		token, err := crypto.CreateJWTToken(ap.Env.JwtSecretKey, otpValue, "guest", time.Minute*10)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, types.ErrMsg{Error: "Something went wrong"})
		}

		println("Otp value", otpValue)

		//  todo: Send otp value to user email via a desired email provider such as resend or twilo

		return c.JSON(http.StatusOK, models.SendEmailResponse{
			Token: token,
			Otp:   otpValue,
		})
	}
}
