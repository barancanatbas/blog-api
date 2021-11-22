package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Base
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int    `json:"age"`
	Job      string `json:"job"`
	Password string `json:"-"`
}

type GetUserResponse struct {
	User  User
	Posts []PostForUser
}

func (u *User) Prepare() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) All(db *gorm.DB) (*[]User, error) {

	users := []User{}

	err := db.Debug().Model(&User{}).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) GetUser(db *gorm.DB, uid string) (*GetUserResponse, error) {

	var rsp = GetUserResponse{}
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&rsp.User).Error

	if err != nil {
		return &rsp, err
	}
	err = db.Debug().Model(Post{}).Where("userfk = ?", uid).Scan(&rsp.Posts).Error

	if err != nil {
		return &rsp, err
	}
	return &rsp, nil
}

func DeleteUser(db *gorm.DB, uid int, userid uint) (int64, error) {

	u := User{}
	db = db.Debug().Where("userfk = ? and id = ?", userid, uid).Delete(&u)
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, db.Error
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	result := db.Debug().Create(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) Login(db *gorm.DB) error {

	err := db.Debug().Model(&u).Where("password = ? and name = ?", u.Password, u.Name).Take(&u).Error

	if err != nil {
		return err
	}
	return nil
}
