package controller

import (
	"app/api/services"
	r "app/response"

	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	err, posts := services.PostS().All()
	if err != nil {
		return r.BadRequest(c, err)
	}

	return r.Success(c, &posts)
}

func GetPost(c echo.Context) error {
	err, post := services.PostS().Get(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, post)
}

func SavePost(c echo.Context) error {
	err := services.PostS().Save(&c)
	if err != nil {
		return r.BadRequest(c, "bir hata var veri yok")
	}

	return r.Success(c, "Ekleme başarılı")
}

func DeletePost(c echo.Context) error {
	err := services.PostS().Delete(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}

	return r.Success(c, "Silme işlemi başarılı")
}

func UpdatePost(c echo.Context) error {
	err := services.PostS().Update(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, "başarılı")
}

func SearchPost(c echo.Context) error {
	posts, err := services.PostS().Search(&c)
	if err != nil {
		return r.BadRequest(c, err)
	}

	return r.Success(c, posts)
}
