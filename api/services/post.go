package services

import (
	"app/api/helper"
	"app/api/model"
	"app/api/repository"
	"app/request"
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostS() PostServices {
	return PostServices{}
}

type PostServices struct{}

// all post
func (ps PostServices) All() (error, []model.Post) {
	posts, err := repository.Get().Post().All()

	return err, posts
}

// only one post by id
func (ps PostServices) Get(c *echo.Context) (error, model.Post) {
	// validate işlemi
	id, _ := strconv.Atoi((*c).Param("id"))

	post, err := repository.Get().Post().GetById(uint(id))

	return err, post
}

// save new post
func (ps PostServices) Save(c *echo.Context) error {
	// validate
	var req request.PostReq
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	// category exists ?
	rowsAffected := repository.Get().Category().Exists(req.CategoryFK)
	if rowsAffected <= 0 {
		fmt.Println(rowsAffected)
		return errors.New("kategori bulunamadı")
	}

	id := helper.AuthId(c)
	post := model.Post{
		Title:      req.Title,
		Content:    req.Content,
		Userfk:     id,
		Categoryfk: uint32(req.CategoryFK),
	}

	err := repository.Get().Post().Save(&post)

	return err
}

// delete post
func (ps PostServices) Delete(c *echo.Context) error {

	var req request.PostDeleteReq
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	userid := helper.AuthId(c)

	err := repository.Get().Post().Delete(uint(req.ID), uint(userid))
	return err
}

// update post
func (ps PostServices) Update(c *echo.Context) error {

	var req request.PostUpdateReq
	if val := helper.Validator(c, &req); val != "" {
		return errors.New(val)
	}

	userid := helper.AuthId(c)

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		Userfk:  userid,
	}
	err := repository.Get().Post().Update(uint(req.ID), post)
	return err
}

func (ps PostServices) Search(c *echo.Context) ([]model.Post, error) {
	var req request.PostSearchReq
	if val := helper.Validator(c, &req); val != "" {
		return []model.Post{}, errors.New(val)
	}

	posts, err := repository.Get().Post().Search(req.Key)
	return posts, err
}
