package routes

import "github.com/labstack/echo/v4"

func Registry(e *echo.Echo) {
	//! users
	e.POST("/users/", UserCreate)
	e.GET("/users/", UserGetAll)
	e.GET("/users/:id", UserGet)
	e.PUT("/users/:id", UserUpdate)
	e.DELETE("/users/:id", UserDelete)
	//! auth
	e.POST("/auth/login", AuthLogin)
	//! product
	e.POST("/products/", ProductCreate)
	e.GET("/products/", ProductGetAll)
	e.GET("/products/:id", ProductGetOne)
	e.PUT("/products/:id", ProductUpdate)
	e.DELETE("/products/:id", ProductDelete)
	//! orders
	e.POST("/orders/", OrderCreate)
	e.GET("/orders/", OrderGetAll)
	e.GET("/orders/:id", OrderGetOne)
	e.PUT("/orders/:id", OrderUpdate)
	e.DELETE("/orders/:id", OrderDelete)
}
