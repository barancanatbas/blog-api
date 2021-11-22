package controller

import (
	"app/api/config"
	"app/api/helper"
	"app/api/model"
	"app/request"
	r "app/response"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// özel bir jwt struct

func GetUsers(c echo.Context) error {

	user := model.User{}

	users, err := user.All(config.Database) // model config kısmında yer alan database nesnesi burada

	if err != nil {
		return r.BadRequest(c, "hata var veri gelmedi")
	}

	return r.Success(c, &users)

}

func GetUser(c echo.Context) error {
	db := model.User{}
	id := c.Param("id")

	user, err := db.GetUser(config.Database, id)

	if err != nil {
		return r.BadRequest(c, "hata var user bulunamadı")
	}

	return r.Success(c, &user)
}

func DeleteUser(c echo.Context) error {

	var req request.UserDelRequest

	if helper.Validator(&c, &req) != nil {
		return nil
	}
	userid := helper.AuthId(&c)

	row, err := model.DeleteUser(config.Database, int(req.ID), uint(userid))

	if err != nil {
		return r.BadRequest(c, "kullanıcı silinemedi")
	}

	// çıktı
	return r.Success(c, strconv.Itoa(int(row))+" adet veri silindi")

}

func SaveUser(c echo.Context) error {
	user := model.User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return r.BadRequest(c, "veriler alınamadı")
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return r.BadRequest(c, "veriler dönüştürülmedi")
	}

	user.Prepare()

	result, err := user.SaveUser(config.Database)
	if err != nil {
		return r.BadRequest(c, "veri eklenemedi")
	}

	return r.Success(c, result)
}

func Login(c echo.Context) error {
	var req request.UserLogin
	if helper.Validator(&c, &req) != nil {
		return nil
	}

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
	}

	err := user.Login(config.Database)

	if err != nil {
		return r.BadRequest(c, "hata var user login")
	}

	// özel oluşturulmuş bir struct tan bir nesne oluşturduk
	claims := &config.JwtCustom{
		Name:          req.Name,
		Authorization: 1,
		ID:            uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return r.Success(c, echo.Map{
		"token": &t,
		"name":  req.Name,
	})
}
