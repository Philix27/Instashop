package middleware

import (
	"instashop/infra/config"
	"instashop/infra/crypto"
	"instashop/infra/types"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikespook/gorbac"
)

// Middleware for RBAC
func CheckAuthorization(ap *config.AppState, requiredRole gorbac.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("Hello from user middleware")
			tokenString := c.Request().Header.Get("Authorization")[7:] //Remove "Bearer " prefix

			err, _, role := crypto.ValidateAndGetTokenPayload(ap.Env.JwtSecretKey, tokenString)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, types.ErrMsg{
					Error: "Unauthorized",
				})
			}

			// Simulate user role (e.g., retrieve it from headers, JWT, or session)
			userRole := c.Request().Header.Get(role)
			if userRole == "" {
				return c.JSON(http.StatusForbidden, types.ErrMsg{
					Error: "role not provided"})
			}

			// Check if user has the required role
			if !ap.Rbac.IsGranted(userRole, gorbac.NewStdPermission(requiredRole.ID()), nil) {

				return c.JSON(http.StatusForbidden, types.ErrMsg{
					Error: "access denied",
				})
			}

			return next(c)
		}
	}
}
