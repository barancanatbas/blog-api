package config

import (
	"app/api/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
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
	godotenv.Load(".env")

	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	DBNAME := os.Getenv("DBNAME")

	db, err := gorm.Open("mysql", USER+":"+PASSWORD+HOST+DBNAME+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	return db
}

// otomatik migrate yapar
func AutoMigrate() *gorm.DB {
	migrate := Database.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Post{},
	)

	Database.Model(&model.Post{}).AddForeignKey("userfk", "users(id)", "cascade", "cascade")
	Database.Model(&model.Post{}).AddForeignKey("categoryfk", "categories(id)", "cascade", "cascade")

	return migrate
}
