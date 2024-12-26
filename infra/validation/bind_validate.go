package validation

import "github.com/labstack/echo/v4"

func BindAndValidate(c echo.Context, dto interface{}) error {
	if err := c.Bind(dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return err
	}
	return nil
}
