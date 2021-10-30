package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	Base
	User    User   `gorm:"foreignkey:userfk"`
	Userfk  uint32 `gorm:"column:userfk" json:"userfk"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *Post) Prepare() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) All(db *gorm.DB) ([]Post, error) {
	posts := []Post{}

	err := db.Debug().Model(&Post{}).Preload("User").Find(&posts).Error

	if err != nil {
		return []Post{}, err
	}

	return posts, err
}

func (p *Post) GetById(db *gorm.DB, pid int) (*Post, error) {

	post := Post{}

	err := db.Debug().Model(&Post{}).Where("id = ?", pid).Preload("User").Limit(1).Take(&post).Error

	if err != nil {
		return &post, err
	}

	return &post, nil
}

func (p *Post) SavePost(db *gorm.DB) error {
	err := db.Debug().Model(&Post{}).Create(&p)

	if err != nil {
		return err.Error
	}

	return nil
}

func (p *Post) DeletePost(db *gorm.DB) error {
	err := db.Debug().Delete(&p)

	if err != nil {
		return err.Error
	}
	return nil
}

func (p *Post) UpdatePost(db *gorm.DB) error {
	err := db.Debug().Model(&Post{}).Where("id = ?", p.ID).Updates(Post{
		Title:   p.Title,
		Content: p.Content,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
