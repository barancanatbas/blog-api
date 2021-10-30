package controller

import (
	r "app/response"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return r.Success(c, "Home page")
}
