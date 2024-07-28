package models

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB = Init()
var RDB = InitRedis()

func Init() *gorm.DB {
	dsn := "root:yangchen22@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm init error: ", err)
	}
	return db
}

func InitRedis() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
