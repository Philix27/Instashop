package middleware

import (
	"instashop/infra/crypto"
	"instashop/infra/types"
	"net/http"

	"github.com/labstack/echo/v4"
)



func IsUser(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		println("Hello from user middleware")
		tokenString := c.Request().Header.Get("Authorization")[7:] // Remove "Bearer " prefix

		err, claim := crypto.ValidateAndGetTokenPayload(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, types.ErrMsg{
				Error: "Unauthorized",
			})
		}

		println("request to /users" + claim)
		return next(c)
	}
}
