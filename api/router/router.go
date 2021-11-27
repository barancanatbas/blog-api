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
	e.GET("/post/:id", controller.GetPosts)
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
	post.POST("/add/post", controller.SavePost)
	post.PUT("/post", controller.DeletePost)
	post.POST("/post/guncelle", controller.UpdatePost)
	post.POST("/post", controller.GetPost)

	category := admin.Group("")
	category.POST("/category", controller.SaveCategory)
	category.DELETE("/category", controller.DeleteCategory)
	category.PUT("/category", controller.UpdateCategory)

	e.Start(":8080")
}
