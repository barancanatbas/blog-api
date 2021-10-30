package router

import (
	"app/api/controller"

	"github.com/labstack/echo/v4"
)

func Set(e *echo.Echo) {
	e.GET("/", controller.Home)

	e.GET("/post", controller.GetPosts)

	e.GET("/users", controller.GetUsers)
	e.GET("/user/:id", controller.GetUser)
	e.PUT("/user", controller.DeleteUser)
	e.POST("/user", controller.SaveUser)

	e.GET("/posts", controller.GetPosts)
	e.POST("/post", controller.GetPost)
	e.POST("/add/post", controller.SavePost)
	e.PUT("/post", controller.DeletePost)
	e.POST("/post/guncelle", controller.UpdatePost)

	e.Start(":8080")
}
