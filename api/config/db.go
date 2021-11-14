package config

import (
	"app/api/model"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// bu nesneyi model kısmında çağırırım
var Database *gorm.DB
var pool = newPool()

var Client = pool.Get()

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
		&model.Post{},
	)

	Database.Model(&model.Post{}).AddForeignKey("userfk", "users(id)", "cascade", "cascade")

	return migrate
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
