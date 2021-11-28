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

	e.POST("/login", controller.Login)
	e.POST("/register", controller.SaveUser)

	e.GET("/posts", controller.GetPosts)
	e.GET("/post/:id", controller.GetPost)
	e.GET("/users", controller.GetUsers)
	e.GET("/user/:id", controller.GetUser)

	e.GET("/categories", controller.AllCategory)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_middleware.Auth)

	user := admin.Group("")
	user.PUT("/user", controller.DeleteUser)

	post := admin.Group("")
	e.GET("/post/search/:key", controller.SearchPost)
	post.POST("/post", controller.SavePost)
	post.DELETE("/post", controller.DeletePost)
	post.PUT("/post", controller.UpdatePost)

	category := admin.Group("")
	category.POST("/category", controller.SaveCategory)
	category.DELETE("/category", controller.DeleteCategory)
	category.PUT("/category", controller.UpdateCategory)

	e.Start(":8080")
}
