package main

import (
	"instashop/app/controllers"
	_ "instashop/docs"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Instashop Swagger API
//	@version		1.0
//	@description	This is a simple e-commerce server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	instashop@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:1323
func main() {
	////! @BasePath	/v2
	e := echo.New()
	e.Use(middleware.Logger())
	// e.Use(middleware.CSRF())
	// e.Use(echojwt.JWT([]byte("secret")))
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net", "*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	rateLimitConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(rateLimitConfig))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Validator = &BodyValidator{validator: validator.New()}
	controllers.Registry(e)
	controllers.PrintRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}

// CustomValidator struct
type BodyValidator struct {
	validator *validator.Validate
}

// Validate method
func (cv *BodyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
