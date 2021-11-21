package model

import (
	"errors"
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

func GetById(db *gorm.DB, pid int, userid uint) (*Post, error) {

	post := Post{}

	err := db.Debug().Model(&Post{}).Where("id = ? and userfk = ?", pid, userid).Preload("User").Limit(1).Take(&post)

	if err.Error != nil {
		return &post, err.Error
	}

	if err.RowsAffected <= 0 {
		return &post, errors.New("veri yok")
	}

	return &post, nil
}

func (p *Post) SavePost(db *gorm.DB) error {
	err := db.Model(&Post{}).Create(&p)

	if err != nil {
		return err.Error
	}

	return nil
}

func DeletePost(db *gorm.DB, id uint32, userid uint32) error {
	post := &Post{}
	err := db.Debug().Where("id = ? and userfk = ?", id, userid).Delete(&post)

	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return errors.New("bilgiler silinmedi")
	}
	return nil
}

func (p *Post) UpdatePost(db *gorm.DB, id uint) error {
	val := db.Model(&Post{}).Where("id = ? and userfk = ?", id, p.Userfk).Updates(Post{
		Title:   p.Title,
		Content: p.Content,
	})
	if val.Error != nil {
		return val.Error
	}
	if val.RowsAffected <= 0 {
		return errors.New("id değeri geçerli değil")
	}
	return nil
}
