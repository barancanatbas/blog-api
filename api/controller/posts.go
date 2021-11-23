package controller

import (
	"app/api/config"
	"app/api/helper"
	"app/api/model"
	"app/request"
	r "app/response"
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	post := model.Post{}

	posts, err := post.All(config.Database)

	if err != nil {
		return r.BadRequest(c, "veriler gelmedi")
	}

	return r.Success(c, &posts)

}

func GetPost(c echo.Context) error {
	var req request.PostDeleteReq

	if helper.Validator(&c, &req) != nil {
		return nil
	}
	userid := helper.AuthId(&c)

	post, err := model.GetById(config.Database, int(req.ID), uint(userid))
	if err != nil {
		return r.BadRequest(c, err.Error())
	}

	return r.Success(c, post)
}

func SavePost(c echo.Context) error {

	var req request.PostReq

	if helper.Validator(&c, &req) != nil {
		return r.BadRequest(c, "hata var")
	}
	id := helper.AuthId(&c)
	db := model.Post{
		Title:   req.Title,
		Content: req.Content,
		Userfk:  id,
	}

	fmt.Println(id)
	err := db.SavePost(config.Database)

	if err != nil {
		return r.BadRequest(c, "bir hata var veri yok")
	}

	return r.Success(c, "Ekleme başarılı")

}

func DeletePost(c echo.Context) error {

	var req request.PostDeleteReq

	if helper.Validator(&c, &req) != nil {
		return nil
	}
	userid := helper.AuthId(&c)
	err := model.DeletePost(config.Database, req.ID, userid)

	if err != nil {
		return r.BadRequest(c, err.Error())
	}

	return r.Success(c, "Silme işlemi başarılı")
}

func UpdatePost(c echo.Context) error {

	var req request.PostUpdateReq

	// validate işlemi
	if helper.Validator(&c, &req) != nil {
		return nil
	}
	userid := helper.AuthId(&c)

	db := model.Post{
		Title:   req.Title,
		Content: req.Content,
		Userfk:  userid,
	}
	// güncelleme işlemi
	err := db.UpdatePost(config.Database, uint(req.ID))

	if err != nil {
		return r.BadRequest(c, err.Error())
	}

	// response
	return r.Success(c, "başarılı")
}

func SearchPost(c echo.Context) error {

	var req request.PostSearchReq

	if helper.Validator(&c, &req) != nil {
		return nil
	}

	post, err := model.SearchPost(config.Connect(), req.Key)
	if err != nil {
		return r.BadRequest(c, err)
	}

	return r.Success(c, post)
}
