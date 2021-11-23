package controller

import (
	"app/api/config"
	"app/api/helper"
	"app/api/model"
	"app/api/repository"
	"app/request"
	r "app/response"
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
	users, _ := repository.Get().UserR().All()
	return r.Success(c, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := repository.Get().UserR().GetUser(id)

	if err != nil {
		return r.BadRequest(c, "hata var user bulunamadı")
	}

	return r.Success(c, &user)
}

func DeleteUser(c echo.Context) error {
	userid := helper.AuthId(&c)

	row, err := repository.Get().UserR().DeleteUser(uint(userid))

	if err != nil {
		return r.BadRequest(c, "kullanıcı silinemedi")
	}

	// çıktı
	return r.Success(c, strconv.Itoa(int(row))+" adet veri silindi")

}

func SaveUser(c echo.Context) error {

	var req request.UserInsert

	if helper.Validator(&c, &req) != nil {
		return nil
	}

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		Age:      int(req.Age),
		Job:      req.Job,
		Surname:  req.Surname,
	}
	user.Prepare()

	result, err := repository.Get().UserR().SaveUser(user)
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

	loginuser, err := repository.Get().UserR().Login(user)

	if err != nil {
		return r.BadRequest(c, "hata var user login")
	}

	// özel oluşturulmuş bir struct tan bir nesne oluşturduk
	claims := &config.JwtCustom{
		User:          *loginuser,
		Authorization: 1,
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
		"user":  loginuser,
		"time":  time.Now().Add(time.Hour * 72).Unix(),
	})
}
