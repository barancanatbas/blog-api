package main

import (
	"app/api/config"
	"app/api/repository"
	"app/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// burada database ayarları ve migration olayları var
	config.Init()

	repository.Set()
	// url yapıları burada
	router.Set(e)
}
