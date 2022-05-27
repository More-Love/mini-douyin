package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
