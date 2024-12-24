package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login Auth Login
//
//	@Summary		Login routes
//	@Description	Login route
//	@Tags			Login
//	@Success		200	{array}	LoginDto
//	@Router			/login [POST]
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

// AuthSendEmailOtp  AuthSendEmailOtp
//
//	@Summary		Send otp route
//	@Description	Login routes
//	@Tags			SendOtp
//	@Success		200	{string} string
//	@Router			/send-otp [POST]
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

func AuthRegister(c echo.Context) error {
	dto := new(LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto)
}
