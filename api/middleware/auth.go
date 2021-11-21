package middleware

import (
	"app/api/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustom)
		auth := claims.Authorization
		if auth != 1 {
			return c.JSON(200, "yetkiniz yok")
		}

		return next(c)
	}
}
