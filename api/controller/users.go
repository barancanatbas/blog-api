package controller

import (
	"app/api/services"
	r "app/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// özel bir jwt struct

func GetUsers(c echo.Context) error {
	users, err := services.UserS().All()
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, users)
}

func GetUser(c echo.Context) error {
	user, err := services.UserS().Get(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, &user)
}

func DeleteUser(c echo.Context) error {
	row, err := services.UserS().Delete(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, strconv.Itoa(int(row))+" adet veri silindi")
}

func SaveUser(c echo.Context) error {
	result, err := services.UserS().Register(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, result)
}

func Login(c echo.Context) error {
	loginuser, t, err := services.UserS().Login(&c)
	if err != nil {
		return r.BadRequest(c, err.Error())
	}
	return r.Success(c, echo.Map{
		"token": &t,
		"user":  loginuser,
		"time":  time.Now().Add(time.Hour * 72).Unix(),
	})
}
