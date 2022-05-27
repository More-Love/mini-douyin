package models

import (
	"strconv"
	"time"
)

func SetToken(token string, userID uint) error {
	return rdb.Set(ctx, strconv.FormatUint(uint64(userID), 10), token, time.Hour*2).Err()
}

func GetToken(userID uint) (string, error) {
	return rdb.Get(ctx, strconv.FormatUint(uint64(userID), 10)).Result()
}
