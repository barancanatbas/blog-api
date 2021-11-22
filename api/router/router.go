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

	e.POST("/register", controller.SaveUser)

	e.GET("/posts", controller.GetPosts)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_middleware.Auth)

	user := admin.Group("")
	user.PUT("/user", controller.DeleteUser)

	post := admin.Group("")
	e.GET("/post/search/:key", controller.SearchPost)
	post.POST("/add/post", controller.SavePost)
	post.PUT("/post", controller.DeletePost)
	post.POST("/post/guncelle", controller.UpdatePost)
	post.POST("/post", controller.GetPost)

	e.Start(":8080")
}
