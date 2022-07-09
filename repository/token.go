package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type TokenRepository struct {
	rdb      *redis.Client
	ctx      context.Context
	duration time.Duration
}

func (m *TokenRepository) SetToken(token string, userID int64) error {

	return m.rdb.Set(m.ctx, token, strconv.FormatInt(int64(userID), 10), m.duration).Err()
}

func (m *TokenRepository) GetUID(token string) int64 {
	uidStr, err := m.rdb.Get(m.ctx, token).Result()
	if err != nil {
		return 0
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		panic("UID is not a int")
	}
	return uid
}
