package middleware

import (
	"instashop/infra/crypto"
	"instashop/infra/types"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikespook/gorbac"
)

// Middleware for RBAC
func CheckAuthorization(rbac *gorbac.RBAC, requiredRole gorbac.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("Hello from user middleware")
			tokenString := c.Request().Header.Get("Authorization")[7:] //Remove "Bearer " prefix

			err, claim := crypto.ValidateAndGetTokenPayload(tokenString)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, types.ErrMsg{
					Error: "Unauthorized",
				})
			}

			println("request to /users" + claim)

			// Simulate user role (e.g., retrieve it from headers, JWT, or session)
			userRole := c.Request().Header.Get("role")
			if userRole == "" {
				return c.JSON(http.StatusForbidden, types.ErrMsg{
					Error: "role not provided"})
			}

			// Check if user has the required role
			if !rbac.IsGranted(userRole, gorbac.NewStdPermission(requiredRole.ID()), nil) {

				return c.JSON(http.StatusForbidden, types.ErrMsg{
					Error: "access denied",
				})
			}

			return next(c)
		}
	}
}
