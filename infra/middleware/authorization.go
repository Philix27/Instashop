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
			tokenString := c.Request().Header.Get("Authorization")[7:] //Remove "Bearer " prefix

			err, userId, userRole := crypto.ValidateAndGetTokenPayload(ap.Env.JwtSecretKey, tokenString)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, types.ErrMsg{
					Error: "Unauthorized",
				})
			}

			// if userRole != roles.Admin || userRole != roles.User {
			// 	return c.JSON(http.StatusForbidden, types.ErrMsg{
			// 		Error: "invalid role provided " + userRole,
			// 	})
			// }

			println("Gonna: ", userRole)
			println("RequiredRole: ", requiredRole.ID())
			// Check if user has the required role
			if !ap.Rbac.IsGranted(userRole, gorbac.NewStdPermission(requiredRole.ID()), nil) {

				return c.JSON(http.StatusForbidden, types.ErrMsg{
					Error: "access denied",
				})
			}

			c.Set("userRole", userRole)
			c.Set("userId", userId)
			return next(c)
		}
	}
}
