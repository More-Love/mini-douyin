package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-douyin/models"
)



var (
	UserRepo *UserRepository
	VideoRepo *VideoRepository
	TokenRepo *TokenRepository
)

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	models := []any {
		&models.User{},
		&models.Video{},
		&models.Comment{},
		&models.Followship{},
	}

	for _, model := range models {
		err = db.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}

	UserRepo = &UserRepository{db}
	VideoRepo = &VideoRepository{db}
	

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	TokenRepo = &TokenRepository{rdb, context.Background()}
}
