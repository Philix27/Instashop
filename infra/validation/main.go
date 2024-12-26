package validation

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Request struct with validation tags
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18,lte=100"`
}

// CustomValidator struct
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)

	// Bind request
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "user created successfully"})
}

func main() {
	e := echo.New()

	// Set up custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	e.POST("/users", CreateUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
