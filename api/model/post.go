package model

import (
	"time"
)

type Post struct {
	Base
	User       User     `gorm:"foreignkey:userfk"`
	Userfk     uint32   `gorm:"column:userfk" json:"userfk"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Category   Category `gorm:"foreignkey:categoryfk"`
	Categoryfk uint32   `gorm:"column:categoryfk" json:"categoryfk"`
}

type PostForUser struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *Post) Prepare() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
