package services

import (
	"app/api/helper"
	"app/api/model"
	"app/api/repository"
	"app/request"

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
	var req request.PostDeleteReq
	helper.Validator(c, &req)

	post, err := repository.Get().Post().GetById(uint(req.ID))

	return err, post
}

// save new post
func (ps PostServices) Save(c *echo.Context) error {
	// validate
	var req request.PostReq
	helper.Validator(c, &req)

	id := helper.AuthId(c)
	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		Userfk:  id,
	}

	err := repository.Get().Post().Save(post)

	return err
}

// delete post
func (ps PostServices) Delete(c *echo.Context) error {

	var req request.PostDeleteReq
	helper.Validator(c, &req)

	userid := helper.AuthId(c)

	err := repository.Get().Post().Delete(uint(req.ID), uint(userid))
	return err
}

// update post
func (ps PostServices) Update(c *echo.Context) error {

	var req request.PostUpdateReq
	helper.Validator(c, &req)

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
	helper.Validator(c, &req)

	posts, err := repository.Get().Post().Search(req.Key)
	return posts, err
}
