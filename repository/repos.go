package repository

import (
	"context"
	"mini-douyin/config"
	"mini-douyin/models"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UserRepo  *UserRepository
	VideoRepo *VideoRepository
	TokenRepo *TokenRepository
)

func init() {

	db, err := gorm.Open(postgres.Open(config.Config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Video{},
		&models.User{},
		&models.Comment{},
		&models.Followship{},
		&models.Favorite{},
	)

	if err != nil {
		panic(err)
	}

	UserRepo = &UserRepository{db}
	VideoRepo = &VideoRepository{db}

	rdb := redis.NewClient(&redis.Options{
		Addr: config.Config.RedisAddr,
	})

	dur, err := time.ParseDuration(config.Config.TokenDuration)
	if err != nil {
		panic(err)
	}

	TokenRepo = &TokenRepository{rdb, context.Background(), dur}
}
