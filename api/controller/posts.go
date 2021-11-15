package controller

import (
	"app/api/config"
	"app/api/helper"
	"app/api/model"
	r "app/response"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	post := []model.Post{}
	val := helper.GetPosts(post)
	if val == true {
		val, err := redis.String(config.Client.Do("get", "posts"))
		helper.Err(err)
		err = json.Unmarshal([]byte(val), &post)
		helper.Err(err)

		return r.Success(c, post)
	} else {
		post := model.Post{}
		posts, err := post.All(config.Database)

		if err != nil {
			return r.BadRequest(c, "veriler gelmedi")
		}
		val := helper.SetPosts(&posts)
		helper.Err(err)
		if val == false {
			fmt.Println("belleğe alınmadı")
		}
		return r.Success(c, &posts)
	}

}

func GetPost(c echo.Context) error {

	read, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return r.BadRequest(c, "bir hata oluştu okuma")
	}

	reqpost := model.Post{}

	err = json.Unmarshal(read, &reqpost)

	if err != nil {
		return r.BadRequest(c, "bir hata oluştu json ")
	}

	post, err := reqpost.GetById(config.Database, int(reqpost.ID))
	if err != nil {
		return r.BadRequest(c, "bir hata oluştu veri ")
	}

	return r.Success(c, post)
}

func SavePost(c echo.Context) error {
	read, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return r.BadRequest(c, "bir hata var read")
	}

	db := model.Post{}

	err = json.Unmarshal(read, &db)

	if err != nil {
		return r.BadRequest(c, "bir hata var json unmarshal")
	}
	err = db.SavePost(config.Database)

	if err != nil {
		return r.BadRequest(c, "bir hata var veri yok")
	}

	return r.Success(c, "Ekleme başarılı")

}

func DeletePost(c echo.Context) error {
	read, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return r.BadRequest(c, "bir hata var read")
	}

	db := model.Post{}

	err = json.Unmarshal(read, &db)

	if err != nil {
		return r.BadRequest(c, "bir hata var json")
	}

	err = db.DeletePost(config.Database)

	if err != nil {
		return r.BadRequest(c, "silmedi")
	}

	return r.Success(c, "Silme işlemi başarılı")
}

func UpdatePost(c echo.Context) error {
	read, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return r.BadRequest(c, "bir hata var read")
	}

	db := model.Post{}

	err = json.Unmarshal(read, &db)

	if err != nil {
		return r.BadRequest(c, "bir hata var json")
	}

	err = db.UpdatePost(config.Database)

	if err != nil {
		return r.BadRequest(c, "bir hata update")
	}

	return r.Success(c, "başarılı")
}
