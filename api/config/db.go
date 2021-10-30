package config

import (
	"app/api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// bu nesneyi model kısmında çağırırım
var Database *gorm.DB

// vt ayarlarını çalıştırma methodu
func Init() {
	Database = Connect()
	AutoMigrate()
}

// bağlantı açar
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql123@/restapideneme?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	return db
}

// otomatik migrate yapar
func AutoMigrate() *gorm.DB {
	migrate := Database.AutoMigrate(
		&model.User{},
		&model.Post{},
	)

	Database.Model(&model.Post{}).AddForeignKey("userfk", "users(id)", "cascade", "cascade")

	return migrate
}
