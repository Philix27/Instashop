package auth

import (
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type VerifyOtpInput struct {
	Token string `json:"token" validate:"required"`
	Otp   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type VerifyOtpResponse struct {
	Token string `json:"token"`
}

// @Tags		Auth_VerifyOtp
// @Success	200		{object}	VerifyOtpResponse	"success"
// @Param		request	body		VerifyOtpInput		true	"Auth_VerifyOtp"
// @Router		/auth/verify-otp [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthVerifyOtp(c echo.Context) error {

	dto := new(VerifyOtpInput)

	if err := validation.BindAndValidate(c, dto); err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: err.Error(),
		})
	}

	if !crypto.ValidateAndCompareClaimToken(dto.Token, dto.Otp) {
		return c.JSON(http.StatusBadRequest, types.ErrMsg{
			Error: "Invalid Otp",
		})
	}

	aToken, err := crypto.CreateJWTToken(dto.Email, time.Minute*30)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrMsg{
			Error: "Error in generating token",
		})
	}

	return c.JSON(http.StatusOK, VerifyOtpResponse{
		Token: aToken,
	})
}
