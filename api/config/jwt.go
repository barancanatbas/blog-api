package config

import (
	"app/api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustom struct {
	User          model.User `json:"user"`
	Authorization uint       `json:"authorization"` // 1 = admin, 2 = sınırlı kullanıcı
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustom{},
	SigningKey: []byte("secret"),
}
