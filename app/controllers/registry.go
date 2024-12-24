package controllers

import (
	"instashop/app/controllers/auth"
	"instashop/app/controllers/orders"
	"instashop/app/controllers/products"
	"instashop/app/controllers/users"

	"github.com/labstack/echo/v4"
)

func Registry(e *echo.Echo) {

	e.POST("/users/", users.UserCreate)
	e.GET("/users/", users.UserGetAll)
	e.GET("/users/:id", users.UserGet)
	e.PUT("/users/:id", users.UserUpdate)
	e.DELETE("/users/:id", users.UserDelete)

	e.POST("/auth/login", auth.AuthLogin)
	e.POST("/auth/send-otp", auth.AuthSendEmailOtp)
	e.POST("/auth/verify-otp", auth.AuthVerifyOtp)
	e.POST("/auth/register", auth.AuthRegister)

	e.POST("/products/", products.ProductCreate)
	e.GET("/products/", products.ProductGetAll)
	e.GET("/products/:id", products.ProductGetOne)
	e.PUT("/products/:id", products.ProductUpdate)
	e.DELETE("/products/:id", products.ProductDelete)

	e.POST("/orders/", orders.OrderCreate)
	e.GET("/orders/", orders.OrderGetAll)
	e.GET("/orders/:id", orders.OrderGetOne)
	e.PUT("/orders/:id", orders.OrderUpdate)
	e.DELETE("/orders/:id", orders.OrderDelete)
}
