package controller

import (
	"app/api/config"
	"app/api/model"
	r "app/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtCustom struct {
	Name string `json:"name"`
	Auth bool   `json:"auth"`
	jwt.StandardClaims
}

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

	// gelicek olan json objesini bir stract değişkenine atadık
	user := model.User{}

	// gelen request body değerini okuttuk
	read, err := ioutil.ReadAll(c.Request().Body)
	// hata kontrolü
	if err != nil {
		return r.BadRequest(c, "bir hata oluştu")
	}
	// read daki json nesnesini bir json objesine döndürüyor ve user nesnesine atıyor.
	err = json.Unmarshal(read, &user)
	// hata kontrolü
	if err != nil {
		return r.BadRequest(c, "bir hata oluştu")
	}

	// bu şekilde id değerini aldık
	fmt.Print(user.ID)

	row, hata := user.DeleteUser(config.Database, int(user.ID))

	if hata != nil {
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
	read, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return r.BadRequest(c, "hata var read")
	}

	loginobj := LoginUser{}

	err = json.Unmarshal(read, &loginobj)

	if err != nil {
		return r.BadRequest(c, "hata var json")
	}
	user := model.User{
		Name:     loginobj.Name,
		Password: loginobj.Password,
	}

	err = user.Login(config.Database)

	if err != nil {
		return r.BadRequest(c, "hata var user login")
	}

	claims := &JwtCustom{
		loginobj.Name,
		true,
		jwt.StandardClaims{
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
		"name":  loginobj.Name,
	})
}
