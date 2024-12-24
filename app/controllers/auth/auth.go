package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Tags		Auth_Login
// @Success	200	{array}	LoginDto
// @Router		/login [POST]
func AuthLogin(c echo.Context) error {
	dto := new(LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto)
}

type SendEmailDto struct {
	Email string `json:"email"`
}

// @Tags		Auth_SendEmailOtp
// @Success	200	{string}	string
// @Router		/auth/send-otp [POST]
func AuthSendEmailOtp(c echo.Context) error {
	dto := new(LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto)
}

type VerifyOtpDtp struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
	Token string `json:"token"`
}

// @Tags		Auth_VerifyOtp
// @Success	200	{string}	string
// @Router		/auth/verify-otp [POST]
func AuthVerifyOtp(c echo.Context) error {
	dto := new(LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto)
}

type RegisterDto struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// @Tags		Auth_Register
// @Success	200	{string}	string
// @Router		/auth/register [POST]
func AuthRegister(c echo.Context) error {
	dto := new(LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto)
}
