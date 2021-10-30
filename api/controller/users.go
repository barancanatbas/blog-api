package controller

import (
	"app/api/config"
	"app/api/model"
	r "app/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
