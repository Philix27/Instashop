package controllers

import (
	"instashop/app/controllers/auth"
	"instashop/app/controllers/orders"
	"instashop/app/controllers/products"
	"instashop/app/controllers/users"
	"instashop/infra/config"
	"instashop/infra/middleware"
	"instashop/infra/rbac/permissions"

	"github.com/labstack/echo/v4"
)

func Registry(e *echo.Echo, appState config.AppState) {

	e.POST("/users", users.UserCreate(appState))
	e.GET("/users/", users.UserGetAll(appState), middleware.CheckAuthorization(appState.Rbac, permissions.UserGet))
	e.GET("/users/:id", users.UserGet(appState), middleware.CheckAuthorization(appState.Rbac, permissions.UserGet))
	e.PUT("/users/:id", users.UserUpdate(appState), middleware.CheckAuthorization(appState.Rbac, permissions.UserUpdate))
	e.DELETE("/users/:id", users.UserDelete(appState), middleware.CheckAuthorization(appState.Rbac, permissions.UserDelete))

	e.POST("/auth/login", auth.AuthLogin(appState))
	e.POST("/auth/send-otp", auth.AuthSendEmailOtp(appState))
	e.POST("/auth/verify-otp", auth.AuthVerifyOtp(appState))
	e.POST("/auth/register", auth.AuthRegister(appState))

	e.POST("/products", products.ProductCreate(appState), middleware.CheckAuthorization(appState.Rbac, permissions.ProductCreate))
	e.GET("/products", products.ProductGetAll(appState), middleware.CheckAuthorization(appState.Rbac, permissions.ProductGet))
	e.GET("/products/:id", products.ProductGetOne(appState), middleware.CheckAuthorization(appState.Rbac, permissions.ProductGet))
	e.PUT("/products/:id", products.ProductUpdate(appState), middleware.CheckAuthorization(appState.Rbac, permissions.ProductUpdate))
	e.DELETE("/products/:id", products.ProductDelete(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderDelete))

	e.POST("/orders", orders.OrderCreate, middleware.CheckAuthorization(appState.Rbac, permissions.OrderCreate))
	e.GET("/orders", orders.OrderGetAll(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderGet))
	e.GET("/orders/:id", orders.OrderGetOne(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderGet))
	e.PUT("/orders/:id", orders.OrderUpdate(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderUpdate))
	e.PATCH("/orders/cancel/:id", orders.OrderCancel(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderCancel))
	e.DELETE("/orders/:id", orders.OrderDelete(appState), middleware.CheckAuthorization(appState.Rbac, permissions.OrderDelete))
}
