package model

import (
	"time"
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
