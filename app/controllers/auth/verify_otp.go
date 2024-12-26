package auth

import (
	"instashop/app/models"
	"instashop/infra/crypto"
	"instashop/infra/types"
	"instashop/infra/validation"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// @Tags		Auth_VerifyOtp
// @Success	200		{object}	models.VerifyOtpResponse	"success"
// @Param		request	body		models.VerifyOtpInput		true	"Auth_VerifyOtp"
// @Router		/auth/verify-otp [POST]
// @Accept		json
// @Produce	json
// @Failure	400	{object}	types.ErrMsg	"error"
func AuthVerifyOtp(c echo.Context) error {

	dto := new(models.VerifyOtpInput)

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

	return c.JSON(http.StatusOK, models.VerifyOtpResponse{
		Token: aToken,
	})
}
