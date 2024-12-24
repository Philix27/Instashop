package main

import (
	_ "instashop/docs"
	"instashop/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/swaggo/echo-swagger"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// PATH=$PATH:$GOROOT/bin:$GOPATH/bin

//	@title			Instashop Swagger API
//	@version		1.0
//	@description	This is a simple e-commerce server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	instashop@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		petstore.swagger.io
// @BasePath	/v2
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.POST("/users", show)

	routes.Registry(e)
	routes.PrintRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
