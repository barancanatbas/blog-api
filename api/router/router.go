package router

import (
	"app/api/config"
	"app/api/controller"
	_middleware "app/api/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Set(e *echo.Echo) {
	e.GET("/", controller.Home)

	e.GET("/post", controller.GetPosts)

	e.POST("/login", controller.Login)
	e.GET("/users", controller.GetUsers)
	e.GET("/user/:id", controller.GetUser)
	e.PUT("/user", controller.DeleteUser)
	e.POST("/user", controller.SaveUser)

	e.GET("/posts", controller.GetPosts)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_middleware.Auth)

	admin.POST("/add/post", controller.SavePost)
	admin.PUT("/post", controller.DeletePost)
	admin.POST("/post/guncelle", controller.UpdatePost)
	admin.POST("/post", controller.GetPost)

	e.Start(":8080")
}
