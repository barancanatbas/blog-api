package services

import (
	"app/api/config"
	"app/api/helper"
	"app/api/model"
	"app/api/repository"
	"app/request"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserServices struct{}

func UserS() UserServices {
	return UserServices{}
}

func (us UserServices) Login(c *echo.Context) (model.User, string, error) {
	var req request.UserLogin
	helper.Validator(c, &req)

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
	}

	loginuser, err := repository.Get().UserR().Login(user)

	if err != nil {
		return model.User{}, "", err
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
		return model.User{}, "", err
	}

	return *loginuser, t, nil
}

func (us UserServices) Register(c *echo.Context) (model.User, error) {
	var req request.UserInsert
	helper.Validator(c, &req)

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		Age:      int(req.Age),
		Job:      req.Job,
		Surname:  req.Surname,
	}
	user.Prepare()

	result, err := repository.Get().UserR().SaveUser(user)

	return *result, err
}

func (us UserServices) Delete(c *echo.Context) (int64, error) {
	userid := helper.AuthId(c)

	row, err := repository.Get().UserR().DeleteUser(uint(userid))

	return row, err
}

func (us UserServices) Get(c *echo.Context) (model.GetUserResponse, error) {
	id := (*c).Param("id")

	user, err := repository.Get().UserR().GetUser(id)

	return *user, err
}

func (us UserServices) All() ([]model.User, error) {
	users, err := repository.Get().UserR().All()
	return *users, err
}
