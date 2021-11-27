package controller

import (
	"app/api/services"
	r "app/response"

	"github.com/labstack/echo/v4"
)

func AllCategory(c echo.Context) error {

	categories, err := services.CategoryS().All()
	if err != nil {
		return r.BadRequest(c, err)
	}

	return r.Success(c, categories)
}

func SaveCategory(c echo.Context) error {
	err := services.CategoryS().Save(&c)
	if err != nil {
		return r.BadRequest(c, "başarısız")
	}

	return r.Success(c, "ekleme başarılı")
}

func DeleteCategory(c echo.Context) error {
	err := services.CategoryS().Delete(&c)

	if err != nil {
		return r.BadRequest(c, "silme başarısız")
	}

	return r.Success(c, "silme başarılı")
}

func UpdateCategory(c echo.Context) error {
	err := services.CategoryS().Update(&c)

	if err != nil {
		return r.BadRequest(c, "güncelleme başarısız")
	}

	return r.Success(c, "güncelleme başarılı")

}
