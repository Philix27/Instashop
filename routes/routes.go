package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func PrintRoutes(e *echo.Echo) {
	routes := e.Routes()

	fmt.Println("Retrieved Routes:")
	for _, route := range routes {
		fmt.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}
