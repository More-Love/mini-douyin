package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type TokenRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func (m *TokenRepository) GenerateToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (m *TokenRepository) SetToken(token string, userID uint) error {
	return m.rdb.Set(m.ctx, token, strconv.FormatUint(uint64(userID), 10), time.Hour*24).Err()
}

func (m *TokenRepository) GetUID(token string) uint {
	uidStr, err := m.rdb.Get(m.ctx, token).Result()
	if err != nil {
		return 0
	}
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		panic("UID is not a uint")
	}
	return uint(uid)
}
