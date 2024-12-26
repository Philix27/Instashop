package auth

import (
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SendEmailOtpInput struct {
	Email string `json:"email" validate:"required"`
}

type SendEmailResponse struct {
	Token string `json:"token"`
	// todo: remove otp
	Otp string `json:"otp"`
}

// @Tags		Auth_SendEmailOtp
// @Success	200		{object}	SendEmailResponse	"success"
// @Param		request	body		SendEmailOtpInput	true	"Auth_SendEmailOtp"
// @Router		/auth/send-otp [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthSendEmailOtp(c echo.Context) error {
	dto := new(SendEmailOtpInput)

	if err := validation.BindAndValidate(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: err.Error(),
		})
	}

	otpValue := crypto.GenerateOTP()
	token, err := crypto.CreateJWTToken(otpValue, time.Minute*10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrMsg{Error: "Something went wrong"})
	}

	println("Otp value", otpValue)

	//  todo: Send otp value to user email via a desired email provider such as resend or twilo

	return c.JSON(http.StatusOK, SendEmailResponse{
		Token: token,
		Otp: otpValue,
	})
}
