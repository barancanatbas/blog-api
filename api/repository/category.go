package repository

import (
	"app/api/model"

	"github.com/jinzhu/gorm"
)

func (rootRepo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: rootRepo.Db}
}

type CategoryRepo struct {
	db *gorm.DB
}

func (c CategoryRepo) All() ([]model.Category, error) {
	categories := []model.Category{}
	err := c.db.Model(&model.Category{}).Find(&categories)

	return categories, err.Error
}

func (c CategoryRepo) Save(category model.Category) error {
	err := c.db.Debug().Model(&model.Category{}).Save(&category)
	return err.Error
}

func (c CategoryRepo) Update(id uint, category model.Category) error {
	err := c.db.Model(&model.Category{}).Where("id = ? ", id).Updates(&category)
	return err.Error
}

func (c CategoryRepo) Delete(id uint) error {
	err := c.db.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{})
	return err.Error
}

func (c CategoryRepo) Exists(id uint) int {
	category := model.Category{}
	rowsAffected := c.db.Debug().Model(&model.Category{}).Where("id = ?", id).Find(&category)

	return int(rowsAffected.RowsAffected)
}
