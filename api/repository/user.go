package repository

import (
	"app/api/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

func (rootRepo *Repositories) UserR() UserRepo {
	return UserRepo{db: rootRepo.Db}
}

type UserRepo struct {
	db *gorm.DB
}

func (repo UserRepo) All() (*[]model.User, error) {

	users := []model.User{}

	err := repo.db.Debug().Model(&model.User{}).Find(&users).Error
	fmt.Println("burdasın")
	if err != nil {
		return &[]model.User{}, err
	}
	return &users, nil
}

func (repo UserRepo) GetUser(uid string) (*model.GetUserResponse, error) {

	var rsp = model.GetUserResponse{}
	err := repo.db.Debug().Model(model.User{}).Where("id = ?", uid).Take(&rsp.User).Error

	if err != nil {
		return &rsp, err
	}
	err = repo.db.Debug().Model(model.Post{}).Where("userfk = ?", uid).Scan(&rsp.Posts).Error

	if err != nil {
		return &rsp, err
	}
	return &rsp, nil
}

func (repo UserRepo) DeleteUser(userid uint) (int64, error) {

	u := model.User{}
	err := repo.db.Debug().Where("id = ?", userid).Delete(&u)
	if err.Error != nil {
		return 0, err.Error
	}

	return err.RowsAffected, err.Error
}

func (repo UserRepo) SaveUser(user model.User) (*model.User, error) {

	result := repo.db.Debug().Create(&user)
	if result.Error != nil {
		return &model.User{}, result.Error
	}

	return &user, nil
}

func (repo UserRepo) Login(user model.User) (*model.User, error) {

	loginuser := model.User{}
	err := repo.db.Debug().Model(&model.User{}).Where("password = ? and name = ?", user.Password, user.Name).Take(&loginuser).Error

	if err != nil {
		return &loginuser, err
	}
	return &loginuser, nil
}
