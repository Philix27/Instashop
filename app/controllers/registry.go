package controllers

import (
	"instashop/app/controllers/auth"
	"instashop/app/controllers/orders"
	"instashop/app/controllers/products"
	"instashop/app/controllers/users"
	"instashop/infra/middleware"

	"github.com/labstack/echo/v4"
)

func Registry(e *echo.Echo) {

	e.POST("/users/", users.UserCreate)
	e.GET("/users/", users.UserGetAll, middleware.IsAdmin)
	e.GET("/users/:id", users.UserGet, middleware.IsUser)
	e.PUT("/users/:id", users.UserUpdate, middleware.IsUser)
	e.DELETE("/users/:id", users.UserDelete, middleware.IsAdmin)

	e.POST("/auth/login", auth.AuthLogin)
	e.POST("/auth/send-otp", auth.AuthSendEmailOtp)
	e.POST("/auth/verify-otp", auth.AuthVerifyOtp)
	e.POST("/auth/register", auth.AuthRegister)

	e.POST("/products/", products.ProductCreate, middleware.IsAdmin)
	e.GET("/products/", products.ProductGetAll, middleware.IsUser)
	e.GET("/products/:id", products.ProductGetOne, middleware.IsUser)
	e.PUT("/products/:id", products.ProductUpdate, middleware.IsAdmin)
	e.DELETE("/products/:id", products.ProductDelete, middleware.IsAdmin)

	e.POST("/orders/", orders.OrderCreate, middleware.IsUser)
	e.GET("/orders/", orders.OrderGetAll, middleware.IsUser)
	e.GET("/orders/:id", orders.OrderGetOne, middleware.IsUser)
	e.PUT("/orders/:id", orders.OrderUpdate, middleware.IsUser)
	e.DELETE("/orders/:id", orders.OrderDelete, middleware.IsAdmin)
}
