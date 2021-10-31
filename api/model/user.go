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

func (u *User) GetUser(db *gorm.DB, uid string) (*User, error) {
	user := User{}

	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid int) (int64, error) {

	db = db.Debug().Delete(&u)
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
	user := User{}

	err := db.Debug().Model(&user).Where("password = ? and name = ?", u.Password, u.Name).Take(&user).Error

	if err != nil {
		return err
	}
	return nil
}
