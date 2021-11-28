package repository

import (
	"app/api/model"

	"github.com/jinzhu/gorm"
)

func (rootRepo Repositories) Post() PostRepo {
	return PostRepo{db: rootRepo.Db}
}

type PostRepo struct {
	db *gorm.DB
}

func (p PostRepo) All() ([]model.Post, error) {
	posts := []model.Post{}
	err := p.db.Model(&model.Post{}).Preload("User").Preload("Category").Find(&posts)

	return posts, err.Error
}

func (p PostRepo) GetById(postId uint) (model.Post, error) {
	post := model.Post{}
	err := p.db.Model(&model.Post{}).Preload("User").Where("id = ?", postId).Take(&post)

	return post, err.Error
}

func (p PostRepo) Search(key string) ([]model.Post, error) {
	post := []model.Post{}
	err := p.db.Model(&model.Post{}).Where("title like ?", "%"+key+"%").Find(&post)

	return post, err.Error
}

func (p PostRepo) Delete(id uint, userid uint) error {
	err := p.db.Model(&model.Post{}).Where("id = ? and userfk = ?", id, userid).Delete(&model.Post{})

	return err.Error
}

func (p PostRepo) Update(id uint, post model.Post) error {
	err := p.db.Model(&model.Post{}).Where("id = ? and userfk = ?", id, post.Userfk).Updates(post)

	return err.Error
}

func (p PostRepo) Save(post *model.Post) error {
	err := p.db.Debug().Model(&model.Post{}).Save(&post)
	return err.Error
}
