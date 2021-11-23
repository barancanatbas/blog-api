package main

import (
	"app/api/config"
	"app/api/repository"
	"app/api/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// burada database ayarları ve migration olayları var
	config.Init()

	repository.Set()
	// url yapıları burada
	router.Set(e)
}
