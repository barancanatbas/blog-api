package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustom struct {
	Name          string `json:"name"`
	Authorization uint   `json:"authorization"` // 1 = admin, 2 = sınırlı kullanıcı
	ID            uint   `json:"id"`
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustom{},
	SigningKey: []byte("secret"),
}
