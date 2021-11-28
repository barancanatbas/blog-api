package services

import (
	"app/api/helper"
	"app/api/model"
	"app/api/repository"
	"app/request"
	"errors"

	"github.com/labstack/echo/v4"
)

type CategoryServices struct{}

func CategoryS() CategoryServices {
	return CategoryServices{}
}

func (cs CategoryServices) All() ([]model.Category, error) {
	categories, err := repository.Get().Category().All()
	return categories, err
}

func (cs CategoryServices) Save(c *echo.Context) error {
	var req request.CategorySave
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	category := model.Category{
		Name: req.Name,
	}

	err := repository.Get().Category().Save(category)
	return err
}

func (cs CategoryServices) Delete(c *echo.Context) error {
	var req request.CategoryDelete
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	err := repository.Get().Category().Delete(req.Id)
	return err
}

func (cs CategoryServices) Update(c *echo.Context) error {
	var req request.CategoryUpdate
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	category := model.Category{
		Name: req.Name,
	}

	err := repository.Get().Category().Update(uint(req.Id), category)
	return err
}
